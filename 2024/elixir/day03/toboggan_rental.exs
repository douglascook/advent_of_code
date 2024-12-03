IO.puts("Day 3 - North Pole Toboggan Rental")

defmodule Day3 do
  def scan_corrupted_memory(filepath) do
    IO.puts("Processing input #{filepath}")
    memory = File.read!(filepath)

    naive_scan(memory)
    full_scan(memory)
  end

  defp naive_scan(memory) do
    pattern = ~r/mul\((\d{1,3}),(\d{1,3})\)/
    matches = Regex.scan(pattern, memory)

    output =
      matches
      |> Enum.map(fn [_, x, y] -> String.to_integer(x) * String.to_integer(y) end)
      |> Enum.sum()

    IO.puts("Output value from naive scan = #{output}")
  end

  defp full_scan(memory) do
    pattern = ~r/(do\(\)|don't\(\)|mul\((\d{1,3}),(\d{1,3})\))/
    matches = Regex.scan(pattern, memory, capture: :all_but_first)

    total = compute_total(matches, true, 0)
    IO.puts("Output value from advanced scan = #{total}")
  end

  defp compute_total(memory, enabled, acc) do
    if memory == [] do
      acc
    else
      [match | remainder] = memory

      case {match, enabled} do
        {[], _} ->
          acc

        # Enable multiplication
        {["do()"], _} ->
          compute_total(remainder, true, acc)

        # Disable multiplication
        {["don't()"], _} ->
          compute_total(remainder, false, acc)

        # Multiplication disabled -> continue without adding to acc
        {["mul(" <> _rest, x, y], false} ->
          compute_total(remainder, enabled, acc)

        # Multiplication enabled -> Update acc
        {["mul(" <> _rest, x, y], true} ->
          compute_total(remainder, enabled, acc + String.to_integer(x) * String.to_integer(y))
      end
    end
  end
end

Day3.scan_corrupted_memory("../../inputs/day03/test_input.txt")
Day3.scan_corrupted_memory("../../inputs/day03/test_input_2.txt")
Day3.scan_corrupted_memory("../../inputs/day03/input.txt")
