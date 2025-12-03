alias Advent2025.Utils

input_type = Utils.input_type(System.argv())

data = Utils.get_string_list(Utils.get_input(input_type, Day03))

defmodule Day3 do
  def highest(number, high, rest) when byte_size(number) == 1 do
    {high, rest}
  end

  def highest(<<first::binary-size(1), tail::binary>> = number, high, rest)
      when byte_size(number) > 0 do
    if first > high, do: highest(tail, first, tail), else: highest(tail, high, rest)
  end

  def highest(_, high, rest) do
    {high, rest}
  end

  def lowest(<<first::binary-size(1), tail::binary>> = number, low)
      when byte_size(number) > 0 do
    lowest(tail, max(first, low))
  end

  def lowest(_, low), do: low

  def puzzle1(number, high) do
    {highest, rest} = highest(number, high, "")
    lowest = lowest(rest, "0")

    String.to_integer(highest <> lowest)
  end

  def puzzle2_sol(_, _, acc) when byte_size(acc) == 12 do
    String.to_integer(acc)
  end

  def puzzle2_sol(number, battery_length, acc) when length(number) <= battery_length do
    rest = Enum.map_join(number, &elem(&1, 0))

    String.to_integer(acc <> rest)
  end

  def puzzle2_sol(number, battery_length, acc) do
    candidates = Enum.take(number, length(number) - battery_length + 1)

    {{nr, _original_idx}, position} =
      candidates
      |> Enum.with_index()
      |> Enum.max_by(fn {{val, _}, pos} -> {val, -pos} end)

    new_acc = acc <> nr
    new_nr = Enum.drop(number, position + 1)
    puzzle2_sol(new_nr, battery_length - 1, new_acc)
  end

  def puzzle2(number, battery_length) do
    graphemes =
      String.graphemes(number)
      |> Enum.with_index()

    puzzle2_sol(graphemes, battery_length, "")
  end

  # AI solution.....
  # def puzzleai(number, _) do
  #   graphemes = String.graphemes(number)

  #   graphemes
  #   |> Enum.with_index()
  #   |> Enum.flat_map(fn {first, i} ->
  #     graphemes
  #     |> Enum.drop(i + 1)
  #     |> Enum.map(fn second -> String.to_integer(first <> second) end)
  #   end)
  #   |> Enum.max()
  # end

  def solve(data, analyzer_fn, start) do
    data
    |> Enum.map(fn str ->
      analyzer_fn.(str, start)
    end)
    |> Enum.sum()
  end
end

{time, {one, two}} =
  :timer.tc(fn ->
    {Day3.solve(data, &Day3.puzzle1/2, "0"), Day3.solve(data, &Day3.puzzle2/2, 12)}
  end)

IO.puts("Puzzle 1: #{one}\nPuzzle 2: #{two}\nExecuted in #{time / 1_000}ms")
