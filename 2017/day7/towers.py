def get_wrong_weight():
    with_above = {k: t for k, t in PROGRAMS.items() if t['above']}
    bottom_program = get_bottom(with_above)
    calculate_total_weight(bottom_program)
    import ipdb; ipdb.set_trace()


def calculate_total_weight(program_name):
    program = PROGRAMS[program_name]
    if not program['above']:
        program['total_weight'] = program['weight']
        return program['weight']

    total = program['weight']
    weights = []
    for other in program['above']:
        other_weight = calculate_total_weight(other)
        if len(weights) and other_weight not in weights:
            print(f'{program_name} is imbalanced due to {other}, weights = {weights}, other = {other_weight}')
        weights.append(other_weight)
    total += sum(weights)
    program['total_weight'] = total
    return total


def get_bottom(with_above):
    all_above = []
    for data in with_above.values():
        all_above.extend(data['above'])

    for program in with_above.keys():
        if program not in all_above:
            print(f'{program} is at the bottom')
            return program


def parse_tower_data():
    input_ = open('input.txt').read()
    lines = [l for l in input_.split('\n') if l]

    programs = {}
    for line in lines:
        program_name, weight, *rest = line.split()
        programs[program_name] = {
            'weight': int(weight.strip('()')),
            'above': [r.strip(',') for r in rest[1:]]
        }
    return programs


if __name__ == '__main__':
    PROGRAMS = parse_tower_data()
    get_wrong_weight()
