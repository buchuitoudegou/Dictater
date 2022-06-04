package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/buchuitoudegou/dictater/reader"
	"github.com/spf13/cobra"
)

const defaultDictateCnt = 1

func DictateSemantic(v reader.Vocabulary) {
	for i := 0; i < defaultDictateCnt; i++ {
		word := v[i]
		fmt.Printf("semantic: %s, please enter %d noun(s), %d verb(s), and %d adj(s)\n", word.Semantic, len(word.Nouns), len(word.Verbs), len(word.Adjs))
		expected := make(map[string]interface{}, 0)
		for _, w := range word.Nouns {
			expected[w] = true
		}
		for _, w := range word.Verbs {
			expected[w] = true
		}
		for _, w := range word.Adjs {
			expected[w] = true
		}
		ioReader := bufio.NewReader(os.Stdin)
		input, err := ioReader.ReadString('\n')
		if err != nil {
			panic("read input: " + err.Error())
		}
		input = strings.Split(input, "\n")[0]
		inputWords := strings.Split(input, " ")
		unexpected := make([]string, 0)
		correct := make([]string, 0)
		for _, w := range inputWords {
			if _, ok := expected[w]; ok {
				delete(expected, w)
				correct = append(correct, w)
			} else {
				unexpected = append(unexpected, w)
			}
		}
		leftWords := make([]string, 0)
		for w := range expected {
			leftWords = append(leftWords, w)
		}
		fmt.Printf("correct: %v, incorrect: %v, expected more: %v\n", correct, unexpected, leftWords)
	}
}

func NewDictateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Short: "dictation",
		RunE: func(cmd *cobra.Command, _ []string) error {
			path, err := cmd.Flags().GetString("path")
			if err != nil {
				return err
			}
			v := reader.GetVocabulary(path)
			DictateSemantic(v)
			return nil
		},
	}
	cmd.Flags().String("path", "vocabulary/gre/semantic.json", "setting vocabulary path")
	return cmd
}

func main() {
	rootCmd := NewDictateCommand()
	args := os.Args[1:]
	rootCmd.SetArgs(args)
	rootCmd.Execute()
}
