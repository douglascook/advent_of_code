IO.puts("Day 2 - Reindeer Reactor Reports!")

defmodule Day2 do
  def go(filepath) do
    IO.puts("Processing input #{filepath}")

    valid_count =
      parse_reports(filepath)
      |> Enum.map(fn r -> Day2.report_is_valid(r) end)
      |> Enum.sum()

    IO.puts("Number of valid reports = #{valid_count}")
  end

  defp parse_reports(filepath) do
    File.read!(filepath)
    |> String.split("\n", trim: true)
    |> Enum.map(fn s -> String.split(s) end)
    |> Enum.map(fn row -> Enum.map(row, fn v -> String.to_integer(v) end) end)
  end

  def report_is_valid(report) do
    [head | tail] = report
    delta = head - hd(tail)
    report_is_valid(tail, delta)
  end

  # Presumably can tidy this up with more idiomatic Elixir? Pattern matching?
  def report_is_valid(report, prev_delta) do
    if not delta_is_valid(prev_delta) do
      0
    else
      [head | tail] = report

      if tail == [] do
        1
      else
        delta = head - hd(tail)

        if prev_delta * delta > 0 do
          report_is_valid(tail, delta)
        else
          0
        end
      end
    end
  end

  defp delta_is_valid(delta) do
    delta != 0 and abs(delta) <= 3
  end
end

Day2.go("../../inputs/day02/test_input.txt")
Day2.go("../../inputs/day02/input.txt")
