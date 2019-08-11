root_path = './vocabulary/'

def extract_words(filename):
  dictionary = {}
  with open(root_path + filename) as f:
    all_words = f.read()
    all_words = all_words.split('\n')
    for line in all_words:
      l = line.split(',')
      key = l[0]
      answer = l[1:]
      dictionary[key] = answer
  return dictionary