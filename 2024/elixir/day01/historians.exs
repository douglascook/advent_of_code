IO.puts("Day 1 - Historian Hysteria")

defmodule Day1 do
  def go() do
    IO.puts("Processing test input...")
    {list1, list2} = parse("../../inputs/day01/test_input.txt")
    calculateDistance(list1, list2)
    calculateSimilarity(list1, list2)

    IO.puts("Processing puzzle input...")
    {list1, list2} = parse("../../inputs/day01/input.txt")
    calculateDistance(list1, list2)
    calculateSimilarity(list1, list2)
  end

  defp parse(filepath) do
    File.read!(filepath)
    |> String.split("\n", trim: true)
    |> Enum.map(fn s -> String.split(s) end)
    |> Enum.map(fn [x, y] -> {String.to_integer(x), String.to_integer(y)} end)
    |> Enum.unzip()
  end

  defp calculateDistance(list1, list2) do
    distance =
      Enum.zip(Enum.sort(list1), Enum.sort(list2))
      |> Enum.map(fn {x, y} -> abs(x - y) end)
      |> Enum.sum()

    IO.puts("Distance = #{distance}")
  end

  defp calculateSimilarity(list1, list2) do
    similarity =
      list1
      |> Enum.map(fn v -> calculateOccurrenceScore(v, list2) end)
      |> Enum.sum()

    IO.puts("Similarity = #{similarity}")
  end

  defp calculateOccurrenceScore(target, values) do
    occurrences =
      values
      |> Enum.map(fn v -> if v == target, do: 1, else: 0 end)
      |> Enum.sum()

    occurrences * target
  end
end

Day1.go()
