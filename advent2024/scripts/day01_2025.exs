alias Advent2024.Utils

args = System.argv()

# Determine the input type based on the argument (default to :sample)
input_type =
  case args do
    ["input"] -> :input
    _ -> :sample
  end

data =
  Utils.get_string_grapheme(Utils.get_input(input_type, Day01_2025))
  |> Stream.map(fn {x, y} -> {x, String.to_integer(String.trim(y))} end)
  |> Enum.to_list()

defmodule Day1 do
  def wrap_around(init, nr, operator) do
    delta = if operator == "+", do: nr, else: -nr
    Integer.mod(init + delta, 100)
  end
end

{counter, _state} =
  data
  |> Enum.reduce({0, 50}, fn {direction, nr}, {counter, state} ->
    cond do
      direction == "R" ->
        new_dial = Day1.wrap_around(state, nr, "+")

        cond do
          new_dial == 0 -> {counter + 1, new_dial}
          true -> {counter, new_dial}
        end

      true ->
        new_dial = Day1.wrap_around(state, nr, "-")

        cond do
          new_dial == 0 -> {counter + 1, new_dial}
          true -> {counter, new_dial}
        end
    end
  end)

IO.inspect(counter, label: "Puzzle1 password")
