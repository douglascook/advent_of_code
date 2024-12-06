IO.puts("Day 5 - Sleigh Safety Report Printer Problems")

defmodule Day5 do
  def solve(filepath) do
    IO.puts("Processing input #{filepath}")

    {rules, page_updates} = parse_report(filepath)

    compute_valid_updates_score(rules, page_updates)
    compute_invalid_updates_score(rules, page_updates)
  end

  def compute_valid_updates_score(rules, page_updates) do
    score =
      page_updates
      |> Enum.filter(&page_update_is_valid(&1, rules, []))
      |> Enum.map(&get_middle_value/1)
      |> Enum.sum()

    IO.puts("Valid page updates score = #{score}")
  end

  def compute_invalid_updates_score(rules, page_updates) do
    score =
      page_updates
      |> Enum.reject(&page_update_is_valid(&1, rules, []))
      |> Enum.map(&fix_page_update_order(&1, rules))
      |> Enum.map(&get_middle_value/1)
      |> Enum.sum()

    IO.puts("Corrected version of incorrect page updates score = #{score}")
  end

  def page_update_is_valid(page_numbers, rules, seen) do
    [current_page | remainder] = page_numbers
    prior_pages = Map.get(rules, current_page)

    # All prior pages must appear before current page
    valid =
      if prior_pages != nil do
        prior_pages
        # Rules only apply if *both* pages numbers are present in the page updates
        |> Enum.filter(&Enum.member?(page_numbers ++ seen, &1))
        |> Enum.map(&Enum.member?(seen, &1))
        |> Enum.all?()
      else
        true
      end

    seen = seen ++ [current_page]

    case {valid, remainder} do
      {true, []} -> true
      {true, remainder} -> page_update_is_valid(remainder, rules, seen)
      {false, _} -> false
    end
  end

  def fix_page_update_order(page_numbers, rules) do
    IO.puts("Original order")
    IO.inspect(page_numbers, charlists: :as_lists)

    # Filter down rules to only include the relevant page numbers.
    matching_rules =
      rules
      |> Map.filter(fn {k, _} -> Enum.member?(page_numbers, k) end)
      |> Map.new(fn {k, v} -> {k, Enum.filter(v, &Enum.member?(page_numbers, &1))} end)

    IO.puts("Matching rules")
    IO.inspect(matching_rules, charlists: :as_lists)

    # Now sort by length of preceding pages to get correct order.
    ordered =
      matching_rules
      |> Enum.to_list()
      |> Enum.sort_by(fn {_, v} -> length(v) end)
      |> Enum.map(&elem(&1, 0))

    # If a page has no preceding page rules then it must be the first page.
    ordered =
      if length(ordered) < length(page_numbers) do
        (page_numbers -- ordered) ++ ordered
      else
        ordered
      end

    IO.puts("Fixed order")
    IO.inspect(ordered, charlists: :as_lists)
    ordered
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
      |> Enum.map(fn p -> Enum.map(String.split(p, ","), &String.to_integer(&1)) end)

    {rules, pages}
  end

  @doc "Build a map from page number to all page numbers that must occur *before* it."
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
