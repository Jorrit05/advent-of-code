alias Advent2024.Days.Day06

args = System.argv()

# Determine the input type based on the argument (default to :sample)
input_type =
  case args do
    ["input"] -> :input
    _ -> :sample
  end

floor_map = Day06.get_input(input_type)

# {results, incorrect_updates} =
Day06.part1(floor_map)
