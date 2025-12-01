alias Advent2024.Days.Day01_2021

args = System.argv()

# Determine the input type based on the argument (default to :sample)
input_type =
  case args do
    ["input"] -> :input
    _ -> :sample
  end

data = Day01_2021.get_input(input_type)
IO.inspect(data, label: "Data")

# |> Enum.reduce({0, []}, fn line, {counter, incorrect_updates} ->

Day01_2021.puzzle1(data)
|> IO.inspect(label: "Number of increases")

# |> Enum.reduce(lst, acc, fn x, acc -> if x > acc, do: acc + 1, else: acc end)

# |> IO.inspect(label: "Number of increases")
# Day01.puzzle2(columns)
# |> IO.inspect(label: "Similarity score")
