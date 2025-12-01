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
  def wrap_around(current_pos, rotations, operator) do
    rotation = if operator == "+", do: rotations, else: -rotations
    Integer.mod(current_pos + rotation, 100)
  end

  def step_to_zero(current_pos, rotations, operator, acc) do
    if operator == "+" do
      clicks_to_zero = 100 - current_pos

      cond do
        # It will pass 0
        clicks_to_zero < rotations ->
          step_to_zero(0, rotations - clicks_to_zero, operator, acc + 1)

        true ->
          acc
      end
    else
      clicks_to_zero = if current_pos == 0, do: 100, else: current_pos

      cond do
        clicks_to_zero < rotations ->
          step_to_zero(0, rotations - clicks_to_zero, operator, acc + 1)

        true ->
          acc
      end
    end
  end

  def dial_count(current_pos, rotations, operator) do
    step_to_zero(current_pos, rotations, operator, 0)
  end
end

{counter, _state} =
  data
  |> Enum.reduce({0, 50}, fn {direction, nr}, {counter, state} ->
    cond do
      direction == "R" ->
        new_dial = Day1.wrap_around(state, nr, "+")
        passed = Day1.dial_count(state, nr, "+")

        cond do
          new_dial == 0 -> {counter + 1 + passed, new_dial}
          true -> {counter + passed, new_dial}
        end

      true ->
        new_dial = Day1.wrap_around(state, nr, "-")
        passed = Day1.dial_count(state, nr, "-")

        cond do
          new_dial == 0 -> {counter + 1 + passed, new_dial}
          true -> {counter + passed, new_dial}
        end
    end
  end)

IO.inspect(counter, label: "Puzzle1 password")
