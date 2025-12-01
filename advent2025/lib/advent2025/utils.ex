defmodule Advent2025.Utils do
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

  def get_integer_list(file_path) do
    GenServer.call(__MODULE__, {:integer_list, file_path})
  end

  def get_integer_matrix(file_path) do
    GenServer.call(__MODULE__, {:integer_matrix, file_path})
  end

  def get_split_strings(file_path) do
    GenServer.call(__MODULE__, {:split_strings, file_path})
  end

  def get_string_grapheme(file_path) do
    GenServer.call(__MODULE__, {:string_grapheme, file_path})
  end

  def get_character_matrix(file_path) do
    GenServer.call(__MODULE__, {:character_matrix, file_path})
  end

  def get_raw_file(file_path) do
    GenServer.call(__MODULE__, {:raw_file, file_path})
  end

  def handle_call({:integer_matrix, file_path}, _from, state) do
    content =
      file_path
      |> File.stream!()
      |> Stream.map(&String.split(&1))
      |> Enum.map(fn list -> Enum.map(list, fn val -> String.to_integer(val) end) end)

    {:reply, content, state}
  end

  def handle_call({:integer_list, file_path}, _from, state) do
    content =
      file_path
      |> File.stream!()
      |> Stream.map(&String.split(&1))
      |> Enum.map(fn [val] -> String.to_integer(val) end)

    {:reply, content, state}
  end

  def handle_call({:raw_file, file_path}, _from, state) do
    content =
      file_path
      |> File.read!()

    {:reply, content, state}
  end

  def handle_call({:character_matrix, file_path}, _from, state) do
    content =
      file_path
      |> File.read!()
      |> String.split("\n")
      |> Enum.map(&String.graphemes/1)

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

  def handle_call({:string_grapheme, file_path}, _from, state) do
    content =
      file_path
      |> File.stream!()
      |> Stream.map(&String.next_grapheme/1)
      |> Enum.to_list()

    {:reply, content, state}
  end
end
