import re
from collections import defaultdict

input_ = open('input.txt').read()
lines = [l for l in input_.split('\n') if l]
parser = re.compile(r'([a-z]+) (inc|dec) (-?\d+) if ([a-z]+) (.*)')

registers = defaultdict(lambda: 0)
max_value = 0

for line in lines:
    parsed = re.match(parser, line)
    to_update, operator, value, to_check, check = parsed.groups()

    if eval(f'{registers[to_check]} {check}'):
        if operator == 'inc':
            registers[to_update] += int(value)
        else:
            registers[to_update] -= int(value)
        if registers[to_update] > max_value:
            max_value = registers[to_update]

print(f'Max value at the end of the procedure == {max(registers.values())}')
print(f'Max value at any point during the procedure == {max_value}')
