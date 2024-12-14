alias Advent2024.Days.Day05

args = System.argv()

# Determine the input type based on the argument (default to :sample)
input_type =
  case args do
    ["input"] -> :input
    _ -> :sample
  end

data = Day05.get_input(input_type)

Day05.part1(data) |> IO.inspect(label: "Part 1: ")
