alias Advent2025.Utils

data =
  Utils.get_split_strings(Utils.input_type(System.argv(), Day06))
  |> Advent2025.ListUtils.transpose()
  |> Enum.map(&Enum.sort/1)

defmodule Solve do
  def calculate_line(operator, number_list) do
    numbers = Enum.map(number_list, &String.to_integer/1)

    case operator do
      "+" -> Enum.sum(numbers)
      "*" -> Enum.product(numbers)
    end
  end

  def puzzle1(math_homework) do
    math_homework
    |> Enum.reduce(0, fn [operator | tail], acc ->
      acc + calculate_line(operator, tail)
    end)
  end
end

{time, {one, two}} =
  :timer.tc(fn ->
    {Solve.puzzle1(data), 0}
    # Solve.puzzle2(ranges)}
  end)

IO.puts("Puzzle 1: #{one}\nPuzzle 2: #{two}\nExecuted in #{time / 1_000}ms")
