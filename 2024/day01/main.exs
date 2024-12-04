defmodule Puzzle1 do
  def total_distance(file_path) do
    file_path
    |> File.stream!()
    |> Stream.map(&String.split(&1))
    |> Stream.map(fn [l, r] -> {String.to_integer(l), String.to_integer(r)} end)
    # Before unzip: [{3, 4}, {4, 3}, {2, 5}, {1, 3}, {3, 9}, {3, 3}]
    |> Enum.unzip()
    # After unzip: {[3, 4, 2, 1, 3, 3], [4, 3, 5, 3, 9, 3]}
    |> then(fn {left, right} ->
      Enum.zip(Enum.sort(left), Enum.sort(right))
    end)
    |> Enum.reduce(0, fn {l, r}, acc -> acc + abs(l - r) end)
  end
end

result = Puzzle1.total_distance("input.txt")
IO.inspect(result, label: "Total Distance")
