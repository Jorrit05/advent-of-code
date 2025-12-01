alias Advent2024.Days.Day07

args = System.argv()

# Determine the input type based on the argument (default to :sample)
input_type =
  case args do
    ["input"] -> :input
    _ -> :sample
  end

data = Day07.get_input(input_type)

# {results, incorrect_updates} =
Day07.part1(data) |> IO.inspect()


# defmodule OperatorCombinations do
# end

# # Example usage:
# target = 3267
# numbers = [81, 40, 27]
# # 3267: 81 40 27

# OperatorCombinations.check_combination(target, numbers)
# |> IO.inspect(label: "Can match target?")
