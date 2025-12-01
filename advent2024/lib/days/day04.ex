defmodule Advent2024.Days.Day04 do
  @moduledoc """
  Solution for Day 04 of Advent of Code 2024.
  """

  alias Advent2024.Utils
  alias Advent2024.ListUtils

  def get_input(type \\ :sample) do
    Utils.get_character_matrix(Utils.get_input(type, __MODULE__))
  end

  def count_xmas(row) do
    xmas = ["X", "M", "A", "S"]

    count =
      Enum.chunk_every(row, 4, 1, :discard)
      |> Enum.count(&(&1 == xmas))

    count_reversed =
      Enum.chunk_every(Enum.reverse(row), 4, 1, :discard)
      |> Enum.count(&(&1 == xmas))

    count + count_reversed
  end

  def part1(graphemes) do
    horizontal = Enum.reduce(graphemes, 0, fn row, acc -> acc + count_xmas(row) end)

    transposed = ListUtils.transpose(graphemes)
    vertical = Enum.reduce(transposed, 0, fn row, acc -> acc + count_xmas(row) end)

    # Positive-slope diagonals
    grid = graphemes

    diagonals_top_left =
      for start <- 0..(length(grid) - 1) do
        for i <- 0..(length(grid) - start - 1), do: Enum.at(Enum.at(grid, start + i), i)
      end

    diagonals_bottom_left =
      for start <- 1..(length(Enum.at(grid, 0)) - 1) do
        for i <- 0..(length(grid) - start - 1), do: Enum.at(Enum.at(grid, i), start + i)
      end

    # Negative-slope diagonals
    diagonals_top_right =
      for start <- 0..(length(Enum.at(grid, 0)) - 1) do
        for i <- 0..min(start, length(grid) - 1), do: Enum.at(Enum.at(grid, i), start - i)
      end

    diagonals_bottom_right =
      for start <- 1..(length(grid) - 1) do
        for i <- 0..min(length(Enum.at(grid, 0)) - 1, length(grid) - start - 1),
            do: Enum.at(Enum.at(grid, start + i), length(Enum.at(grid, 0)) - i - 1)
      end

    all_diagonals =
      diagonals_top_left ++ diagonals_bottom_left ++ diagonals_top_right ++ diagonals_bottom_right

    diagonal_count =
      Enum.reduce(all_diagonals, 0, fn row, acc -> acc + count_xmas(row) end)

    horizontal + vertical + diagonal_count
  end
end
