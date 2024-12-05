IO.puts("Day 5 - Sleigh Safety Report Printer Problems")

defmodule Day5 do
  def solve(filepath) do
    IO.puts("Processing input #{filepath}")

    {rules, page_updates} = parse_report(filepath)
    IO.inspect(rules, charlists: :as_lists)
    IO.inspect(page_updates, charlists: :as_lists)

    compute_valid_updates_score(rules, page_updates)
  end

  def compute_valid_updates_score(rules, page_updates) do
    score =
      page_updates
      |> Enum.filter(fn p -> page_update_is_valid(p, rules, []) end)
      |> Enum.map(&get_middle_value/1)
      |> Enum.sum()

    IO.puts("Valid page updates score = #{score}")
  end

  def page_update_is_valid(page_numbers, rules, seen) do
    IO.puts("All remaining page numbers")
    IO.inspect(page_numbers, charlists: :as_lists)

    [current_page | remainder] = page_numbers
    prior_pages = Map.get(rules, current_page)

    IO.puts("Page number #{current_page}")
    IO.inspect(remainder, charlists: :as_lists)
    IO.inspect(seen, charlists: :as_lists)
    # All prior pages must appear before current page
    valid =
      if prior_pages != nil do
        prior_pages
        # Rules only apply if *both* pages numbers are present in the page updates
        |> Enum.filter(fn p -> Enum.member?(page_numbers ++ seen, p) end)
        |> Enum.map(fn p -> Enum.member?(seen, p) end)
        |> Enum.all?()
      else
        true
      end

    IO.inspect(valid)

    seen = seen ++ [current_page]

    case {valid, remainder} do
      {true, []} -> true
      {true, remainder} -> page_update_is_valid(remainder, rules, seen)
      {false, _} -> false
    end
  end

  def get_middle_value(list) do
    Enum.at(list, div(length(list), 2))
  end

  def parse_report(filepath) do
    [rules, pages] = String.split(File.read!(filepath), "\n\n", trim: true)

    rules = parse_rules(String.split(rules), %{})

    pages =
      pages
      |> String.split()
      |> Enum.map(fn p -> Enum.map(String.split(p, ","), fn n -> String.to_integer(n) end) end)

    {rules, pages}
  end

  def parse_rules([head | tail], lookup) do
    [before, after_] = Enum.map(String.split(head, "|"), &String.to_integer/1)

    page_rules = Map.get(lookup, after_, []) ++ [before]
    lookup = Map.put(lookup, after_, page_rules)

    if tail == [] do
      lookup
    else
      parse_rules(tail, lookup)
    end
  end
end

Day5.solve("../../inputs/day05/test_input.txt")
Day5.solve("../../inputs/day05/input.txt")
