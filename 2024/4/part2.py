#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import sys


def count_xmas_words(grid):
    """For every 'A' in the grid, check if two 'MAS' words intersect diagonally."""
    count = 0
    for i in range(len(grid)):
        for j in range(len(grid[i])):
            if i == 0 or j == 0 or i == len(grid) - 1 or j == len(grid[i]) - 1:
                continue

            if grid[i][j] == 'A':  # check diagonally
                hits = 0

                # check MAS top-left to bottom-right
                if grid[i-1][j-1] == 'M' and grid[i+1][j+1] == 'S':
                    hits += 1

                # check MAS top-right to bottom-left
                if grid[i-1][j+1] == 'M' and grid[i+1][j-1] == 'S':
                    hits += 1

                # check SAM top-left to bottom-right
                if grid[i-1][j-1] == 'S' and grid[i+1][j+1] == 'M':
                    hits += 1

                # check SAM top-right to bottom-left
                if grid[i-1][j+1] == 'S' and grid[i+1][j-1] == 'M':
                    hits += 1

                if hits == 2:
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

    count = count_xmas_words(grid)
    print(count)
