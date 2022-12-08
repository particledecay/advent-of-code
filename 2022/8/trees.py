class Grid:

    def __init__(self, height_map: list):
        self.grid = self.build_grid(height_map)

    def build_grid(self, height_map: list) -> list:
        grid = []

        for line in height_map:
            row = []
            for height in line:
                row.append(int(height))
            grid.append(row)

        return grid

    def is_visible(self, row: int, col: int) -> bool:
        this_height = self.grid[row][col]

        # if at any boundary, tree is visible
        if (row == 0) or (row == len(self.grid) - 1) or \
           (col == 0) or (col == len(self.grid[row]) - 1):
            return True

        visibility = [True, True, True, True]

        # check left
        left = 0
        i = col
        while i > 0:
            i -= 1
            if self.grid[row][i] >= this_height:
                visibility[left] = False
                break
        if visibility[left]:
            return True

        # check up
        up = 1
        y = row
        while y > 0:
            y -= 1
            if self.grid[y][col] >= this_height:
                visibility[up] = False
                break
        if visibility[up]:
            return True

        # check right
        right = 2
        x = col
        while x < len(self.grid[row]) - 1:
            x += 1
            if self.grid[row][x] >= this_height:
                visibility[right] = False
                break
        if visibility[right]:
            return True

        # check down
        down = 3
        j = row
        while j < len(self.grid) - 1:
            j += 1
            if self.grid[j][col] >= this_height:
                visibility[down] = False
                break
        if visibility[down]:
            return True

        return False


if __name__ == "__main__":
    import sys

    with open(sys.argv[1], 'r') as f:
        height_map = [line.strip() for line in f.readlines()]
        tree_map = Grid(height_map)
        visible_trees = 0

        for idx, row in enumerate(tree_map.grid):
            for idx2, tree in enumerate(row):
                if tree_map.is_visible(idx, idx2):
                    visible_trees += 1

        print(f"There are {visible_trees} visible trees.")
