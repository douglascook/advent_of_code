IO.puts("Day 4 - Xmas Word Search")

defmodule Day4 do
  def solve(filepath) do
    IO.puts("Processing input #{filepath}")

    puzzle = parse_puzzle(filepath)
    count_xmas_occurrences(puzzle)
    count_cross_mas_occurrences(puzzle)
  end

  def count_xmas_occurrences(puzzle) do
    find_matches_in_puzzle = &find_matches(puzzle, &1)

    total_matches =
      get_char_coords(puzzle, "X")
      |> Enum.map(find_matches_in_puzzle)
      |> Enum.sum()

    IO.puts("Found #{total_matches} matching XMAS words")
  end

  def get_char_coords(puzzle, target_char) do
    find_target = &find_char_in_row(&1, target_char)

    puzzle
    |> Enum.with_index()
    |> Enum.flat_map(find_target)
  end

  def find_char_in_row({row, row_index}, target_char) do
    row
    |> Enum.with_index()
    |> Enum.filter(fn {v, _} -> v == target_char end)
    |> Enum.map(fn {_, i} -> {row_index, i} end)
  end

  def find_matches(puzzle, coords) do
    # Down, Up, Left, Right, Up-Left, Up-Right, Down-Left, Down-Right
    directions = [{1, 0}, {-1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}]
    Enum.count(directions, fn d -> is_xmas_match(puzzle, coords, d, "M") end)
  end

  def is_xmas_match(puzzle, coords, direction, target_char) do
    # Move to the next character in given direction
    y = elem(coords, 0) + elem(direction, 0)
    x = elem(coords, 1) + elem(direction, 1)

    # Out of bounds
    if y < 0 or x < 0 or y >= length(puzzle) or x >= length(hd(puzzle)) do
      false
    else
      char = Enum.at(Enum.at(puzzle, y), x)
      # Check if it matches expected target
      case {char, target_char} do
        {"M", "M"} -> is_xmas_match(puzzle, {y, x}, direction, "A")
        {"A", "A"} -> is_xmas_match(puzzle, {y, x}, direction, "S")
        {"S", "S"} -> true
        {_, _} -> false
      end
    end
  end

  def count_cross_mas_occurrences(puzzle) do
    total_matches =
      get_char_coords(puzzle, "A")
      |> Enum.count(fn c ->
        is_cross_mas_part(puzzle, c, 1, 1) and is_cross_mas_part(puzzle, c, 1, -1)
      end)

    IO.puts("Found #{total_matches} matching cross-MAS")
  end

  def is_cross_mas_part(puzzle, {y, x}, y_delta, x_delta) do
    # Skip out of bounds
    if y == 0 or x == 0 or y == length(puzzle) - 1 or x == length(hd(puzzle)) - 1 do
      false
    else
      one = Enum.at(Enum.at(puzzle, y + y_delta), x + x_delta)
      opposite = Enum.at(Enum.at(puzzle, y - y_delta), x - x_delta)

      case {one, opposite} do
        {"M", "S"} -> true
        {"S", "M"} -> true
        {_, _} -> false
      end
    end
  end

  def parse_puzzle(filepath) do
    File.read!(filepath)
    |> String.split("\n", trim: true)
    |> Enum.map(&String.graphemes/1)
  end
end

Day4.solve("../../inputs/day04/test_input.txt")
Day4.solve("../../inputs/day04/input.txt")
