def compute_total_required_fuel(accumulator, mass):
    if mass <= 0:
        return accumulator

    new_required = compute_required_fuel(mass)
    accumulator += new_required
    return compute_total_required_fuel(accumulator, new_required)


def compute_required_fuel(mass):
    return max(0, mass // 3 - 2)


if __name__ == "__main__":
    total = 0
    with open("input.txt") as input:
        for line in input:
            module_mass = int(line)
            total += compute_total_required_fuel(0, module_mass)

    print(f"Total required fuel = {total}")
