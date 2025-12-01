defmodule Advent2025.ListUtils do
  def ordered?(list, type) do
    case type do
      :increasing ->
        list
        |> Enum.reduce_while(true, fn
          current, true -> {:cont, current}
          current, previous when current > previous -> {:cont, current}
          _, _ -> {:halt, false}
        end) != false

      :decreasing ->
        list
        |> Enum.reduce_while(true, fn
          current, true -> {:cont, current}
          current, previous when current < previous -> {:cont, current}
          _, _ -> {:halt, false}
        end) != false
    end
  end

  def transpose(list) do
    Enum.zip_with(list, &Function.identity/1)
  end

  @spec reverse_matrix(list(list(any()))) :: list(list(any()))
  def reverse_matrix(matrix) do
    matrix
    |> Enum.map(&Enum.reverse(&1))
  end

  def get_matrix_transposed_pos({row, col} = _old_pos) do
    {col, row}
  end

  def get_matrix_transposed_reversed_pos(len, {row, col} = _old_pos) do
    new_row = len - 1 - row
    {col, new_row}
  end
end
