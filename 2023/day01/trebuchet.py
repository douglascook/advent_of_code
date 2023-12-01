import sys

import profiling

DIGITS = {
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9,
}


def _reverse_string(string):
    return "".join(reversed(string))


REVERSED_DIGITS = {_reverse_string(w): v for w, v in DIGITS.items()}


@profiling.profile_it
def fire(filepath):
    calibration_total = 0
    updated_calibration_total = 0
    for line in open(filepath):
        try:
            calibration_total += parse_calibration_value(line.strip())
        except StopIteration:
            # No digit can be found
            pass

        updated_calibration_total += parse_updated_calibration_value(line.strip())

    print("Total calibration value =", calibration_total)
    print("Total updated calibration value =", updated_calibration_total)


def parse_calibration_value(line):
    first_digit = next(c for c in line if c.isdigit())
    last_digit = next(c for c in reversed(line) if c.isdigit())
    return int(f"{first_digit}{last_digit}")


def parse_updated_calibration_value(line):
    first_digit = find_first_digit_recursive("", line, DIGITS)
    last_digit = find_first_digit_recursive("", _reverse_string(line), REVERSED_DIGITS)
    return int(f"{first_digit}{last_digit}")


def find_first_digit(characters, digits):
    first = characters[0]
    if first.isdigit():
        return first

    for word, value in digits.items():
        if characters.startswith(word):
            return value

    return find_first_digit(characters[1:], digits)


def find_first_digit_recursive(matched, remaining, digits):
    if not remaining:
        raise ValueError("Could not find any digit spelt out")

    if matched in digits:
        return digits[matched]

    next_ = remaining[0]
    if next_.isdigit():
        return next_

    digits = {w: v for w, v in digits.items() if w[len(matched)] == next_}

    return find_first_digit_recursive(matched + next_, remaining[1:], digits)


if __name__ == "__main__":
    fire(sys.argv[1])
