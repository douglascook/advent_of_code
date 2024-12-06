IO.puts("Day 6 - Gallivanting Guards")

defmodule Day6 do
  def follow_the_guard(filepath) do
    IO.puts("Processing input #{filepath}")

    {map_shape, guard, obstacles} = parse_map(filepath)
    IO.inspect(guard, charlists: :as_lists, label: "Guard")
    IO.puts("#{length(obstacles)} obstacles on map.")

    {visited, _in_loop} = follow_guard(obstacles, map_shape, guard, {-1, 0}, MapSet.new([guard]))

    unique_coords = visited |> Enum.map(&elem(&1, 0)) |> MapSet.new()
    IO.puts("Guard visited #{MapSet.size(unique_coords)} squares.")

    # This is just brute force - must be some sort of dynamic programming or something that would help
    possible_loops =
      unique_coords
      |> MapSet.delete(guard)
      |> Enum.map(fn new_obstacle ->
        follow_guard(obstacles ++ [new_obstacle], map_shape, guard, {-1, 0}, MapSet.new([guard]))
      end)
      |> Enum.filter(&elem(&1, 1))
      |> Enum.count()

    IO.puts("#{possible_loops} options to create a loop by adding a single obstacle.")
  end

  def follow_guard(obstacles, map_shape, guard, direction, visited) do
    {guard, direction} = get_next_square(obstacles, guard, direction)

    {height, width} = map_shape

    # In a loop -> return visited and flag for "in loop"
    if already_visited(visited, {guard, direction}) do
      IO.puts("Found loop!")
      {visited, true}
    else
      case {guard, direction} do
        # Return number of squares visited if guard has gone OOB
        {{-1, _}, _} ->
          {visited, false}

        {{_, -1}, _} ->
          {visited, false}

        {{y, x}, _} when y >= height or x >= width ->
          {visited, false}

        # Update visited and continue following
        {{_, _}, _} ->
          follow_guard(
            obstacles,
            map_shape,
            guard,
            direction,
            MapSet.put(visited, {guard, direction})
          )
      end
    end
  end

  def already_visited(visited, guard_direction) do
    MapSet.member?(visited, guard_direction)
  end

  def get_next_square(obstacles, guard, direction) do
    y = elem(guard, 0) + elem(direction, 0)
    x = elem(guard, 1) + elem(direction, 1)

    # No obstacle -> continue in same direction
    if !Enum.member?(obstacles, {y, x}) do
      {{y, x}, direction}
    else
      # Obstacle -> turn right and try in that direction
      next_dir =
        case direction do
          {1, 0} -> {0, -1}
          {0, -1} -> {-1, 0}
          {-1, 0} -> {0, 1}
          {0, 1} -> {1, 0}
        end

      get_next_square(obstacles, guard, next_dir)
    end
  end

  def parse_map(filepath) do
    map =
      File.read!(filepath)
      |> String.split("\n", trim: true)
      |> Enum.map(&String.graphemes/1)

    guard =
      map
      |> Enum.with_index()
      |> Enum.flat_map(&find_char_in_row(&1, "^"))
      |> Enum.at(0)

    obstacles =
      map
      |> Enum.with_index()
      |> Enum.flat_map(&find_char_in_row(&1, "#"))

    {{length(map), length(hd(map))}, guard, obstacles}
  end

  def find_char_in_row({row, row_index}, target_char) do
    row
    |> Enum.with_index()
    |> Enum.filter(fn {v, _} -> v == target_char end)
    |> Enum.map(fn {_, i} -> {row_index, i} end)
  end

  def print_map(filepath, visited) do
    map =
      File.read!(filepath)
      |> String.split("\n", trim: true)
      |> Enum.map(&String.graphemes/1)
      |> Enum.with_index()
      |> Enum.map(fn {r, i} -> add_visited(r, i, visited) end)
      |> Enum.join("\n")

    IO.puts(map)
  end

  def add_visited(row, row_index, visited) do
    row
    |> Enum.with_index()
    |> Enum.map(fn {v, i} ->
      if MapSet.member?(visited, {row_index, i}) do
        "x"
      else
        v
      end
    end)
    |> List.to_string()
  end
end

Day6.follow_the_guard("../../inputs/day06/test_input.txt")
Day6.follow_the_guard("../../inputs/day06/input.txt")
