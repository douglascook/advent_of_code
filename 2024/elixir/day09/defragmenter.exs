require Integer

IO.puts("Day 9 - Disk Defrag")

defmodule Day9 do
  def defrag_file(filepath) do
    IO.puts("Processing input #{filepath}")
    defragment(File.read!(filepath))
  end

  def defragment(disk_map) do
    {files, gaps} =
      disk_map
      |> String.trim()
      |> String.graphemes()
      |> Enum.map(&String.to_integer/1)
      |> Enum.with_index()
      |> Enum.split_with(fn {_, i} -> Integer.is_even(i) end)

    gaps = gaps |> Enum.map(&elem(&1, 0))
    files = files |> Enum.map(&elem(&1, 0)) |> Enum.with_index()

    checksum = defragment(files, gaps, 0, 0)
    IO.puts("Checksum = #{checksum}")
  end

  def defragment([], _, _, checksum) do
    checksum
  end

  def defragment(files, gaps, blocks_pointer, checksum) when files != [] do
    # IO.inspect(files, charlists: :as_lists, label: "defrag files")
    # IO.inspect(gaps, charlists: :as_lists, label: "defrag gaps")
    # IO.inspect(blocks_pointer, charlists: :as_lists, label: "defrag blocks pointer")
    # IO.inspect(checksum, charlists: :as_list, label: "defrag checksum")
    [{file_size, file_id} | remaining_files] = files

    checksum = checksum + calculate_checksum_for_file(file_size, file_id, blocks_pointer)
    IO.puts("CHECKSUM after processing file_id #{file_id} = #{checksum}")
    blocks_pointer = blocks_pointer + file_size

    if gaps != [] do
      [gap_size | remaining_gaps] = gaps

      {remaining_files, blocks_pointer, checksum} =
        increment_checksum_for_gap(gap_size, remaining_files, blocks_pointer, checksum)

      IO.puts("CHECKSUM after processing gap (after file_id #{file_id}) = #{checksum}")

      defragment(
        remaining_files,
        remaining_gaps,
        blocks_pointer,
        checksum
      )
    else
      defragment(
        remaining_files,
        gaps,
        blocks_pointer,
        checksum
      )
    end
  end

  def calculate_checksum_for_file(file_size, file_id, blocks_pointer) do
    # IO.puts("CALCULATING CHECKSUM #{file_size}, #{file_id}, #{blocks_pointer}")
    Enum.to_list(blocks_pointer..(blocks_pointer + file_size - 1))
    |> Enum.map(fn pos -> file_id * pos end)
    |> Enum.sum()
  end

  def increment_checksum_for_gap(0, files, blocks_pointer, checksum) do
    {files, blocks_pointer, checksum}
  end

  def increment_checksum_for_gap(_, [], blocks_pointer, checksum) do
    {[], blocks_pointer, checksum}
  end

  def increment_checksum_for_gap(gap_size, files, blocks_pointer, checksum) do
    # IO.puts("RECURSE")
    # IO.inspect(gap_size, charlists: :as_lists, label: "------gap_size")
    # IO.inspect(files, charlists: :as_lists, label: "------files")
    # IO.inspect(blocks_pointer, charlists: :as_lists, label: "------blocks_pointer")
    # IO.inspect(checksum, charlists: :as_lists, label: "------checksum")

    # IO.puts("Gap size before taking another block = #{gap_size}")
    {file_id, files} = take_last_block(files)
    # IO.inspect(file_id, charlists: :as_lists, label: "------file_id of last block")
    # IO.inspect(files, charlists: :as_lists, label: "------after taking last block")

    increment_checksum_for_gap(
      gap_size - 1,
      files,
      blocks_pointer + 1,
      checksum + blocks_pointer * file_id
    )
  end

  def take_last_block(files) do
    [last_file | others] = Enum.reverse(files)
    others = Enum.reverse(others)

    {file_size, file_id} = last_file
    file_size = file_size - 1

    if file_size == 0 do
      {file_id, others}
    else
      {file_id, others ++ [{file_size, file_id}]}
    end
  end
end

Day9.defragment("12345")
Day9.defragment("2333133121414131402")
Day9.defrag_file("../../inputs/day09/input.txt")
