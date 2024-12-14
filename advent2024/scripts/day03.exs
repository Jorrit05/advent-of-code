alias Advent2024.Days.Day03

args = System.argv()

# Determine the input type based on the argument (default to :sample)
input_type =
  case args do
    ["input"] -> :input
    _ -> :sample
  end

data = Day03.get_input(input_type)

Day03.part1(data) |> IO.inspect(label: "Part 1: ")

Day03.part2(data) |> IO.inspect(label: "Part 2")
