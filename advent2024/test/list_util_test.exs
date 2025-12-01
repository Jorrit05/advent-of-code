defmodule ListUtilTest do
  use ExUnit.Case
  alias Advent2024.ListUtils

  @matrix [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
  @transposed [[1, 4, 7], [2, 5, 8], [3, 6, 9]]
  @reversed [[7, 4, 1], [8, 5, 2], [9, 6, 3]]

  test "transpose" do
    assert ListUtils.transpose(@matrix) == @transposed
  end

  test "transpose and reverse" do
    assert ListUtils.reverse_matrix(@transposed) == @reversed
  end

  test "transposed start pos" do
    assert ListUtils.get_matrix_transposed_pos({1, 1}) == {1, 1}
    assert ListUtils.get_matrix_transposed_pos({1, 2}) == {2, 1}
  end

  test "transpose and reverse get start pos" do
    pos = {2, 2}
    assert ListUtils.get_matrix_transposed_reversed_pos(length(@matrix), pos) == {2, 0}
  end
end
