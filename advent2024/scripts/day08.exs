alias Advent2024.Days.Day08

args = System.argv()

# Determine the input type based on the argument (default to :sample)
input_type =
  case args do
    ["input"] -> :input
    _ -> :sample
  end

data = Day08.get_input(input_type)

# {results, incorrect_updates} =
# Day07.part1(data) |> IO.inspect()
