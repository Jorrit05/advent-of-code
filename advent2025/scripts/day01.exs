alias Advent2025.Utils

args = System.argv()

# Determine the input type based on the argument (default to :sample)
input_type =
  case args do
    ["input"] -> :input
    _ -> :sample
  end

data =
  Utils.get_string_grapheme(Utils.get_input(input_type, Day01))
  |> Stream.map(fn {x, y} -> {x, String.to_integer(String.trim(y))} end)
  |> Enum.to_list()

defmodule Day1 do
  def wrap_around(current_pos, rotations, "R"), do: Integer.mod(current_pos + rotations, 100)
  def wrap_around(current_pos, rotations, _), do: Integer.mod(current_pos - rotations, 100)

  def step_to_zero(current_pos, rotations, "R", acc) do
    clicks_to_zero = 100 - current_pos

    cond do
      # It will pass 0
      clicks_to_zero < rotations ->
        step_to_zero(0, rotations - clicks_to_zero, "R", acc + 1)

      true ->
        acc
    end
  end

  def step_to_zero(current_pos, rotations, "L", acc) do
    clicks_to_zero = if current_pos == 0, do: 100, else: current_pos

    cond do
      clicks_to_zero < rotations ->
        step_to_zero(0, rotations - clicks_to_zero, "L", acc + 1)

      true ->
        acc
    end
  end
end

{counter, passed, _state} =
  data
  |> Enum.reduce({0, 0, 50}, fn {direction, nr}, {exact_zero, passed_zero, state} ->
    new_dial_pos = Day1.wrap_around(state, nr, direction)
    passed = Day1.step_to_zero(state, nr, direction, 0)

    exact_zero = if new_dial_pos == 0, do: exact_zero + 1, else: exact_zero
    {exact_zero, passed_zero + passed, new_dial_pos}
  end)

IO.inspect({counter, counter + passed}, label: "Day01 Solution 1 and 2")
