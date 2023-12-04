import re

CARD_PATTERN = re.compile(r"^Card\s+(\d+): ([\d ]+) \| ([\d ]+)$")


def scratch(filepath):
    scores = []
    winner_counts = []
    cards = list(parse_cards(filepath))

    for card in cards:
        winner_count, score = calculate_score(*card)
        scores.append(score)
        winner_counts.append(winner_count)
    print("Total original score =", sum(scores))

    calculate_card_explosion(winner_counts)


def parse_cards(filepath):
    for line in open(filepath):
        match = CARD_PATTERN.match(line.strip())
        winning_numbers = [int(n) for n in match.group(2).split()]
        numbers = [int(n) for n in match.group(3).split()]
        yield winning_numbers, numbers


def calculate_score(winning_numbers, numbers):
    winners = set(winning_numbers).intersection(set(numbers))
    if not winners:
        return 0, 0

    winner_count = len(winners)
    return winner_count, 2 ** (winner_count - 1)


def calculate_card_explosion(winner_counts):
    card_counts = [1 for _ in range(len(winner_counts))]

    for i, winner_count in enumerate(winner_counts):
        for j in range(i + 1, i + winner_count + 1):
            card_counts[j] += card_counts[i]

    print("Total number of cards =", sum(card_counts))
