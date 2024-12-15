alias Advent2024.Days.Day05

args = System.argv()

# Determine the input type based on the argument (default to :sample)
input_type =
  case args do
    ["input"] -> :input
    _ -> :sample
  end

{rule_map, update_instructions} = Day05.get_input(input_type)

rule_set = MapSet.new(rule_map)

{results, incorrect_updates} =
  Day05.part1(rule_set, update_instructions)

results |> IO.inspect(label: "Part 1: ")

Day05.part2(rule_set, incorrect_updates)
