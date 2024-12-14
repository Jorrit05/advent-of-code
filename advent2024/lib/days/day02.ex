defmodule Advent2024.Days.Day02 do
  @moduledoc """
  Solution for Day 02 of Advent of Code 2024.
  """

  alias Advent2024.Utils

  def get_input(type \\ :sample) do
    Utils.get_integer_matrix(Utils.get_input(type, __MODULE__))
  end

  def process_report([_]) do
    true
  end

  def process_report([fst, snd | tail]) do
    case abs(fst - snd) >= 1 && abs(fst - snd) <= 3 do
      true ->
        process_report([snd | tail])

      _ ->
        false
    end
  end

  @spec get_valid_reports(list(list(integer()))) :: {integer(), list(list(integer()))}
  def get_valid_reports(lists) do
    lists
    |> Enum.reduce({0, []}, fn lst, {counter, unsafe_lists} ->
      cond do
        Utils.ordered?(lst, :increasing) and process_report(lst) ->
          {counter + 1, unsafe_lists}

        Utils.ordered?(lst, :decreasing) and process_report(lst) ->
          {counter + 1, unsafe_lists}

        true ->
          {counter, [lst | unsafe_lists]}
      end
    end)
  end

  @spec delete_list_entry(integer(), list(integer()), {integer(), list(integer())}) ::
          {integer(), list(integer())}
  def delete_list_entry(index, list, {counter, unsafe_lists}) do
    cond do
      length(list) == index ->
        {counter, unsafe_lists}

      Utils.ordered?(List.delete_at(list, index), :increasing) and
          process_report(List.delete_at(list, index)) ->
        {counter + 1, unsafe_lists}

      Utils.ordered?(List.delete_at(list, index), :decreasing) and
          process_report(List.delete_at(list, index)) ->
        {counter + 1, unsafe_lists}

      true ->
        delete_list_entry(index + 1, list, {counter, unsafe_lists})
    end
  end

  @spec get_semi_safe_reports(list(list(integer()))) :: {integer(), list(list(integer()))}
  def get_semi_safe_reports(unsafe_list) do
    unsafe_list
    |> Enum.reduce({0, []}, fn lst, acc ->
      delete_list_entry(0, lst, acc)
    end)
  end
end
