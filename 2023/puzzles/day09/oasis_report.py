def extrapolate(filepath):
    next_values_sum = 0
    prev_values_sum = 0

    for history in load_histories(filepath):
        print("History", history)

        deltas = calculate_deltas(history)
        next_values_sum += predict_next_value(history, deltas)
        prev_values_sum += predict_previous_value(history, deltas)

    print("Sum of next values =", next_values_sum)
    print("Sum of previous values =", prev_values_sum)


def predict_next_value(history, deltas):
    next_value = history[-1] + sum(d[-1] for d in deltas)
    print("Next value =", next_value)
    return next_value


def predict_previous_value(history, deltas):
    previous = deltas[-1][0]
    for d in reversed(deltas[:-1]):
        previous = d[0] - previous

    print("Previous value =", previous)
    return history[0] - previous


def calculate_deltas(history):
    sequences = []
    deltas = history
    while len(set(deltas)) != 1:
        deltas = [deltas[i + 1] - deltas[i] for i in range(len(deltas) - 1)]
        sequences.append(deltas)

    print("Sequences", sequences)
    return sequences


def load_histories(filepath):
    with open(filepath) as f:
        for line in f:
            yield [int(n) for n in line.split()]
