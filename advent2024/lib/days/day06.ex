defmodule Advent2024.Days.Day06 do
  @moduledoc """
  Solution for Day 06 of Advent of Code 2024.
  """

  require Logger
  alias Advent2024.ListUtils
  alias Advent2024.Utils

  def get_input(type \\ :sample) do
    Utils.get_character_matrix(Utils.get_input(type, __MODULE__))
  end

  def find_start(floor_map) do
    floor_map
    |> Enum.with_index()
    |> Enum.reduce_while(nil, fn {line, row}, _acc ->
      case Enum.find_index(line, &(&1 == "^")) do
        nil -> {:cont, nil}
        col -> {:halt, {row, col}}
      end
    end)
  end

  def move(map, {row, col}, row_length, _column_length, :north) do
    new_start = ListUtils.get_matrix_transposed_reversed_pos(row_length, {row, col})

    IO.inspect(new_start)

    IO.inspect(Map.get(map, :north))
  end

  def start_moving(map, {row, col}, row_length, column_length) do
    move(map, {row, col}, row_length, column_length, :north)
  end

  def part1(floor_map) do
    map = %{
      :east => floor_map,
      :west => ListUtils.reverse_matrix(floor_map),
      :south => ListUtils.transpose(floor_map)
    }

    IO.inspect(Map.get(map, :south))

    map = Map.put(map, :north, ListUtils.reverse_matrix(Map.get(map, :south)))
    IO.inspect(Map.get(map, :north))

    row_length = length(floor_map)
    column_length = length(List.first(floor_map))
    start_moving(map, find_start(floor_map), row_length, column_length)
  end
end
