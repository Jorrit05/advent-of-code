memory = File.read!("input.txt")
regex = ~r/mul\((\d{1,3}),(\d{1,3})\)/

results =
  Regex.scan(regex, memory)
  |> Enum.reduce(0, fn [_, a, b], acc ->
    acc + String.to_integer(a) * String.to_integer(b)
  end)

IO.inspect(results, label: "Results")

regex = ~r/(mul\(\d{1,3},\d{1,3}\))|((?:do|don't)\(\))/
matches = Regex.scan(regex, memory)

defmodule Puzzle2 do
  def extract_digits(mul_string) do
    case Regex.run(~r/mul\((\d+),(\d+)\)/, mul_string) do
      [_, a, b] ->
        String.to_integer(a) * String.to_integer(b)

      _ ->
        0
    end
  end

  def resolve_matches([_, "", "don't()"], %{process: true, sum: sum}) do
    %{process: false, sum: sum}
  end

  def resolve_matches([_, "", "do()"], %{process: false, sum: sum}) do
    %{process: true, sum: sum}
  end

  def resolve_matches([_, mul_string], %{process: true, sum: sum}) do
    %{process: true, sum: sum + extract_digits(mul_string)}
  end

  def resolve_matches(_, acc) do
    acc
  end
end

puzzle2 =
  Enum.reduce(matches, %{process: true, sum: 0}, fn match, acc ->
    Puzzle2.resolve_matches(match, acc)
  end)

IO.inspect(puzzle2, label: "Puzzle 2 Result")
