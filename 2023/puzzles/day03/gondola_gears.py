import logging
import dataclasses
import pprint
import itertools

logger = logging.getLogger(__name__)


@dataclasses.dataclass
class EnginePart:
    row: int
    col: int
    value: str = ""
    length: int = 0

    def adjacent_to(self, other_part):
        # Must be in one of neighbouring rows
        if other_part.row < self.row - 1 or other_part.row > self.row + 1:
            return False

        # FIXME this only works if self is a number
        # Can be diagonally touching
        return other_part.col in range(self.col - 1, self.col + self.length + 1)


class Number(EnginePart):
    def __add__(self, character):
        self.value += character
        self.length += 1
        return self

    def to_int(self):
        return int(self.value)


def read_schematic(filepath):
    part_numbers, symbols = parse_schematic(filepath)
    calculate_part_number_sum(part_numbers, symbols)
    calculate_gear_ratio_sum(part_numbers, symbols)


def calculate_part_number_sum(part_numbers, symbols):
    part_number_sum = 0
    for i, row in enumerate(part_numbers):
        # Only symbols in surrounding rows may be adjacent
        symbol_rows = symbols[max(0, i - 1) : i + 2]

        for number in row:
            for symbol in itertools.chain.from_iterable(symbol_rows):
                if number.adjacent_to(symbol):
                    part_number_sum += number.to_int()
                    break

    print("Part number sum =", part_number_sum)


def calculate_gear_ratio_sum(part_numbers, symbols):
    gear_ratio_sum = 0
    for i, row in enumerate(symbols):
        gears = [r for r in row if r.value == "*"]
        # Only symbols in surrounding rows may be adjacent
        number_rows = part_numbers[max(0, i - 1) : i + 2]

        for gear in gears:
            adjacent_parts = []
            for part in itertools.chain.from_iterable(number_rows):
                if part.adjacent_to(gear):
                    adjacent_parts.append(part.to_int())

            if len(adjacent_parts) == 2:
                gear_ratio_sum += adjacent_parts[0] * adjacent_parts[1]

    print("Gear ratio sum =", gear_ratio_sum)


def parse_schematic(filepath):
    numbers = []
    symbols = []

    number = None
    with open(filepath) as f:
        for i, line in enumerate(f):
            # Store everything by row to make
            row_numbers = []
            row_symbols = []

            for j, char in enumerate(line.strip()):
                if char.isdigit():
                    if number is None:
                        number = Number(row=i, col=j)
                    number += char

                else:
                    if number:
                        row_numbers.append(number)
                        number = None

                    if char != ".":
                        row_symbols.append(
                            EnginePart(row=i, col=j, value=char, length=1)
                        )
            if number:
                row_numbers.append(number)
                number = None

            numbers.append(row_numbers)
            symbols.append(row_symbols)

    # print("-------------------------Part numbers-------------------------")
    # pprint.pprint(numbers)
    # print("---------------------------Symbols----------------------------")
    # pprint.pprint(symbols)

    return numbers, symbols
