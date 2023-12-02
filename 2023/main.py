import sys
import functools

from puzzles.day01 import trebuchet
from puzzles.day02 import cubes

import profiling

PUZZLES = {1: trebuchet.fire, 2: cubes.cubes}


def go(function, input_path):
    @profiling.profile_it
    @functools.wraps(function)
    def wrapped():
        function(input_path)

    wrapped()


if __name__ == "__main__":
    day = int(sys.argv[1])
    input_file = sys.argv[2]
    module = f"puzzles/day{day:02}"

    go(PUZZLES[day], f"{module}/{input_file}")
