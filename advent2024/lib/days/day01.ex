defmodule Advent2024.Days.Day01 do
  @moduledoc """
  Solution for Day 01 of Advent of Code 2024.
  """

  alias Advent2024.Utils

  def get_input(type \\ :sample) do
    Utils.get_split_strings(Utils.get_input(type, __MODULE__))
    |> Enum.map(fn [l, r] -> {String.to_integer(l), String.to_integer(r)} end)
    # # Before unzip: [{3, 4}, {4, 3}, {2, 5}, {1, 3}, {3, 9}, {3, 3}]
    |> Enum.unzip()

    # After unzip: {[3, 4, 2, 1, 3, 3], [4, 3, 5, 3, 9, 3]}
  end

  def total_distance({left, right}) do
    Enum.zip(Enum.sort(left), Enum.sort(right))
    |> Enum.reduce(0, fn {l, r}, acc -> acc + abs(l - r) end)
  end

  def puzzle2({left, right}) do
    map_vals =
      right
      |> Enum.reduce(%{}, fn val, acc ->
        Map.update(acc, val, 1, fn existing_value -> existing_value + 1 end)
      end)

    left
    |> Enum.reduce(0, fn val, acc -> acc + val * Map.get(map_vals, val, 0) end)
  end
end
