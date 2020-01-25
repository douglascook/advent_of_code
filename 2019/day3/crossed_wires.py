def get_wire_coordinates(wire):
    steps = wire.split(',')

    coords = []
    current_x = 0
    current_y = 0

    for step in steps:
        direction = step[0]
        step_length = int(step[1:])

        if direction == 'U':
            for _ in range(step_length):
                current_y += 1
                coords.append((current_x, current_y))

        elif direction == 'D':
            for _ in range(step_length):
                current_y -= 1
                coords.append((current_x, current_y))

        elif direction == 'R':
            for _ in range(step_length):
                current_x += 1
                coords.append((current_x, current_y))

        elif direction == 'L':
            for _ in range(step_length):
                current_x -= 1
                coords.append((current_x, current_y))

    return coords


def find_intersections(coords_1, coords_2):
    intersections = set(c[:2] for c in coords_1) & set(c[:2] for c in coords_2)
    print('The wires intersect at', intersections)
    return list(intersections)


def find_closest_distance_from_origin(intersections):
    distances = (abs(x) + abs(y) for x, y in intersections)
    return min(distances)


def find_min_number_steps(intersections, coords_1, coords_2):
    # need to add 1 to each since index zero corresponds with one steptaken
    step_counts = [coords_1.index(i) + 1 + coords_2.index(i) + 1 for i in intersections]
    return min(step_counts)


if __name__ == '__main__':
    wire1, wire2 = open('input.txt').read().strip().split('\n')

    coords_1 = get_wire_coordinates(wire1)
    coords_2 = get_wire_coordinates(wire2)
    intersections = find_intersections(coords_1, coords_2)

    print('PART 1 - find closest point to origin where the wires intersect')
    min_distance = find_closest_distance_from_origin(intersections)
    print('Manhattan distance of the closest intersection from origin =', min_distance)

    print('PART 2 - find lowest combined numbers of steps taken')
    min_step = find_min_number_steps(intersections, coords_1, coords_2)
    print('Minimum combined number of steps taken =', min_step)
