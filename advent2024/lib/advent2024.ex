defmodule Advent2024 do
  use Application

  def start(_type, _args) do
    children = [
      {Advent2024.Utils, []}
    ]

    Supervisor.start_link(children, strategy: :one_for_one, name: Advent2024.Supervisor)
  end
end
