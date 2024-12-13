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

  def get_valid_reports(lists) do
    lists
    |> Enum.reduce(0, fn lst, acc ->
      cond do
        ordered?(lst, :increasing) and process_report(lst) ->
          acc + 1

        ordered?(lst, :decreasing) and process_report(lst) ->
          acc + 1

        true ->
          acc
      end
    end)
  end
end

lists = Puzzle1.get_lists("input.txt")

Puzzle1.get_valid_reports(lists)
|> IO.inspect(label: "Total safe reports")
