from collections import Counter
from enum import IntEnum

CARD_VALUES = {str(i): i for i in range(2, 10)}
CARD_VALUES.update({"T": 10, "J": 11, "Q": 12, "K": 13, "A": 14})

CARD_VALUES_WITH_JOKER = CARD_VALUES.copy()
CARD_VALUES_WITH_JOKER["J"] = 1


class HandType(IntEnum):
    HIGH_CARD = 1
    PAIR = 2
    TWO_PAIR = 3
    THREE_OF_KIND = 4
    FULL_HOUSE = 5
    FOUR_OF_KIND = 6
    FIVE_OF_KIND = 7


class Hand:
    def __init__(self, line):
        cards, bid = line.split()
        self.cards = cards
        self.bid = int(bid)
        self.hand_type = self._calculate_type()

    def _calculate_type(self):
        card_counts = Counter(self.cards)
        max_count = max(card_counts.values())

        if max_count == 5:
            return HandType.FIVE_OF_KIND

        if max_count == 4:
            return HandType.FOUR_OF_KIND

        if max_count == 3:
            if len(card_counts) == 2:
                return HandType.FULL_HOUSE
            else:
                return HandType.THREE_OF_KIND

        if max_count == 2:
            if len(card_counts) == 3:
                return HandType.TWO_PAIR
            else:
                return HandType.PAIR

        return HandType.HIGH_CARD

    def __eq__(self, other):
        return self.cards == other.cards

    def __lt__(self, other, values=CARD_VALUES):
        if self.hand_type == other.hand_type:
            for i in range(5):
                if self.cards[i] == other.cards[i]:
                    continue
                return values[self.cards[i]] < values[other.cards[i]]

        return self.hand_type < other.hand_type

    def __gt__(self, other, values=CARD_VALUES):
        if self.hand_type == other.hand_type:
            for i in range(5):
                if self.cards[i] == other.cards[i]:
                    continue
                return values[self.cards[i]] > values[other.cards[i]]

        return self.hand_type > other.hand_type

    def __repr__(self):
        return f"Hand: {self.cards} {self.hand_type.name}"


# FIXME where is the bug in this???
class JokerHand(Hand):
    def _calculate_type(self):
        card_counts = Counter(self.cards)

        joker_count = card_counts["J"]
        if joker_count == 0:
            return super()._calculate_type()

        max_count = max(card_counts.values())

        # Joker is one of the values so can match up with other value
        if len(card_counts) <= 2:
            return HandType.FIVE_OF_KIND

        # Otherwise there is one set of three, with two others.
        # Add joker to three or vice versa
        if max_count == 3:
            return HandType.FOUR_OF_KIND

        if max_count == 2:
            # Two pairs and one other
            if len(card_counts) == 3:
                # Joker is one of the pairs, add to other pair
                if joker_count == 2:
                    return HandType.FOUR_OF_KIND
                # Joker is single card
                return HandType.THREE_OF_KIND

            # One pair and three others.
            # If joker is the pair then update to match any of the others, if not
            # then update to match the pair.
            return HandType.THREE_OF_KIND

        # All different values, joker pairs up with any
        return HandType.PAIR

    def __lt__(self, other):
        return super().__lt__(other, values=CARD_VALUES_WITH_JOKER)

    def __gt__(self, other):
        return super().__gt__(other, values=CARD_VALUES_WITH_JOKER)


def deal(filepath):
    with open(filepath) as f:
        lines = f.readlines()

    hands = [Hand(l) for l in lines]
    calculate_winnings(hands)

    joker_hands = [JokerHand(l) for l in lines]
    calculate_winnings(joker_hands)


def calculate_winnings(hands):
    ranked_hands = sorted(hands)

    winnings = 0
    for rank, hand in enumerate(ranked_hands, start=1):
        print(hand, hand.bid)
        winnings += rank * hand.bid

    print("Total winnings =", winnings)
