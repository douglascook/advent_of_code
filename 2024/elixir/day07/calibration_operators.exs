IO.puts("Day 7 - Calibration Report Operators")

defmodule Day7 do
  def calibrate(filepath) do
    IO.puts("Processing input #{filepath}")

    rows =
      File.read!(filepath)
      |> String.split("\n", trim: true)
      |> Enum.map(&parse_row/1)

    valid = rows |> Enum.filter(fn {t, v} -> validate_row(t, [hd(v)], tl(v)) end)
    total = valid |> Enum.map(&elem(&1, 0)) |> Enum.sum()
    IO.puts("Calibration result for valid equations = #{total}")

    valid =
      rows
      |> Enum.filter(fn {t, v} -> validate_row(t, [hd(v)], tl(v), with_concat_operator: true) end)

    total = valid |> Enum.map(&elem(&1, 0)) |> Enum.sum()
    IO.puts("Calibration result for valid equations, with concat operator included = #{total}")
  end

  def parse_row(row) do
    [target, values] = String.split(row, ":", trim: true)
    target = String.to_integer(target)

    values = values |> String.split(" ", trim: true) |> Enum.map(&String.to_integer(&1))

    {target, values}
  end

  def validate_row(target, totals, remaining_values, with_concat_operator \\ false) do
    [next | remaining] = remaining_values

    sums = Enum.map(totals, fn t -> t + next end)
    products = Enum.map(totals, fn t -> t * next end)

    concats =
      if with_concat_operator do
        Enum.map(totals, fn t ->
          String.to_integer(Integer.to_string(t) <> Integer.to_string(next))
        end)
      else
        []
      end

    # Drop any totals that are higher than target value.
    new_totals = Enum.reject(sums ++ products ++ concats, fn t -> t > target end)

    case {remaining, Enum.member?(new_totals, target)} do
      {[], true} -> true
      {[], false} -> false
      {_, _} -> validate_row(target, new_totals, remaining, with_concat_operator)
    end
  end
end

Day7.calibrate("../../inputs/day07/test_input.txt")
Day7.calibrate("../../inputs/day07/input.txt")
