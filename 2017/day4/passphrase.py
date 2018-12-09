def count_valid_phrases():
    with open('input.txt') as f:
        input_text = f.read()
    lines = [line.split() for line in input_text.split('\n') if line]
    return sum([len(line) == len(set(line)) for line in lines])


def count_valid_non_anagrams():
    with open('input.txt') as f:
        input_text = f.read()
    lines = [line.split() for line in input_text.split('\n') if line]
    sorted_words = []
    for line in lines:
        sorted_words.append([''.join(sorted(word)) for word in line])
    return sum([len(line) == len(set(line)) for line in sorted_words])


if __name__ == '__main__':
    print(count_valid_phrases())
    print(count_valid_non_anagrams())
