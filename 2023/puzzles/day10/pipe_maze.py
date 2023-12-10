import enum


class Direction(enum.Enum):
    NORTH = (-1, 0)
    SOUTH = (1, 0)
    EAST = (0, 1)
    WEST = (0, -1)


# Coming into pipe from key, next direction is value
PIPE_TO_NEXT = {
    "|": {Direction.NORTH: Direction.NORTH, Direction.SOUTH: Direction.SOUTH},
    "F": {Direction.NORTH: Direction.EAST, Direction.WEST: Direction.SOUTH},
    "7": {Direction.NORTH: Direction.WEST, Direction.EAST: Direction.SOUTH},
    "-": {Direction.EAST: Direction.EAST, Direction.WEST: Direction.WEST},
    "L": {Direction.SOUTH: Direction.EAST, Direction.WEST: Direction.NORTH},
    "J": {Direction.SOUTH: Direction.WEST, Direction.EAST: Direction.NORTH},
    "S": {
        Direction.NORTH: "DONE",
        Direction.SOUTH: "DONE",
        Direction.EAST: "DONE",
        Direction.WEST: "DONE",
    },
}


def search(filepath):
    pipe_map = sketch_pipes(filepath)
    start, start_neighbours = find_start(pipe_map)

    loop_points = set([start])
    # Pick first neigbour, arbitrarily
    pipe, direction = start_neighbours[0]

    print("Following route")
    while pipe not in loop_points:
        loop_points.add(pipe)
        pipe, next_pipe, direction = move(pipe, direction, pipe_map)

    draw_loop(pipe_map, loop_points)
    print("Furthest point in loop is", len(loop_points) // 2, "steps from start")


def sketch_pipes(filepath):
    pipe_map = []
    with open(filepath) as f:
        for line in f:
            pipe_map.append(list(line.strip()))
    return pipe_map


def find_start(pipe_map):
    start = None
    for i, row in enumerate(pipe_map):
        if "S" in row:
            start = i, row.index("S")
            break

    if not start:
        raise ValueError("No Start!?")

    pipe_neigbours = []
    for direction in Direction:
        try:
            next_, _, next_direction = move(start, direction, pipe_map)
            pipe_neigbours.append((next_, next_direction))
        except ValueError:
            continue
    assert len(pipe_neigbours) == 2

    return start, pipe_neigbours


def move(p, direction, pipe_map):
    next_ = p[0] + direction.value[0], p[1] + direction.value[1]
    try:
        next_pipe = pipe_map[next_[0]][next_[1]]
    except IndexError:
        raise ValueError("Cannot move over edge of map", next_)

    next_direction = PIPE_TO_NEXT.get(next_pipe, {}).get(direction)
    # print(direction, next_pipe, next_direction)
    if next_direction is None:
        raise ValueError(
            "Cannot move", direction, "from", p, "next pipe invalid", next_pipe
        )
    return next_, next_pipe, next_direction


def draw_loop(pipe_map, points):
    plot = []
    for i in range(len(pipe_map)):
        plot.append([])
        for j in range(len(pipe_map[0])):
            if (i, j) in points:
                plot[i].append("x")
            else:
                plot[i].append(".")
    for row in plot:
        print("".join(row))


# FIXME this is broken but general approach should work. Search for first pipe
# after encountering non-pipe, checking for start and end of line edge cases.
def find_enclosed_sections(pipe_map, points):
    row_enclosures = []
    for i in range(len(pipe_map)):
        row_enclosures.append([])

        start = None
        enclosed = False
        for j in range(len(pipe_map[0])):
            if (i, j) not in points and enclosed and start is None:
                start = j

            elif (i, j) in points and start is not None:
                row_enclosures[i].append((start, j))
                start = None
                enclosed = False

            elif (i, j) in points:
                enclosed = True

    print(row_enclosures)

    col_enclosures = []
    for j in range(len(pipe_map[0])):
        col_enclosures.append([])

        start = None
        enclosed = False
        for i in range(len(pipe_map)):
            if (i, j) not in points and enclosed and start is None:
                start = i

            elif (i, j) in points and start is not None:
                col_enclosures[j].append((start, i))
                start = None
                enclosed = False

            elif (i, j) in points:
                enclosed = True

    print(col_enclosures)
