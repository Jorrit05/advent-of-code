defmodule Advent2025 do
  use Application

  def start(_type, _args) do
    children = [
      {Advent2025.Utils, []}
    ]

    Supervisor.start_link(children, strategy: :one_for_one, name: Advent2025.Supervisor)
  end
end
