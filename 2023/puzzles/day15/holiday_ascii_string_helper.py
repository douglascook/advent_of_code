def focus_lenses(filepath):
    with open(filepath) as f:
        initialisation_sequence = f.read().strip().split(",")

    hash_sum = sum(compute_hash(step) for step in initialisation_sequence)
    print("Sum of hashes for all steps in sequence =", hash_sum)

    boxes = arrange_lenses(initialisation_sequence)
    total_focus_power = compute_focus_power(boxes)
    print("Total focussing power of lens configuration =", total_focus_power)


def compute_hash(string):
    value = 0
    for c in string:
        value = ((value + ord(c)) * 17) % 256
    return value


def arrange_lenses(initialisation_sequence):
    boxes = [[] for i in range(256)]

    for step in initialisation_sequence:
        if step[-1] == "-":
            label = step[:-1]
            box_number = compute_hash(label)
            box = boxes[box_number]
            # Find and remove lens with matching label from that box (if present)
            try:
                index = [b[0] for b in box].index(label)
                updated = box[:index] + box[index + 1 :]
                boxes[box_number] = updated
            except ValueError:
                pass

        else:
            label, focal_length = step.split("=")
            box_number = compute_hash(label)
            box = boxes[box_number]
            try:
                # If a lens with label is already in box then replace it.
                index = [b[0] for b in box].index(label)
                box[index] = (label, int(focal_length))
            except ValueError:
                # Otherwise place lens at back of the box
                box.append((label, int(focal_length)))

            # TODO maybe don't need this if it's mutating?
            boxes[box_number] = box

    return boxes


def compute_focus_power(boxes):
    total_power = 0

    for i, box in enumerate(boxes, start=1):
        for j, (label, focal_length) in enumerate(box, start=1):
            lens_power = i * j * focal_length
            total_power += lens_power

    return total_power
