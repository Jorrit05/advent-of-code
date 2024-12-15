defmodule Advent2024.Days.Day05 do
  @moduledoc """
  Solution for Day 05 of Advent of Code 2024.
  """

  require Logger
  alias Advent2024.Utils

  def get_input(type \\ :sample) do
    [rules | updates] =
      Utils.get_raw_file(Utils.get_input(type, __MODULE__))
      |> String.split("\n\n")

    rule_map =
      rules
      |> String.split(["\n"])
      |> Enum.map(fn line ->
        [left, right] = String.split(line, "|")
        {String.to_integer(left), String.to_integer(right)}
      end)

    update_instructions =
      updates
      |> List.first()
      |> String.split(["\n"])
      |> Enum.map(fn line ->
        line
        |> String.split(",")
        |> Enum.map(&String.to_integer/1)
      end)

    {rule_map, update_instructions}
  end

  def process_line(_rule_set, [_]) do
    true
  end

  def process_line(rule_set, [fst, snd | tail]) do
    cond do
      MapSet.member?(rule_set, {snd, fst}) -> false
      true -> process_line(rule_set, [snd | tail])
    end
  end

  defp get_middle_number(line) do
    Enum.at(line, div(length(line), 2))
  end

  def part1(rule_set, update_instructions) do
    update_instructions
    |> Enum.reduce({0, []}, fn line, {counter, incorrect_updates} ->
      cond do
        process_line(rule_set, line) ->
          {counter + get_middle_number(line), incorrect_updates}

        true ->
          {counter, [line] ++ incorrect_updates}
      end
    end)
  end

  def part2(rule_set, update_instructions) do
    update_instructions
    |> Enum.reduce({0, []}, fn line, {counter, incorrect_updates} ->
      cond do
        process_line(rule_set, line) ->
          {counter + get_middle_number(line), incorrect_updates}

        true ->
          {counter, [line] ++ incorrect_updates}
      end
    end)
  end
end
