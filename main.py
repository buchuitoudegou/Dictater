from FileReader import extract_words, root_path
import os
from tqdm import tqdm


def dictate(dictionary):
  total = len(dictionary)
  wrong = {}
  correct = 0
  for word in tqdm(dictionary):
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
  dicts = os.listdir(root_path)
  for i in range(len(dicts)):
    print(f'{i}: {dicts[i]}')
  choice = input('select your dictionary (input the number): ')
  dictionary = extract_words(dicts[i])
  result = dictate(dictionary)
  print(result) 