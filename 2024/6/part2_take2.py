#!/usr/bin/env python3
# -*- coding: utf-8 -*-
class Cell:
    """Represents a cell in the grid with direction tracking"""

    def __init__(self):
        self.walk_id = 0
        self.blocked = False
        self.walked_north = False
        self.walked_south = False
        self.walked_east = False
        self.walked_west = False


class PathStep:
    """Records a step in the guard's path"""

    def __init__(self, row, col, direction):
        self.row = row
        self.col = col
        self.direction = direction


def record_original_path(grid, start_row, start_col):
    """Record the guard's original path before any modifications"""
    path = []
    row, col = start_row, start_col
    direction = "up"  # Assuming starts facing up

    while True:
        path.append(PathStep(row, col, direction))

        # Get next position
        next_row, next_col = row, col
        if direction == "up":
            next_row -= 1
        elif direction == "down":
            next_row += 1
        elif direction == "right":
            next_col += 1
        elif direction == "left":
            next_col -= 1

        # Check bounds
        if (next_row < 0 or next_row >= len(grid) or
                next_col < 0 or next_col >= len(grid[0])):
            break

        # Check for obstacle
        if grid[next_row][next_col].blocked:
            direction = {"up": "right", "right": "down",
                         "down": "left", "left": "up"}[direction]
            continue

        row, col = next_row, next_col

    return path


def maybe_reset_cell(cell, current_walk):
    """Reset cell's walked states if it's a new walk"""
    if cell.walk_id != current_walk:
        cell.walk_id = current_walk
        cell.walked_north = False
        cell.walked_south = False
        cell.walked_east = False
        cell.walked_west = False


def walk_from_position(grid, row, col, direction, current_walk):
    """Simulate walk from a position, return True if a loop is found"""
    while True:
        next_row, next_col = row, col

        if direction == "up":
            while True:
                next_row -= 1
                if next_row < 0:
                    return False
                if grid[next_row][next_col].blocked:
                    next_row += 1
                    cell = grid[next_row][next_col]
                    maybe_reset_cell(cell, current_walk)
                    if cell.walked_north:
                        return True
                    cell.walked_north = True
                    direction = "right"
                    break

        elif direction == "down":
            while True:
                next_row += 1
                if next_row >= len(grid):
                    return False
                if grid[next_row][next_col].blocked:
                    next_row -= 1
                    cell = grid[next_row][next_col]
                    maybe_reset_cell(cell, current_walk)
                    if cell.walked_south:
                        return True
                    cell.walked_south = True
                    direction = "left"
                    break

        elif direction == "right":
            while True:
                next_col += 1
                if next_col >= len(grid[0]):
                    return False
                if grid[next_row][next_col].blocked:
                    next_col -= 1
                    cell = grid[next_row][next_col]
                    maybe_reset_cell(cell, current_walk)
                    if cell.walked_east:
                        return True
                    cell.walked_east = True
                    direction = "down"
                    break

        elif direction == "left":
            while True:
                next_col -= 1
                if next_col < 0:
                    return False
                if grid[next_row][next_col].blocked:
                    next_col += 1
                    cell = grid[next_row][next_col]
                    maybe_reset_cell(cell, current_walk)
                    if cell.walked_west:
                        return True
                    cell.walked_west = True
                    direction = "up"
                    break

        row, col = next_row, next_col


def find_loops(grid_lines):
    # Initialize grid with Cell objects
    height = len(grid_lines)
    width = len(grid_lines[0])
    grid = [[Cell() for _ in range(width)] for _ in range(height)]

    # Find start position and mark obstacles
    start_row = start_col = 0
    for row in range(height):
        for col in range(width):
            if grid_lines[row][col] == '#':
                grid[row][col].blocked = True
            elif grid_lines[row][col] == '^':
                start_row, start_col = row, col

    # Record original path
    path = []
    row, col = start_row, start_col
    direction = "up"  # Assuming starts facing up

    while True:
        path.append(PathStep(row, col, direction))

        # Get next position
        next_row, next_col = row, col
        if direction == "up":
            next_row -= 1
        elif direction == "down":
            next_row += 1
        elif direction == "right":
            next_col += 1
        elif direction == "left":
            next_col -= 1

        # Check bounds
        if (next_row < 0 or next_row >= height or
                next_col < 0 or next_col >= width):
            break

        # Check for obstacle
        if grid[next_row][next_col].blocked:
            direction = {"up": "right", "right": "down",
                         "down": "left", "left": "up"}[direction]
            continue

        row, col = next_row, next_col

    # Try blocking each position in the path
    tested_cells = {(start_row, start_col)}  # Don't block start
    current_walk = 0
    loop_count = 0

    while len(path) > 1:
        prev_step = path.pop(0)
        blocking_step = path[0]

        if (blocking_step.row, blocking_step.col) in tested_cells:
            continue

        tested_cells.add((blocking_step.row, blocking_step.col))

        # Try blocking this cell
        grid[blocking_step.row][blocking_step.col].blocked = True

        # Simulate walk from previous position
        current_walk += 1
        if simulate_walk(grid, prev_step.row, prev_step.col, prev_step.direction, current_walk):
            loop_count += 1

        # Restore cell
        grid[blocking_step.row][blocking_step.col].blocked = False

    return loop_count


def simulate_walk(grid, row, col, direction, current_walk):
    """Simulate walk from a position, return True if a loop is found"""
    while True:
        next_row, next_col = row, col

        if direction == "up":
            while True:
                next_row -= 1
                if next_row < 0:
                    return False
                if grid[next_row][next_col].blocked:
                    next_row += 1
                    cell = grid[next_row][next_col]
                    maybe_reset_cell(cell, current_walk)
                    if cell.walked_north:
                        return True
                    cell.walked_north = True
                    direction = "right"
                    break

        elif direction == "down":
            while True:
                next_row += 1
                if next_row >= len(grid):
                    return False
                if grid[next_row][next_col].blocked:
                    next_row -= 1
                    cell = grid[next_row][next_col]
                    maybe_reset_cell(cell, current_walk)
                    if cell.walked_south:
                        return True
                    cell.walked_south = True
                    direction = "left"
                    break

        elif direction == "right":
            while True:
                next_col += 1
                if next_col >= len(grid[0]):
                    return False
                if grid[next_row][next_col].blocked:
                    next_col -= 1
                    cell = grid[next_row][next_col]
                    maybe_reset_cell(cell, current_walk)
                    if cell.walked_east:
                        return True
                    cell.walked_east = True
                    direction = "down"
                    break

        elif direction == "left":
            while True:
                next_col -= 1
                if next_col < 0:
                    return False
                if grid[next_row][next_col].blocked:
                    next_col += 1
                    cell = grid[next_row][next_col]
                    maybe_reset_cell(cell, current_walk)
                    if cell.walked_west:
                        return True
                    cell.walked_west = True
                    direction = "up"
                    break

        row, col = next_row, next_col


if __name__ == "__main__":
    import sys

    if len(sys.argv) != 2:
        print(f"Usage: python3 {sys.argv[0]} <filename>")
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
                if char in "^v<>":
                    starting_x = len(row) - 1
                    starting_y = len(rows)
            rows.append(row)

    loops = find_loops(rows)
    print(loops)
