defmodule Puzzle1 do
  def get_lists(file_path) do
    file_path
    |> File.stream!()
    |> Stream.map(&String.split(&1))
    |> Enum.map(fn list -> Enum.map(list, fn val -> String.to_integer(val) end) end)
  end

  def ordered?(list, :increasing) do
    list
    |> Enum.reduce_while(true, fn
      current, true -> {:cont, current}
      current, previous when current > previous -> {:cont, current}
      _, _ -> {:halt, false}
    end) != false
  end

  def ordered?(list, :decreasing) do
    list
    |> Enum.reduce_while(true, fn
      current, true -> {:cont, current}
      current, previous when current < previous -> {:cont, current}
      _, _ -> {:halt, false}
    end) != false
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
        ordered?(lst, :increasing) and process_report(lst) ->
          {counter + 1, unsafe_lists}

        ordered?(lst, :decreasing) and process_report(lst) ->
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

      ordered?(List.delete_at(list, index), :increasing) and
          process_report(List.delete_at(list, index)) ->
        {counter + 1, unsafe_lists}

      ordered?(List.delete_at(list, index), :decreasing) and
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

lists = Puzzle1.get_lists("input.txt")
start_time = System.monotonic_time(:millisecond)

{nr_of_safe_lists, unsafe_lists} =
  Puzzle1.get_valid_reports(lists)

nr_of_safe_lists |> IO.inspect(label: "Total safe reports")

{nr_of_semi_safe_reports, _} =
  Puzzle1.get_semi_safe_reports(unsafe_lists)

IO.inspect(nr_of_semi_safe_reports + nr_of_safe_lists,
  label: "Total safe reports including semi-safe"
)

end_time = System.monotonic_time(:millisecond)
IO.puts("Execution time: #{end_time - start_time} milliseconds")

# -----------------------------------------------------------------------------------------------
# ChatGPT solution:
# Nice use of Enum.any? and more function abstraction
# -----------------------------------------------------------------------------------------------

defmodule ChatGpt do
  def get_lists(file_path) do
    file_path
    |> File.stream!()
    |> Stream.map(&String.split(&1))
    |> Enum.map(fn list -> Enum.map(list, fn val -> String.to_integer(val) end) end)
  end

  def ordered?(list, :increasing) do
    Enum.reduce_while(list, nil, fn
      current, nil -> {:cont, current}
      current, previous when current > previous -> {:cont, current}
      _, _ -> {:halt, false}
    end) != false
  end

  def ordered?(list, :decreasing) do
    Enum.reduce_while(list, nil, fn
      current, nil -> {:cont, current}
      current, previous when current < previous -> {:cont, current}
      _, _ -> {:halt, false}
    end) != false
  end

  defp valid_list?(list) do
    (ordered?(list, :increasing) or ordered?(list, :decreasing)) and process_report(list)
  end

  def process_report([_]), do: true

  def process_report([fst, snd | tail]) do
    case abs(fst - snd) in 1..3 do
      true -> process_report([snd | tail])
      _ -> false
    end
  end

  @spec get_valid_reports(list(list(integer()))) :: {integer(), list(list(integer()))}
  def get_valid_reports(lists) do
    Enum.reduce(lists, {0, []}, fn list, {counter, unsafe_lists} ->
      if valid_list?(list) do
        {counter + 1, unsafe_lists}
      else
        {counter, [list | unsafe_lists]}
      end
    end)
  end

  @spec get_semi_safe_reports(list(list(integer()))) :: {integer(), list(list(integer()))}
  def get_semi_safe_reports(unsafe_lists) do
    Enum.reduce(unsafe_lists, {0, []}, fn list, {counter, remaining_unsafe} ->
      semi_safe_found =
        Enum.any?(0..(length(list) - 1), fn idx ->
          list
          |> List.delete_at(idx)
          |> valid_list?()
        end)

      if semi_safe_found do
        {counter + 1, remaining_unsafe}
      else
        {counter, [list | remaining_unsafe]}
      end
    end)
  end
end

start_time = System.monotonic_time(:millisecond)

IO.inspect("ChatGPT:")
{nr_of_safe_lists, unsafe_lists} = ChatGpt.get_valid_reports(lists)
IO.inspect(nr_of_safe_lists, label: "Total safe reports")

{nr_of_semi_safe_reports, _} = ChatGpt.get_semi_safe_reports(unsafe_lists)

IO.inspect(nr_of_semi_safe_reports + nr_of_safe_lists,
  label: "Total safe reports including semi-safe"
)

end_time = System.monotonic_time(:millisecond)
IO.puts("Execution time: #{end_time - start_time} milliseconds")
