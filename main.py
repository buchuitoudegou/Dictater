from FileReader import extract_words
import argparse

parser = argparse.ArgumentParser()
parser.add_argument('--dictionary', required=False, default='synonym0',  help='select a dictionary to dictate')

args = parser.parse_args()

def dictate(dictionary):
  total = len(dictionary)
  wrong = {}
  correct = 0
  for word in dictionary:
    print(f'current word: {word}')
    reference = dictionary[word]
    answer = input('answer: ')
    answer = answer.split(',')
    for w in answer:
      try:
        reference.remove(w)
      except:   
        break
    if len(reference) == 0:
      correct += 1
    else:
      wrong[word] = reference
  return {
    'total': total,
    'correct num': correct,
    'wrong words': wrong
  }

if __name__ == "__main__":
  dictionary = extract_words(args.dictionary)
  result = dictate(dictionary)
  print(result) 