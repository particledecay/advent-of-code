#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import sys


def count_words(grid, word):
    count = 0
    for y in range(len(grid)):
        for x in range(len(grid[y])):
            count += count_words_at_position(grid, word, x, y)
    return count


def check_left(grid, word, x, y):
    if x - (len(word) - 1) >= 0:
        for i in range(len(word)):
            if grid[y][x - i] != word[i]:
                break
        else:
            return True
    return False


def check_right(grid, word, x, y):
    if x + len(word) <= len(grid[y]):
        for i in range(len(word)):
            if grid[y][x + i] != word[i]:
                break
        else:
            return True
    return False


def check_down(grid, word, x, y):
    if y + len(word) <= len(grid):
        for i in range(len(word)):
            if grid[y + i][x] != word[i]:
                break
        else:
            return True
    return False


def check_up(grid, word, x, y):
    if y - (len(word) - 1) >= 0:
        for i in range(len(word)):
            if grid[y - i][x] != word[i]:
                break
        else:
            return True
    return False


def check_up_left(grid, word, x, y):
    if x - (len(word) - 1) >= 0 and y - (len(word) - 1) >= 0:
        for i in range(len(word)):
            if grid[y - i][x - i] != word[i]:
                break
        else:
            return True
    return False


def check_up_right(grid, word, x, y):
    if x + len(word) <= len(grid[y]) and y - (len(word) - 1) >= 0:
        for i in range(len(word)):
            if grid[y - i][x + i] != word[i]:
                break
        else:
            return True
    return False


def check_down_left(grid, word, x, y):
    if x - (len(word) - 1) >= 0 and y + len(word) <= len(grid):
        for i in range(len(word)):
            if grid[y + i][x - i] != word[i]:
                break
        else:
            return True
    return False


def check_down_right(grid, word, x, y):
    if x + len(word) <= len(grid[y]) and y + len(word) <= len(grid):
        for i in range(len(word)):
            if grid[y + i][x + i] != word[i]:
                break
        else:
            return True
    return False


def count_words_at_position(grid, word, x, y):
    count = 0

    # Check left
    if check_left(grid, word, x, y):
        count += 1

    # Check right
    if check_right(grid, word, x, y):
        count += 1

    # Check down
    if check_down(grid, word, x, y):
        count += 1

    # Check up
    if check_up(grid, word, x, y):
        count += 1

    # Check diagonal
    if check_up_left(grid, word, x, y):
        count += 1
    if check_up_right(grid, word, x, y):
        count += 1
    if check_down_left(grid, word, x, y):
        count += 1
    if check_down_right(grid, word, x, y):
        count += 1

    return count


if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: $0 <input-file>")
        sys.exit(1)

    grid = []
    with open(sys.argv[1]) as f:
        for line in f:
            line = line.strip()
            if len(line) == 0:
                continue

            chars = [c for c in line]
            grid.append(chars)

    word = "XMAS"
    count = count_words(grid, word)
    print(count)
