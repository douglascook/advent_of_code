import re
from collections import namedtuple

Mapping = namedtuple("Mapping", ["start", "end", "delta"])
Range = namedtuple("Range", ["start", "end"])

ALMANAC_PATTERN = re.compile(
    "^"
    "seeds:([\d\s]+)"
    "seed-to-soil map:([\d\s]+)"
    "soil-to-fertilizer map:([\d\s]+)"
    "fertilizer-to-water map:([\d\s]+)"
    "water-to-light map:([\d\s]+)"
    "light-to-temperature map:([\d\s]+)"
    "temperature-to-humidity map:([\d\s]+)"
    "humidity-to-location map:([\d\s]+)"
    "$"
)


def plant(filepath):
    seeds, maps = read_almanac(filepath)

    locations = [get_location_for_seed(s, maps) for s in seeds]
    print("Lowest location number =", min(locations))

    seed_ranges = [
        Range(seeds[i], seeds[i] + seeds[i + 1]) for i in range(0, len(seeds), 2)
    ]
    location_ranges = get_location_ranges(seed_ranges, maps)

    print(
        "Lowest seed range input location number =",
        min(l.start for l in location_ranges),
    )


def get_location_for_seed(seed, maps):
    for map_ in maps:
        for mapping in map_:
            if mapping.start <= seed < mapping.end:
                seed += mapping.delta
                # Mappings are not chained together, each map modifies each value at most once.
                break
    return seed


def get_location_ranges(seed_ranges, maps):
    ranges = seed_ranges

    for map_ in maps:
        updated_ranges = []
        for r in ranges:
            updated = apply_map(map_, r)
            assert all(r.start < r.end for r in updated)
            updated_ranges.extend(updated)
        ranges = updated_ranges
    return ranges


def apply_map(map_, range_):
    """Applying a map to a range will produce 1 - 3 updated ranges:

    Cases:
        * map and range disjoint -> range is unchanged
            [ ---- range ------]
                                  [ --- map ---- ]

        * map contains range -> entire range updated
            [ ----------- map -----------------]
                   [ --- range ---- ]
                   [ --- updated -- ]

        * range contains map -> outer ranges unchanged, inner one updated
            [ ----------- range -----------------]
                   [ --- map ---- ]
            [ -r- ][ ---updated---][ --- r ------]


        * overlap but neither container in other -> half of range updated, other bit unchanged
                [ ----------- range -----------------]
           [ --- map ---- ]
                [ updated ][---------------r---------]

    """
    # Values that have been mapped should not be mapped again.
    mapped = []
    # Values that were not mapped but are part of a new smaller range should be
    # considered by later maps.
    unmapped = [range_]

    for m in map_:
        for r in unmapped:
            # Range and map disjoint
            if r.end <= m.start or r.start >= m.end:
                unmapped = [r]
            # Map contains entire range
            elif m.start <= r.start and m.end >= r.end:
                mapped.append(Range(r.start + m.delta, r.end + m.delta))
                unmapped = []
            # Range contains entire map
            elif r.start <= m.start and r.end >= m.end:
                mapped.append(Range(m.start + m.delta, m.end + m.delta))
                unmapped = []
                if r.start != m.start:
                    unmapped.append(Range(r.start, m.start))
                if r.end != m.end:
                    unmapped.append(Range(m.end, r.end))
            # Map overlaps on left
            elif m.start < r.start and m.end < r.end:
                mapped.append(Range(r.start + m.delta, m.end + m.delta))
                unmapped = [Range(m.end, r.end)]
            # Map overlaps on right
            else:
                mapped.append(Range(m.start + m.delta, r.end + m.delta))
                unmapped = [Range(r.start, m.start)]

    return mapped + unmapped


def read_almanac(filepath):
    with open(filepath) as f:
        match = ALMANAC_PATTERN.match(f.read())

    seeds, *maps = (g.strip().split("\n") for g in match.groups())

    seeds = [int(s) for s in seeds[0].split()]
    maps = [[parse_mapping(m) for m in map_] for map_ in maps]

    return seeds, maps


def parse_mapping(mapping_string):
    dest_start, source_start, length = [int(m) for m in mapping_string.split()]

    return Mapping(
        start=source_start,
        end=source_start + length,
        delta=dest_start - source_start,
    )
