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

    # TODO figure out start direction?
    loop = [(start, "S", "UNKNOWN")]
    loop_points = set([start])
    # Pick first neigbour, arbitrarily
    point, pipe, direction = start_neighbours[0]

    print("Following route")
    while point not in loop_points:
        loop_points.add(point)
        loop.append((point, pipe, direction))
        point, pipe, direction = move(point, direction, pipe_map)

    draw_loop(pipe_map, loop_points)
    print("Furthest point in loop is", len(loop_points) // 2, "steps from start")

    # find_enclosed_sections(pipe_map, loop_points)
    inside_outside_map(pipe_map, loop)


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
            next_, next_pipe, next_direction = move(start, direction, pipe_map)
            pipe_neigbours.append((next_, next_pipe, next_direction))
        except ValueError:
            continue
    assert len(pipe_neigbours) == 2

    return start, pipe_neigbours


def move(point, direction, pipe_map):
    next_ = point[0] + direction.value[0], point[1] + direction.value[1]
    try:
        next_pipe = pipe_map[next_[0]][next_[1]]
    except IndexError:
        raise ValueError("Cannot move over edge of map", next_)

    next_direction = PIPE_TO_NEXT.get(next_pipe, {}).get(direction)
    # print(direction, next_pipe, next_direction)
    if next_direction is None:
        raise ValueError(
            "Cannot move", direction, "from", point, "next pipe invalid", next_pipe
        )
    return next_, next_pipe, next_direction


def draw_loop(pipe_map, points):
    plot = []
    for i in range(len(pipe_map)):
        plot.append([])
        for j in range(len(pipe_map[0])):
            if (i, j) in points:
                if pipe_map[i][j] == "S":
                    plot[i].append("S")
                else:
                    plot[i].append("x")
            else:
                plot[i].append(".")
    for row in plot:
        print("".join(row))


def inside_outside_map(pipe_map, loop):
    map_ = [["." for _ in range(len(pipe_map[0]))] for _ in range(len(pipe_map))]

    prev_x, prev_y = loop[0][0]
    map_[prev_x][prev_y] = "+"

    # Direction is the direction in which you *exit* the current pipe
    for (x, y), pipe, direction in loop[1:]:
        map_[x][y] = "+"

        if pipe == "|":
            if direction is Direction.SOUTH:
                try_update(map_, x, y - 1, "R")
                try_update(map_, x, y + 1, "L")
            else:
                try_update(map_, x, y - 1, "L")
                try_update(map_, x, y + 1, "R")
        elif pipe == "-":
            if direction is Direction.EAST:
                try_update(map_, x - 1, y, "L")
                try_update(map_, x + 1, y, "R")
            else:
                try_update(map_, x - 1, y, "R")
                try_update(map_, x + 1, y, "L")

        # Fill in pipes on outside of corner, nothing to fill on inside
        elif pipe == "F":
            side = "L" if direction is Direction.NORTH else "R"
            try_update(map_, x, y - 1, side)
            try_update(map_, x - 1, y - 1, side)
            try_update(map_, x - 1, y, side)

        elif pipe == "L":
            side = "L" if direction is Direction.NORTH else "R"
            try_update(map_, x, y - 1, side)
            try_update(map_, x + 1, y - 1, side)
            try_update(map_, x + 1, y, side)

        elif pipe == "J":
            side = "R" if direction is Direction.NORTH else "L"
            try_update(map_, x, y + 1, side)
            try_update(map_, x + 1, y + 1, side)
            try_update(map_, x + 1, y, side)

        elif pipe == "7":
            side = "L" if direction is Direction.SOUTH else "R"
            try_update(map_, x - 1, y, side)
            try_update(map_, x, y + 1, side)
            try_update(map_, x - 1, y + 1, side)

        # print("\n\n")
        # print("x", x, "y", y, "Pipe", pipe, "direction", direction)
        # for row in map_:
        #     print("".join(row))

    for row in map_:
        print("".join(row))


def try_update(map_, x, y, value):
    if x < 0 or y < 0:
        return

    try:
        if map_[x][y] != "+":
            map_[x][y] = value
    except IndexError:
        pass
