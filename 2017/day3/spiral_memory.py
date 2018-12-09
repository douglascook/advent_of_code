def distance_from_centre(target):
    ring = find_ring(target)
    offset = find_target_offset_in_ring(ring, target)
    print(ring + offset)
    return ring + offset


def find_ring(target_square):
    """Return depth of ring containing the target square.

    This is the distance in some direction from the row/column containing 1.
    """
    highest = 1
    rings = 0
    while highest < target_square:
        rings += 1
        highest = max_number_in_ring(rings)
    return rings


def find_target_offset_in_ring(ring, target):
    """Returns the offset of the given target from the central row/column."""
    edges = get_ring_edges(ring)
    for edge in edges:
        try:
            index = edge.index(target)
            centre = len(edge) // 2 - 1
            return abs(index - centre)
        except ValueError:
            continue


def get_ring_edges(ring):
    ring_values = range(max_number_in_ring(ring - 1) + 1,
                        max_number_in_ring(ring) + 1)
    slice_size = len(ring_values) // 4
    edges  = [ring_values[i * slice_size : (i + 1) * slice_size] for i in range(4)]
    return edges


def max_number_in_ring(ring):
    return (ring * 2 + 1) ** 2


if __name__ == '__main__':
    distance_from_centre(12)
    distance_from_centre(23)
    distance_from_centre(1024)
    distance_from_centre(289326)
