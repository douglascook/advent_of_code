import sys

from puzzles.day01 import trebuchet
from puzzles.day02 import cubes
from puzzles.day03 import gondola_gears
from puzzles.day04 import scratchcards
from puzzles.day05 import seeds

import profiling

PUZZLES = {
    1: trebuchet.fire,
    2: cubes.cubes,
    3: gondola_gears.read_schematic,
    4: scratchcards.scratch,
    5: seeds.plant,
}


if __name__ == "__main__":
    day = int(sys.argv[1])
    input_file = sys.argv[2]
    module = f"puzzles/day{day:02}"

    puzzle_run = PUZZLES[day]
    input_path = f"{module}/{input_file}"

    if len(sys.argv) > 3 and sys.argv[3] == "--profile":
        puzzle_run = profiling.profile_it(puzzle_run, repeats=100)

    puzzle_run(input_path)
