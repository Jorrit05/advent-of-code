alias Advent2024.Days.Day01

args = System.argv()

# Determine the input type based on the argument (default to :sample)
input_type =
  case args do
    ["input"] -> :input
    _ -> :sample
  end

data = Day01.get_input(input_type)

Day01.total_distance(data)
|> IO.inspect(label: "Total Distance")

Day01.puzzle2(columns)
|> IO.inspect(label: "Similarity score")
