defmodule Puzzle5 do
end

[rules | update] =
  File.read!("sample.txt")
  |> String.split("\n\n")

rule_map =
  rules
  |> String.split(["|", "\n"])

IO.inspect(rule_map)

# Enum.chunk_every(row, 4, 1, :discard)
