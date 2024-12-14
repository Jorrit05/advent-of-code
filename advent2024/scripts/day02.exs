alias Advent2024.Days.Day02

args = System.argv()

# Determine the input type based on the argument (default to :sample)
input_type =
  case args do
    ["input"] -> :input
    _ -> :sample
  end

data = Day02.get_input(input_type)

{nr_of_safe_lists, unsafe_lists} =
  Day02.get_valid_reports(data)

nr_of_safe_lists |> IO.inspect(label: "Total safe reports")

{nr_of_semi_safe_reports, _} =
  Day02.get_semi_safe_reports(unsafe_lists)

IO.inspect(nr_of_semi_safe_reports + nr_of_safe_lists,
  label: "Total safe reports including semi-safe"
)
