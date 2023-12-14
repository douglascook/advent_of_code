def reflect(filepath):
    with open(filepath) as f:
        mirrors = [m.split() for m in f.read().split("\n\n")]

    for remove_smudges in (False, True):
        reflection_total = 0

        for m in mirrors:
            score = get_reflection_score(m, remove_smudges)
            reflection_total += score

        print(
            "Total reflection score with remove_smudges",
            remove_smudges,
            "=",
            reflection_total,
        )


def get_reflection_score(mirror, remove_smudges=False):
    print()
    print("\n".join(mirror))
    for i in range(1, len(mirror)):
        if is_horizontal_reflection_line(mirror, i, remove_smudges):
            print("Horizontal reflection at row", i)
            return 100 * i

    for j in range(1, len(mirror[0])):
        if is_vertical_reflection_line(mirror, j, remove_smudges):
            print("Vertical reflection at column", j)
            return j
    return 0


def is_horizontal_reflection_line(mirror, i, remove_smudges):
    cleaned = False

    d = 0
    while i - d - 1 >= 0 and i + d < len(mirror):
        if mirror[i - d - 1] != mirror[i + d]:
            if (
                remove_smudges
                and not cleaned
                and row_is_smudged(mirror, i - d - 1, i + d)
            ):
                cleaned = True
            else:
                return False
        d += 1

    if not remove_smudges:
        return True

    # Only cleaned mirrors count when removing smudges
    return cleaned


def row_is_smudged(mirror, row, other_row):
    row_length = len(mirror[row])
    mismatches = (mirror[row][j] != mirror[other_row][j] for j in range(row_length))
    return sum(mismatches) == 1


def is_vertical_reflection_line(mirror, j, remove_smudges):
    cleaned = False

    d = 0
    while j - d - 1 >= 0 and j + d < len(mirror[0]):
        left = [m[j - d - 1] for m in mirror]
        right = [m[j + d] for m in mirror]
        if left != right:
            if remove_smudges and not cleaned and col_is_smudged(left, right):
                cleaned = True
            else:
                return False
        d += 1

    if not remove_smudges:
        return True

    return cleaned


def col_is_smudged(col, other_col):
    mismatches = [col[i] != other_col[i] for i in range(len(col))]
    return sum(mismatches) == 1
