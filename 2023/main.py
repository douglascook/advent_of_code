import sys

from puzzles.day01 import trebuchet
from puzzles.day02 import cubes
from puzzles.day03 import gondola_gears
from puzzles.day04 import scratchcards
from puzzles.day05 import seeds
from puzzles.day06 import boat_race
from puzzles.day07 import camel_cards
from puzzles.day08 import ghost_maps
from puzzles.day09 import oasis_report
from puzzles.day10 import pipe_maze
from puzzles.day11 import cosmic_expansion
from puzzles.day12 import hot_springs
from puzzles.day13 import mirrors
from puzzles.day14 import reflector
from puzzles.day15 import holiday_ascii_string_helper

import profiling

PUZZLES = {
    1: trebuchet.fire,
    2: cubes.cubes,
    3: gondola_gears.read_schematic,
    4: scratchcards.scratch,
    5: seeds.plant,
    6: boat_race.charge_up,
    7: camel_cards.deal,
    8: ghost_maps.count_steps,
    9: oasis_report.extrapolate,
    10: pipe_maze.search,
    11: cosmic_expansion.expand,
    12: hot_springs.repair,
    13: mirrors.reflect,
    14: reflector.tilt,
    15: holiday_ascii_string_helper.focus_lenses,
}


if __name__ == "__main__":
    day = int(sys.argv[1])
    input_file = sys.argv[2]
    module = f"puzzles/day{day:02}"

    puzzle_run = PUZZLES[day]
    input_path = f"{module}/{input_file}"

    if len(sys.argv) > 3 and sys.argv[3] == "--profile":
        puzzle_run = profiling.profile_it(puzzle_run, repeats=int(sys.argv[4]))

    puzzle_run(input_path)
