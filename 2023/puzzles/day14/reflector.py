def tilt(filepath):
    rocks_map = read_rocks_map(filepath)

    rolled = tilt_north(rocks_map)
    load = measure_load(rolled)
    print("Total load on support beams after tilting north =", load)

    spun = read_rocks_map(filepath)
    cycle, cycle_start = find_cycle(spun)
    print("Found cycle of length", len(cycle), "at iteration", cycle_start)

    # Subtract initial states outside cycle and then see what's left after repeating
    # the cycle.
    target = (1000000000 - cycle_start) % len(cycle)
    load = measure_load(cycle[target])
    print("Total load on support beams after 1000000000 spins =", load)


def find_cycle(spun):
    previous_states = [deepcopy(spun)]

    i = 0
    while True:
        for tilt in [tilt_north, tilt_west, tilt_south, tilt_east]:
            spun = tilt(spun)

        i += 1
        if spun in previous_states:
            cycle_start = previous_states.index(spun)
            return previous_states[cycle_start:], cycle_start

        previous_states.append(deepcopy(spun))


def tilt_north(rocks_map):
    """Roll rocks north until they all stop."""
    rolled, moved = roll_north(rocks_map)
    while moved:
        rolled, moved = roll_north(rolled)
    return rolled


def tilt_south(rocks_map):
    """Roll rocks south until they all stop."""
    rolled = tilt_north(list(reversed(rocks_map)))
    return list(reversed(rolled))


def tilt_west(rocks_map):
    """Roll rocks west until they all stop."""
    rotated = rotate_map_right(rocks_map)
    rolled = tilt_north(rotated)
    return rotate_map_left(rolled)


def tilt_east(rocks_map):
    """Roll rocks east until they all stop."""
    rotated = rotate_map_left(rocks_map)
    rolled = tilt_north(rotated)
    return rotate_map_right(rolled)


def rotate_map_right(rocks_map):
    rotated = []
    for j in range(len(rocks_map[0])):
        row = [rocks_map[i][j] for i in range(len(rocks_map) - 1, -1, -1)]
        rotated.append(row)
    return rotated


def rotate_map_left(rocks_map):
    rotated = []
    for j in range(len(rocks_map[0]) - 1, -1, -1):
        row = [rocks_map[i][j] for i in range(len(rocks_map))]
        rotated.append(row)
    return rotated


def roll_north(rocks_map):
    """Roll rocks one square north."""
    after_tilt = []
    rocks_moved = False

    prev_row = rocks_map[0]
    for i in range(1, len(rocks_map)):
        row = rocks_map[i]
        row_updated = row.copy()

        for j in range(len(row)):
            if row[j] == "O" and prev_row[j] == ".":
                prev_row[j] = "O"
                row_updated[j] = "."
                rocks_moved = True

        after_tilt.append(prev_row)
        prev_row = row_updated

    after_tilt.append(prev_row)

    return after_tilt, rocks_moved


def measure_load(rocks_map):
    total_load = 0
    for i, row in enumerate(rocks_map):
        load = len(rocks_map) - i
        rock_count = sum(r == "O" for r in row)
        total_load += load * rock_count
    return total_load


def read_rocks_map(filepath):
    rocks = []
    with open(filepath) as f:
        for line in f:
            rocks.append(list(line.strip()))
    return rocks


def deepcopy(list_):
    return [r.copy() for r in list_]
