defmodule Puzzle1 do
  use GenServer

  # Starts the GenServer
  def start_link(_) do
    GenServer.start_link(__MODULE__, %{}, name: __MODULE__)
  end

  # Public API to process the file and store the sorted lists
  def process_and_store(file_path) do
    {left, right} =
      file_path
      |> File.stream!()
      |> Stream.map(&String.split(&1))
      |> Stream.map(fn [l, r] -> {String.to_integer(l), String.to_integer(r)} end)
      |> Enum.unzip()
      |> then(fn {left, right} ->
        {Enum.sort(left), Enum.sort(right)}
      end)

    GenServer.call(__MODULE__, {:store_sorted, {left, right}})
  end

  # Public API to calculate the total distance between the lists
  def calculate_distance() do
    GenServer.call(__MODULE__, :get_sorted_lists)
    |> then(fn {left, right} ->
      Enum.zip(left, right)
      |> Enum.reduce(0, fn {l, r}, acc -> acc + abs(l - r) end)
    end)
  end

  # Public API to retrieve the sorted lists (optional for further puzzles)
  def get_sorted_lists() do
    GenServer.call(__MODULE__, :get_sorted_lists)
  end

  ## GenServer Callbacks

  @impl true
  def init(_) do
    {:ok, %{sorted_lists: nil}}
  end

  @impl true
  def handle_call({:store_sorted, lists}, _from, state) do
    {:reply, :ok, %{state | sorted_lists: lists}}
  end

  @impl true
  def handle_call(:get_sorted_lists, _from, state) do
    {:reply, state.sorted_lists, state}
  end
end
