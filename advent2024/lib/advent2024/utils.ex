defmodule Advent2024.Utils do
  use GenServer

  def start_link(_opts) do
    GenServer.start_link(__MODULE__, %{}, name: __MODULE__)
  end

  def init(state), do: {:ok, state}

  def day(day) do
    day
    |> Module.split()
    |> List.last()
    |> String.downcase()
  end

  def get_input(:sample, day) do
    "priv/#{day(day)}_sample.txt"
  end

  def get_input(:input, day) do
    "priv/#{day(day)}_input.txt"
  end

  def get_integer_matrix(file_path) do
    GenServer.call(__MODULE__, {:integer_matrix, file_path})
  end

  def get_split_strings(file_path) do
    GenServer.call(__MODULE__, {:split_strings, file_path})
  end

  def get_raw_file(file_path) do
    GenServer.call(__MODULE__, {:raw_file, file_path})
  end

  def ordered?(list, :increasing) do
    GenServer.call(__MODULE__, {:increasing, list})
  end

  def ordered?(list, :decreasing) do
    GenServer.call(__MODULE__, {:decreasing, list})
  end

  def handle_call({:integer_matrix, file_path}, _from, state) do
    content =
      file_path
      |> File.stream!()
      |> Stream.map(&String.split(&1))
      |> Enum.map(fn list -> Enum.map(list, fn val -> String.to_integer(val) end) end)

    {:reply, content, state}
  end

  def handle_call({:raw_file, file_path}, _from, state) do
    content =
      file_path
      |> File.read!()

    {:reply, content, state}
  end

  def handle_call({:split_strings, file_path}, _from, state) do
    content =
      file_path
      |> File.stream!()
      |> Stream.map(&String.split/1)
      |> Enum.to_list()

    {:reply, content, state}
  end

  def handle_call({type, list}, _from, state) do
    resp =
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

    {:reply, resp, state}
  end
end
