defmodule Puzzle1 do
  def get_columns(file_path) do
    file_path
    |> File.stream!()
    |> Stream.map(&String.split(&1))
    |> Stream.map(fn [l, r] -> {String.to_integer(l), String.to_integer(r)} end)
    # Before unzip: [{3, 4}, {4, 3}, {2, 5}, {1, 3}, {3, 9}, {3, 3}]
    |> Enum.unzip()

    # After unzip: {[3, 4, 2, 1, 3, 3], [4, 3, 5, 3, 9, 3]}
  end

  def total_distance({left, right}) do
    Enum.zip(Enum.sort(left), Enum.sort(right))
    |> Enum.reduce(0, fn {l, r}, acc -> acc + abs(l - r) end)
  end

  def puzzle2({left, right}) do
    map_vals =
      right
      |> Enum.reduce(%{}, fn val, acc ->
        Map.update(acc, val, 1, fn existing_value -> existing_value + 1 end)
      end)

    left
    |> Enum.reduce(0, fn val, acc -> acc + val * Map.get(map_vals, val, 0) end)
  end
end

columns = Puzzle1.get_columns("input.txt")

puzzle1 = Puzzle1.total_distance(columns)
IO.inspect(puzzle1, label: "Total Distance")

puzzle2 = Puzzle1.puzzle2(columns)
IO.inspect(puzzle2, label: "Similarity score")

# -----------------------------------------------------------------------------------------------
# ChatGPT solution:
# -----------------------------------------------------------------------------------------------

defmodule ChatGpt do
  def get_columns(file_path) do
    file_path
    |> File.stream!()
    |> Stream.map(fn line ->
      [l, r] = String.split(line)
      {String.to_integer(l), String.to_integer(r)}
    end)
    |> Enum.unzip()
  end

  def total_distance({left, right}) do
    left
    |> Enum.sort()
    |> Enum.zip(Enum.sort(right))
    |> Enum.map(fn {l, r} -> abs(l - r) end)
    |> Enum.sum()
  end

  def puzzle2({left, right}) do
    frequencies = Enum.frequencies(right)

    left
    |> Enum.reduce(0, fn val, acc ->
      acc + val * Map.get(frequencies, val, 0)
    end)
  end
end

column = ChatGpt.get_columns("input.txt")

ChatGpt.total_distance(column)
|> IO.inspect(label: "Total Distance")

ChatGpt.puzzle2(column)
|> IO.inspect(label: "Similarity score")
