alias Advent2025.Utils

input_type = Utils.input_type(System.argv())

data =
  Utils.get_raw_file(Utils.get_input(input_type, Day04))
  |> String.split(",")
  |> Enum.map(fn range_str ->
    [start, finish] = String.split(range_str, "-")
    {String.to_integer(start), String.to_integer(finish)}
  end)

defmodule Day4 do
end

{time, {one, two}} =
  :timer.tc(fn ->
    {Day4.solve(data, &Day4.puzzle1/2), 0}
    #  Day4.solve(data, &Day4.puzzle2/2)}
  end)

IO.puts("Puzzle 1: #{one}\nPuzzle 2: #{two}\nExecuted in #{time / 1_000}ms")
