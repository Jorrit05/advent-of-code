alias Advent2025.Utils

input_type = Utils.input_type(System.argv())

# data =
defmodule Day3 do
  def solve(data, analyzer_fn) do
    data
    |> Stream.flat_map(fn {start, finish} ->
      Stream.map(start..finish, fn n ->
        str = Integer.to_string(n)
        analyzer_fn.(str, String.length(str))
      end)
    end)
    |> Enum.sum()
  end
end

{time, {one, two}} =
  :timer.tc(fn ->
    {Day3.solve(data, &Day2.analyze_number/2), 0}
    #  Day3.solve(data, &Day2.difficult_analysis(&1, &2, 1))}
  end)

IO.puts("Puzzle 1: #{one}\nPuzzle 2: #{two}\nExecuted in #{time / 1_000}ms")
