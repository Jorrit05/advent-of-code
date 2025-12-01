defmodule Advent2024.Days.Day07 do
  @moduledoc """
  Solution for Day 07 of Advent of Code 2024.
  """

  require Logger
  alias Advent2024.Utils

  def get_input(type \\ :sample) do
    Utils.get_raw_file(Utils.get_input(type, __MODULE__))
    |> String.split("\n", trim: true)
    |> Enum.map(&String.split(&1, ":", trim: true))
    |> Enum.map(fn [head, str | _] ->
      rest =
        String.split(str, " ", trim: true)
        |> Enum.map(fn val -> String.to_integer(val) end)

      {String.to_integer(head), rest}
    end)
  end

  def check_combination(target, numbers) when is_list(numbers) do
    operators = [:plus, :mult]
    evaluate_combinations(target, numbers, operators)
  end

  defp evaluate_combinations(target, [h | t], operators) do
    # Start recursion with first number
    evaluate_combinations(target, h, t, operators)
  end

  defp evaluate_combinations(target, result, [], _operators) do
    if result == target, do: target, else: 0
  end

  defp evaluate_combinations(target, current, [h | t], operators) do
    operators
    |> Enum.reduce(0, fn operator, acc ->
      case operator do
        :plus -> acc + evaluate_combinations(target, current + h, t, operators)
        :mult -> acc + evaluate_combinations(target, current * h, t, operators)
      end
    end)
  end

  def part1(data) do
    data
    |> Enum.reduce(0, fn {target, numbers}, acc ->
      acc + check_combination(target, numbers)
    end)
  end
end
