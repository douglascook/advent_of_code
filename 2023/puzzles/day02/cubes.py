import collections

BAG_COUNTS = {"red": 12, "green": 13, "blue": 14}


def cubes(filepath):
    valid_game_sum = 0
    min_set_power_sum = 0
    for game_id, game in parse_games(filepath):
        if game_is_valid(game):
            valid_game_sum += game_id

        min_set_power_sum += calculate_minimum_set_power(game)

    print("Sum of valid game IDs is", valid_game_sum)
    print("Sum of minimum set powers is", min_set_power_sum)


def game_is_valid(game):
    for draw in game:
        for colour, count in draw.items():
            bag_count = BAG_COUNTS.get(colour)
            if bag_count is None or count > bag_count:
                return False
    return True


def calculate_minimum_set_power(game):
    max_counts = collections.defaultdict(lambda: 0)
    for draw in game:
        for colour, count in draw.items():
            if max_counts[colour] < count:
                max_counts[colour] = count
    power = 1
    for v in max_counts.values():
        power *= v
    return power


def parse_games(filepath):
    for line in open(filepath):
        name, draws = line.split(": ")
        yield int(name.split(" ")[1]), parse_cube_counts(draws)


def parse_cube_counts(draws):
    cube_counts = []
    for draw in draws.split("; "):
        counts = {}
        for count in draw.split(", "):
            n, colour = count.split()
            counts[colour] = int(n)
        cube_counts.append(counts)
    return cube_counts
