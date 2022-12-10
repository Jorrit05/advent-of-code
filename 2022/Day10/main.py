#!/usr/bin/env python

import time
import os

cycle_list = [20, 60, 100, 140, 180, 220]

def check_cycle_count(cycle_count : int, X : int):
     if cycle_count in cycle_list:
        return X * cycle_count
     else:
        return 0

def add_pixel(row : list, current_pos : int, visible_window : list):
    if current_pos in visible_window:
        row.append('#')
    else:
        row.append('.')

    return current_pos + 1

def check_row(row : list, current_pos : int):
    if len(row) == 40:
        screen.append(row)
        return ([], 0)
    else:
        return (row, current_pos)

if __name__ == "__main__":
    st = time.time()
    cycle_count = 0
    total = 0
    X = 1
    current_pos = 0
    visible_window = [0,1,2]
    row = []
    screen = []
    with open(f"{os.getcwd()}/Day10/input.txt") as f:
        lines = f.read().splitlines()

        for line in lines:
            (row, current_pos) = check_row(row, current_pos)
            if line == "noop":
                current_pos = add_pixel(row, current_pos, visible_window)
                cycle_count += 1
                total += check_cycle_count(cycle_count, X)
                continue
            else:
                current_pos = add_pixel(row, current_pos, visible_window)
                cycle_count += 1
                total += check_cycle_count(cycle_count, X)

                (row, current_pos) = check_row(row, current_pos)

                current_pos = add_pixel(row, current_pos, visible_window)
                cycle_count += 1
                total += check_cycle_count(cycle_count, X)

                X += int(line.split()[1])
                visible_window = [X-1, X, X+1]

        (row, current_pos) = check_row(row, current_pos)



    print(total)
    for completeRow in screen:
        print("".join(completeRow))

    et = time.time()
    print('Execution time:', {et - st}, 'seconds')
    os.EX_OK
