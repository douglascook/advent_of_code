def setup_program(noun, verb):
    memory = PROGRAM.copy()
    memory[1] = noun
    memory[2] = verb
    return memory


def run_program(program, index=0):
    op_code, val1, val2, update = program[index : index + 4]

    if op_code == 99:
        return program

    if op_code == 1:
        value = program[val1] + program[val2]

    elif op_code == 2:
        value = program[val1] * program[val2]

    program[update] = value
    return run_program(program, index + 4)


def find_inputs_producing(output):
    for noun in range(100):
        for verb in range(100):
            memory = setup_program(noun, verb)
            done = run_program(memory)
            if memory[0] == output:
                return noun, verb
    return Exception('No inputs found producing output')


if __name__ == '__main__':
    PROGRAM = [int(c) for c in open('input.txt').read().strip().split(',')]

    print('PART 1 - find value at position zero with noun = 12, verb = 2')
    memory = setup_program(12, 2)
    done = run_program(memory)
    print('Program finished running, value at position zero =', PROGRAM[0])

    print('PART 2 - find noun and verb inputs that produce output of 19690720')
    noun, verb = find_inputs_producing(19690720)
    print('Inputs found 100 * noun + verb =', 100 * noun + verb)
