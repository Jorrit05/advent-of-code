alias Advent2025.Utils

input_type = Utils.input_type(System.argv())
data = Utils.get_character_matrix(Utils.get_input(input_type, Day04))

defmodule Day4 do
  def puzzle1(data) do
    data
    |> Enum.map(fn row ->
      Enum.with_index()
    end)
  end
  end

  # def solve(data, analyzer_fn) do
  #   data
  #   |> Enum.map(fn str ->
  #     analyzer_fn.(str, start)
  #   end)
  #   |> Enum.sum()
  # end
end

{time, {one, two}} =
  :timer.tc(fn ->
    {Day4.puzzle1(data), 0}
    #  Day4.solve(data, &Day4.puzzle2/2)}
  end)

IO.puts("Puzzle 1: #{one}\nPuzzle 2: #{two}\nExecuted in #{time / 1_000}ms")
