IO.puts("Day 8 - Antenna Antinodes")

defmodule Day8 do
  def find_antinodes(filepath) do
    IO.puts("Processing input #{filepath}")

    {map_shape, antennae} = parse_map(filepath)

    pairs =
      antennae
      |> Map.to_list()
      |> Enum.flat_map(fn {_char, coords} -> build_pairs(coords, []) end)

    total =
      pairs
      |> Enum.flat_map(&get_antinodes_for_pair/1)
      |> Enum.filter(&is_within_bounds(&1, map_shape))
      |> MapSet.new()
      |> MapSet.size()

    IO.puts("Total number of antinodes within map = #{total}")

    intersecting_lines =
      pairs
      |> Enum.flat_map(&get_line_intersecting_pair(&1, map_shape))
      |> MapSet.new()

    antenna_coords = antennae |> Map.to_list() |> Enum.flat_map(&elem(&1, 1)) |> MapSet.new()

    total = intersecting_lines |> MapSet.union(antenna_coords) |> MapSet.size()

    IO.puts("Updated total number of antinodes = #{total}")
  end

  def build_pairs(values, pairs) when values != [] do
    [head | tail] = values
    pairs = pairs ++ Enum.map(tail, fn v -> {head, v} end)
    build_pairs(tail, pairs)
  end

  def build_pairs([], pairs) do
    pairs
  end

  def get_antinodes_for_pair({{y1, x1}, {y2, x2}}) do
    # Antinodes are found by adding/subtracting the vector between the pair of coords.
    y_delta = y1 - y2
    x_delta = x1 - x2

    [{y1 + y_delta, x1 + x_delta}, {y2 - y_delta, x2 - x_delta}]
  end

  def get_line_intersecting_pair({{y1, x1}, {y2, x2}}, map_shape) do
    y_delta = y1 - y2
    x_delta = x1 - x2

    # Normalise keeping components as integers
    gcd = Integer.gcd(y_delta, x_delta)
    normalised = {div(y_delta, gcd), div(x_delta, gcd)}

    get_points_before({y1, x1}, normalised, map_shape, []) ++
      get_points_after({y1, x1}, normalised, map_shape, [])
  end

  # TODO could tidy this up by passing in operator? instead of copy pasta
  def get_points_before({y, x}, {y_delta, x_delta}, map_shape, points) do
    next_point = {y - y_delta, x - x_delta}

    if not is_within_bounds(next_point, map_shape) do
      points
    else
      get_points_before(next_point, {y_delta, x_delta}, map_shape, points ++ [next_point])
    end
  end

  def get_points_after({y, x}, {y_delta, x_delta}, map_shape, points) do
    next_point = {y + y_delta, x + x_delta}

    if not is_within_bounds(next_point, map_shape) do
      points
    else
      get_points_after(next_point, {y_delta, x_delta}, map_shape, points ++ [next_point])
    end
  end

  def is_within_bounds({y, x}, {map_height, map_width}) do
    y >= 0 and y < map_height and x >= 0 and x < map_width
  end

  def parse_map(filepath) do
    map =
      File.read!(filepath)
      |> String.split("\n", trim: true)
      |> Enum.map(&String.graphemes/1)

    antennae =
      map
      |> Enum.with_index()
      |> Enum.flat_map(fn {row, j} -> find_antennae(row, j) end)
      |> build_antennae_lookup(%{})

    {{length(map), length(hd(map))}, antennae}
  end

  def find_antennae(row, row_index) do
    row
    |> Enum.with_index()
    |> Enum.filter(fn {c, _} -> c != "." end)
    |> Enum.map(fn {c, i} -> {c, {row_index, i}} end)
  end

  def build_antennae_lookup(antennae, lookup) when antennae != [] do
    [head | tail] = antennae
    {char, {j, i}} = head
    build_antennae_lookup(tail, Map.put(lookup, char, Map.get(lookup, char, []) ++ [{j, i}]))
  end

  def build_antennae_lookup([], lookup) do
    lookup
  end
end

Day8.find_antinodes("../../inputs/day08/test_input.txt")
Day8.find_antinodes("../../inputs/day08/test_input_2.txt")
Day8.find_antinodes("../../inputs/day08/input.txt")
