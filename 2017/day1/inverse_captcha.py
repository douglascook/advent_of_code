with open('input.txt') as f:
    input_text = f.read()


def matches_at_distance_total(distance):
    total = 0
    import ipdb; ipdb.set_trace()
    number_digits = len(input_text)
    for i in range(number_digits):
        if input_text[i] == input_text[(i + distance) % number_digits]:
            total += int(input_text[i])
    print(total)
    return total


if __name__ == '__main__':
    matches_at_distance_total(1)
    matches_at_distance_total(len(input_text)//2)
