defmodule Advent2024.Days.Day08 do
  @moduledoc """
  Solution for Day 08 of Advent of Code 2024.
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
end
