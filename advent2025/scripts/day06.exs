alias Advent2025.Utils

input_type = Utils.input_type(System.argv())

dat = Utils.get_raw_file(Utils.get_input(input_type, Day06))

defmodule Solve do
end

{time, {one, two}} =
  :timer.tc(fn ->
    {Solve.puzzle1(ranges, ingredients), 0}
    # Solve.puzzle2(ranges)}
  end)

IO.puts("Puzzle 1: #{one}\nPuzzle 2: #{two}\nExecuted in #{time / 1_000}ms")
