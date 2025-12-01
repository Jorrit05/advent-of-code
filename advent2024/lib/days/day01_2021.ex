defmodule Advent2024.Days.Day01_2021 do
  @moduledoc """
  Solution for Day 01 of Advent of Code 2024.
  """

  alias Advent2024.Utils

  def get_input(type \\ :sample) do
    Utils.get_integer_list(Utils.get_input(type, __MODULE__))
  end

  def get_increases([], _last, counter) do
    counter
  end

  def get_increases([h | tail], last, counter) do
    if h <= last do
      get_increases(tail, h, counter)
    else
      get_increases(tail, h, counter + 1)
    end
  end

  def puzzle1([h | tail]) do
    get_increases(tail, h, 0)
  end
end
