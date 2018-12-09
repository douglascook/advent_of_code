def do_it():
    input_ = open('input.txt').read()
    blocks = [int(b) for b in input_.split()]

    iterations = 0
    seen = [make_block_string(blocks)]

    while True:
        largest = max(blocks)
        index = blocks.index(largest)
        blocks[index] = 0

        for i in range(1, largest + 1):
            blocks[(index + i) % len(blocks)] += 1
        iterations += 1

        block_string = make_block_string(blocks)
        if block_string in seen:
            print(f'Took {iterations} iterations until repeat')
            print(f'Cycle has length {iterations - seen.index(block_string)}')
            return iterations

        seen.append(block_string)


def make_block_string(blocks):
    return '_'.join(str(b) for b in blocks)


if __name__ == '__main__':
    do_it()
