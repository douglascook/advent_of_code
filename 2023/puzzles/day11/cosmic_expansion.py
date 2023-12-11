def expand(filepath):
    cosmos = [list(l.strip()) for l in open(filepath)]

    galaxies = []
    for i in range(len(cosmos)):
        for j in range(len(cosmos[i])):
            if cosmos[i][j] == "#":
                galaxies.append((i, j))
    print(galaxies)

    cols_to_expand = get_cols_to_expand(cosmos)
    rows_to_expand = get_rows_to_expand(cosmos)

    print("Expanding by factor of  2")
    expanded = expand_cosmos(galaxies, rows_to_expand, cols_to_expand, 2)
    measure_distances_between_galaxies(expanded)

    print("Expanding by factor of 10")
    expanded = expand_cosmos(galaxies, rows_to_expand, cols_to_expand, 10)
    measure_distances_between_galaxies(expanded)

    print("Expanding by factor of 100")
    expanded = expand_cosmos(galaxies, rows_to_expand, cols_to_expand, 100)
    measure_distances_between_galaxies(expanded)

    print("Expanding by factor of 1000000")
    expanded = expand_cosmos(galaxies, rows_to_expand, cols_to_expand, 1000000)
    measure_distances_between_galaxies(expanded)


def measure_distances_between_galaxies(galaxies):
    print(galaxies)
    distances = []
    for i, g1 in enumerate(galaxies):
        for j in range(i + 1, len(galaxies)):
            g2 = galaxies[j]
            distances.append(distance_between(*g1, *g2))

    print(
        "Total distance between", len(distances), "pairs of galaxies is", sum(distances)
    )


def get_cols_to_expand(cosmos):
    cols_to_expand = []
    for j in range(len(cosmos[0])):
        if all(row[j] == "." for row in cosmos):
            cols_to_expand.append(j)
    return cols_to_expand


def get_rows_to_expand(cosmos):
    return [i for i, row in enumerate(cosmos) if all(r == "." for r in row)]


def expand_cosmos(galaxies, rows_to_expand, cols_to_expand, expansion_factor):
    expanded = []
    for x, y in galaxies:
        # Subract one from expansion factor to account for existing blank row
        row_shift = (expansion_factor - 1) * len([r for r in rows_to_expand if r < x])
        col_shift = (expansion_factor - 1) * len([c for c in cols_to_expand if c < y])
        expanded.append((x + row_shift, y + col_shift))

    # print_cosmos(expanded)
    return expanded


def distance_between(x1, y1, x2, y2):
    return abs(x1 - x2) + abs(y1 - y2)


def print_cosmos(galaxies):
    x_limit = max(x for x, y in galaxies) + 1
    y_limit = max(y for x, y in galaxies) + 1

    cosmos = []
    for i in range(x_limit):
        cosmos.append(["." for j in range(y_limit)])

    for x, y in galaxies:
        cosmos[x][y] = "#"

    for line in cosmos:
        print("".join(line))
