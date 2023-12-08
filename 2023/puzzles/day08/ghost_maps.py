import re
import math

NODE_PATTERN = re.compile(r"^(\w{3}) = \((\w{3}), (\w{3})\)$")


def count_steps(filepath):
    with open(filepath) as f:
        lines = f.readlines()

    directions = lines[0].strip()
    network = parse_network(lines[2:])

    follow_directions(network, directions)
    follow_ghost_directions(network, directions)


def parse_network(lines):
    network = {}
    for line in lines:
        node, left, right = NODE_PATTERN.match(line).groups()
        network[node] = {"L": left, "R": right}
    return network


def follow_directions(network, directions):
    node = "AAA"
    direction_count = len(directions)

    steps = 0
    while node != "ZZZ":
        direction = directions[steps % direction_count]
        node = network[node][direction]
        steps += 1
    print("Found route for human in", steps, "steps")


def follow_ghost_directions(network, directions):
    nodes = [k for k in network.keys() if k[-1] == "A"]

    periods = [find_path_period(network, directions, n) for n in nodes]
    # When the periods are align all ghosts are at an end node
    steps = math.lcm(*periods)
    print("Found route for ghosts in", steps, "steps")


def find_path_period(network, directions, start):
    node = start
    direction_count = len(directions)

    first_found = 0
    steps = 0
    while True:
        direction = directions[steps % direction_count]
        node = network[node][direction]
        if node[-1] == "Z":
            print(
                "Found",
                node,
                "at steps",
                steps,
                "direction index",
                steps % direction_count,
            )
            # Thought we'd also need to check that we're at same direction index in case
            # there are multiple routes to the end points but that isn't the case,
            # all the direction indexes match up.
            if first_found:
                period = steps - first_found
                print("Path from", start, "has period", period)
                return period

            first_found = steps
        steps += 1
