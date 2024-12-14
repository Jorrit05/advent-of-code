defmodule Advent2024.Days.Day03 do
  @moduledoc """
  Solution for Day 03 of Advent of Code 2024.
  """

  alias Advent2024.Utils

  def get_input(type \\ :sample) do
    Utils.get_raw_file(Utils.get_input(type, __MODULE__))
  end

  def part1(file) do
    regex = ~r/mul\((\d{1,3}),(\d{1,3})\)/

    Regex.scan(regex, file)
    |> Enum.reduce(0, fn [_, a, b], acc ->
      acc + String.to_integer(a) * String.to_integer(b)
    end)
  end

  def part2(file) do
    regex = ~r/(mul\(\d{1,3},\d{1,3}\))|((?:do|don't)\(\))/
    matches = Regex.scan(regex, file)

    Enum.reduce(matches, %{process: true, sum: 0}, fn match, acc ->
      resolve_matches(match, acc)
    end)
  end

  def extract_digits(mul_string) do
    case Regex.run(~r/mul\((\d+),(\d+)\)/, mul_string) do
      [_, a, b] ->
        String.to_integer(a) * String.to_integer(b)

      _ ->
        0
    end
  end

  def resolve_matches([_, "", "don't()"], %{process: true, sum: sum}) do
    %{process: false, sum: sum}
  end

  def resolve_matches([_, "", "do()"], %{process: false, sum: sum}) do
    %{process: true, sum: sum}
  end

  def resolve_matches([_, mul_string], %{process: true, sum: sum}) do
    %{process: true, sum: sum + extract_digits(mul_string)}
  end

  def resolve_matches(_, acc) do
    acc
  end
end
