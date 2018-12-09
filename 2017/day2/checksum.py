def checksum(line_checksum):
    with open('input.txt') as f:
        input_ = f.read()

    total = 0
    for line in input_.split('\n'):
        entries = [int(e) for e in line.split('\t')]
        total += line_checksum(entries)
    print(total)
    return total


def min_max_checksum(entries):
    return max(entries) - min(entries)


def divisible_checksum(entries):
    for i in range(len(entries)):
        entry = entries[i]
        for other in entries[i + 1:]:
            smaller, larger = sorted([entry, other])
            if larger % smaller == 0:
                return larger//smaller



if __name__ == '__main__':
    checksum(min_max_checksum)
    checksum(divisible_checksum)
