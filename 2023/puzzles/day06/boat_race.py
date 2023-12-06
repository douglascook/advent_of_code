import math


def charge_up(filepath):
    with open(filepath) as f:
        lines = f.readlines()

    times = [int(t) for t in lines[0].split()[1:]]
    distances = [int(d) for d in lines[1].split()[1:]]

    winning_times_product = 1
    for t, d in zip(times, distances):
        lower_bound, upper_bound = calculate_winning_charge(t, d)
        print(lower_bound, upper_bound)
        winning_times_product *= upper_bound - lower_bound

    print("Product of all size of winning ranges =", winning_times_product)

    updated_time = int("".join(str(t) for t in times))
    updated_distance = int("".join(str(d) for d in distances))

    lower_bound, upper_bound = calculate_winning_charge(updated_time, updated_distance)
    print("Update size of winning range =", upper_bound - lower_bound)


def calculate_winning_charge(time, distance):
    """How long to charge the boat in order to reach at least <distance> in a race of
    length <time>.

    Range of winning charge times is returned with lower bound inclusive, upper bound
    exclusive.

    Distance = charge_time * (time - charge_time)
    d = n * (t - n)
    d = nt - n**2
    -n**2 + tn -d > 0

    Can solve using quadratic formula: (-b +- sqrt(b*b - 4ac)) / 2a
    where a = -1, b = t, c = -d

    n = -t +- sqrt(t*t - 4d) / -2
    """
    n1 = (-time + math.sqrt((time * time) - 4 * distance)) / -2
    n2 = (-time - math.sqrt((time * time) - 4 * distance)) / -2

    print("n1 =", n1, "n2 =", n2)
    if n1 < n2:
        if n1.is_integer():
            n1 += 0.1
        return math.ceil(n1), math.ceil(n2)

    if n2.is_integer():
        n2 += 0.1
    return math.ceil(n2), math.ceil(n1)
