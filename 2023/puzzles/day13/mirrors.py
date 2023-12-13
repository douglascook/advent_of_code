def reflect(filepath):
    with open(filepath) as f:
        mirrors = [m.split() for m in f.read().split("\n\n")]

    reflection_total = 0
    for m in mirrors:
        score = find_horizontal_reflection(m)
        if not score:
            score = find_vertical_reflection(m)

        reflection_total += score

    print("Total reflection score =", reflection_total)


def find_horizontal_reflection(mirror):
    for i in range(1, len(mirror)):
        # Found potential line of reflection
        if mirror[i] == mirror[i - 1]:
            d = 0
            reflected = True

            while i - d - 1 >= 0 and i + d < len(mirror):
                if mirror[i - d - 1] != mirror[i + d]:
                    reflected = False
                    break
                d += 1

            if reflected:
                return 100 * i

    return 0


def find_vertical_reflection(mirror):
    for j in range(1, len(mirror[0])):
        column = [m[j] for m in mirror]
        previous = [m[j - 1] for m in mirror]
        if column == previous:
            d = 0
            reflected = True

            while j - d - 1 >= 0 and j + d < len(mirror[0]):
                left = [m[j - d - 1] for m in mirror]
                right = [m[j + d] for m in mirror]
                if left != right:
                    reflected = False
                    break
                d += 1

            if reflected:
                return j
    return 0
