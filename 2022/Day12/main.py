#!/usr/bin/env python

import time
import os
from collections import deque

class Mountain():
    def __init__(self, file_name: str):
        with open(f"{os.getcwd()}/Day12/{file_name}") as f:
            lines = f.read().splitlines()
            self.mountain = [[char for char in line.split()]
                                   for line in lines]

        self.row_length = len(self.mountain[0])
        self.column_length = len(self.mountain)
        self.range = range(self.row_length)
        self.queue = deque()
        self.current_pos = self.mountain[0][0]

    def print_mountain(self):
        """Print current mountain to stdout"""
        list(map(lambda xs: print(*xs), self.mountain))

    def check_neighbours(r, c):

        if self.mountain[0][0]

    def solve(self):
        """
        """
        visited = []
        self.queue.append(self.mountain[0][0])
        while True:
            if len(self.queue) == 0:
                print("No Solution Found")
                break

            self.sudoku = self._stack.pop()

            try:
                entry = self.get_entry()
            except (IndexError):
                if self.valid_sudoku():
                    self.show_sudoku()
                    return True

            self.add_moves(entry)

if __name__ == "__main__":
    st = time.time()
    grid = Mountain("test.txt")
    grid.print_mountain()

    print('Execution time:', {time.time() - st}, 'seconds')
    os.EX_OK
