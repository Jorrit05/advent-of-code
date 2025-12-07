alias Advent2025.Utils

[ranges | [ingredients]] =
  Utils.get_raw_file(Utils.input_type(System.argv(), Day05))
  |> String.split("\n\n", trim: true)

ranges = String.split(ranges, "\n", trim: true)
ingredients = String.split(ingredients, "\n", trim: true)

defmodule Solve do
  def fresh?([], _ingredient), do: 0

  def fresh?([head | tail], ingredient) do
    [start, finish] = String.split(head, "-")

    if ingredient >= String.to_integer(start) and ingredient <= String.to_integer(finish),
      do: 1,
      else: fresh?(tail, ingredient)
  end

  def puzzle1(ranges, ingredients) do
    ingredients
    |> Enum.reduce(0, fn ingredient, acc ->
      acc + fresh?(ranges, String.to_integer(ingredient))
    end)
  end

  def merge_ranges([]), do: []

  def merge_ranges([last]), do: [last]

  def merge_ranges([
        {start1, finish1} = first,
        {start2, finish2} = second
        | tail
      ]) do
    cond do
      finish1 + 1 >= start2 ->
        merge_ranges([{start1, max(finish1, finish2)} | tail])

      true ->
        [first | merge_ranges([second | tail])]
    end
  end

  def puzzle2(ranges) do
    ranges
    |> Enum.map(fn range ->
      [start, finish] = String.split(range, "-")
      {String.to_integer(start), String.to_integer(finish)}
    end)
    |> Enum.sort()
    |> merge_ranges()
    |> Enum.reduce(0, fn {start, finish}, acc ->
      acc + (finish - start + 1)
    end)
  end
end

{time, {one, two}} =
  :timer.tc(fn ->
    {Solve.puzzle1(ranges, ingredients), Solve.puzzle2(ranges)}
  end)

IO.puts("Puzzle 1: #{one}\nPuzzle 2: #{two}\nExecuted in #{time / 1_000}ms")
