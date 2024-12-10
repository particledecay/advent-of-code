#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import sys


def get_next_position(x, y, direction):
    """Given the current position and direction, return the next position."""
    if direction == "up":
        return x, y - 1
    elif direction == "down":
        return x, y + 1
    elif direction == "left":
        return x - 1, y
    elif direction == "right":
        return x + 1, y


def guard_walk(grid, x, y, direction):
    """Guard starts at x,y and walks in the direction until just before it hits a '#' character,
    then turns right and continues walking. This continues until the guard reaches
    the edge of the grid. Also mark each unique cell visited."""
    path = set()
    while True:
        # mark the cell as visited
        path.add((x, y))

        # check the next cell is walkable
        next_x, next_y = get_next_position(x, y, direction)

        # is the next cell out of bounds?
        if next_y < 0 or next_y >= len(grid) or next_x < 0 or next_x >= len(grid[next_y]):
            break

        # if the next cell is not walkable, turn right
        if grid[next_y][next_x] == "#":
            if direction == "up":
                direction = "right"
            elif direction == "down":
                direction = "left"
            elif direction == "left":
                direction = "up"
            elif direction == "right":
                direction = "down"
            continue

        # move to the next cell
        x, y = next_x, next_y

    # return count of unique cells visited
    return path


if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: $0 <input-file>")
        sys.exit(1)

    rows = []
    direction = starting_x = starting_y = None
    with open(sys.argv[1]) as f:
        for line in f:
            line = line.strip()
            if len(line) == 0:
                continue

            # build the grid
            row = []
            for char in line:
                row.append(char)
                if char == "^":
                    direction = "up"
                elif char == "v":
                    direction = "down"
                elif char == "<":
                    direction = "left"
                elif char == ">":
                    direction = "right"

                # starting guard position
                if char in ["^", "v", "<", ">"]:
                    starting_x = len(row) - 1
                    starting_y = len(rows)
            rows.append(row)

    path = guard_walk(rows, starting_x, starting_y, direction)
    print(len(path))
