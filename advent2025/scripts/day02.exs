alias Advent2025.Utils

input_type = Utils.input_type(System.argv())

data =
  Utils.get_raw_file(Utils.get_input(input_type, Day02))
  |> String.split(",")
  |> Enum.map(fn range_str ->
    [start, finish] = String.split(range_str, "-")
    {String.to_integer(start), String.to_integer(finish)}
  end)

defmodule Day2 do
  def analyze_number(number, length) when rem(length, 2) == 0 do
    <<left::binary-size(div(length, 2)), right::binary>> = number
    if left == right, do: String.to_integer(number), else: 0
  end

  def analyze_number(_, _), do: 0

  def difficult_analysis(_, len, substring_len) when substring_len >= len do
    0
  end

  def difficult_analysis(number, len, substring_len) do
    times = div(len, substring_len)
    <<left::binary-size(substring_len), _::binary>> = number

    if String.duplicate(left, times) == number,
      do: String.to_integer(number),
      else: difficult_analysis(number, len, substring_len + 1)
  end

  def solve(data, analyzer_fn) do
    data
    |> Stream.flat_map(fn {start, finish} ->
      Stream.map(start..finish, fn n ->
        str = Integer.to_string(n)
        analyzer_fn.(str, String.length(str))
      end)
    end)
    |> Enum.sum()
  end
end

{time, {one, two}} =
  :timer.tc(fn ->
    {Day2.solve(data, &Day2.analyze_number/2),
     Day2.solve(data, &Day2.difficult_analysis(&1, &2, 1))}
  end)

IO.puts("Puzzle 1: #{one}\nPuzzle 2: #{two}\nExecuted in #{time / 1_000}ms")
