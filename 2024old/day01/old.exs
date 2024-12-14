defmodule Puzzle1 do
  def run_regex(regex, str, {left, right}) do
    case Regex.run(regex, str) do
      [_, first_number, last_number] ->
        {
          # Prepend to the left list
          [String.to_integer(first_number) | left],
          # Prepend to the right list
          [String.to_integer(last_number) | right]
        }

      _ ->
        # Return the same lists if there's no match
        {left, right}
    end
  end

  def process_file(file_path, regex) do
    # Initialize empty lists
    initial_state = {[], []}

    # Read and process the file
    File.read!(file_path)
    |> String.split("\n", trim: true)
    |> Enum.reduce(initial_state, fn line, acc ->
      run_regex(regex, line, acc)
    end)
    |> then(fn {left, right} ->
      # Reverse the lists to maintain original order since we prepended elements
      {Enum.reverse(left), Enum.reverse(right)}
    end)
  end
end

regex = ~r/^(\d+)\s+(\d+)$/
{left_list, right_list} = Puzzle1.process_file("input.txt", regex)
initial = 0

result =
  Enum.zip(
    Enum.sort(left_list),
    Enum.sort(right_list)
  )
  |> Enum.reduce(initial, fn {l, r}, acc ->
    acc + abs(l - r)
  end)

IO.inspect(result, label: "Total Distance")
# 2_815_556
