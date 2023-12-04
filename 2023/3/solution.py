#!/usr/bin/env python
# -*- coding: utf-8 -*-

import sys


# 2d grid
GRID = []


class Number:
    def __init__(self):
        self.start_idx = None
        self.end_idx = None
        self.row = None
        self.number = []

    @property
    def value(self):
        return int(''.join(self.number))


def is_symbol(char):
    return not char.isdigit() and char != '.'


def check_around_number(num):
    left_index = num.start_idx - 1
    if left_index < 0:
        left_index = 0
    right_index = num.end_idx + 1
    if right_index >= len(GRID[num.row]):
        right_index = len(GRID[num.row]) - 1

    # check the row above if it exists
    if num.row > 0:
        row_above = GRID[num.row - 1]
        for i in range(left_index, right_index + 1):
            if is_symbol(row_above[i]):
                return True

    # check left of the number
    if left_index > 0:
        if is_symbol(GRID[num.row][left_index]):
            return True

    # check right of the number
    if right_index < len(GRID[num.row]) - 1:
        if is_symbol(GRID[num.row][right_index]):
            return True

    # check the row below if it exists
    if num.row < len(GRID) - 1:
        row_below = GRID[num.row + 1]
        for i in range(left_index, right_index + 1):
            if is_symbol(row_below[i]):
                return True

    return False
        

# iterate over the lines to build a grid of numbers
def check_grid(lines):
    for row in lines:
        new_row = []
        for col in row:
            new_row.append(col)
        GRID.append(new_row)

    total = 0
    for i, row in enumerate(lines):
        in_number = False
        num = Number()
        for j, col in enumerate(row):
            if not col.isdigit():
                if in_number:
                    num.end_idx = j - 1
                in_number = False
                if len(num.number) > 0:
                    valid = check_around_number(num)
                    if valid:
                        # print(f"VALID: {num.value}")
                        total += num.value
                    else:
                        print(f"{num.value} is not valid")
                num = Number()
            else:
                if not in_number:
                    num.start_idx = j
                    num.row = i
                    in_number = True
                num.number.append(col)

    print(f"total is {total}")


if __name__ == '__main__':
    if len(sys.argv) != 2:
        print('Usage: python solution.py <input_file>')
        sys.exit(1)

    with open(sys.argv[1], 'r') as f:
        lines = f.readlines()

    check_grid(lines)
