package reader

import (
	"encoding/json"
	"os"
)

type Word struct {
	Semantic string   `json:"semantic"`
	Nouns    []string `json:"noun,omitempty"`
	Verbs    []string `json:"verb,omitempty"`
	Adjs     []string `json:"adj,omitempty"`
}

type Vocabulary = []Word

func (w *Word) UnMarshal(s string) error {
	t := Word{}
	err := json.Unmarshal([]byte(s), &t)
	if err != nil {
		return err
	}
	w = &t
	return nil
}

func readFile(path string) string {
	f, err := os.Open(path)
	if err != nil {
		return ""
	}
	buffer := make([]byte, 102400) // 100KB
	ret := ""
	var length int
	for {
		if length, err = f.Read(buffer); err != nil {
			break
		}
		ret += string(buffer[:length])
		buffer = make([]byte, 102400)
	}
	return ret
}

func convertToVocabulary(s string) Vocabulary {
	var v Vocabulary
	err := json.Unmarshal([]byte(s), &v)
	if err != nil {
		panic("unmarshal string failed: " + err.Error())
	}
	return v
}

func GetVocabulary(path string) Vocabulary {
	s := readFile(path)
	return convertToVocabulary(s)
}
