import operator
from functools import reduce


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

    def scenic_score(self, row: int, col: int) -> int:
        this_height = self.grid[row][col]
        scores = [0, 0, 0, 0]

        # check left
        left = 0
        i = col
        while i > 0:
            i -= 1
            if i == 0 or self.grid[row][i] >= this_height:
                scores[left] += 1
                break
            scores[left] += 1

        # check up
        up = 1
        y = row
        while y > 0:
            y -= 1
            if y == 0 or self.grid[y][col] >= this_height:
                scores[up] += 1
                break
            scores[up] += 1

        # check right
        right = 2
        x = col
        right_bound = len(self.grid[row]) - 1
        while x < right_bound:
            x += 1
            if x == right_bound or self.grid[row][x] >= this_height:
                scores[right] += 1
                break
            scores[right] += 1

        # check down
        down = 3
        j = row
        down_bound = len(self.grid) - 1
        while j < down_bound:
            j += 1
            if j == down_bound or self.grid[j][col] >= this_height:
                scores[down] += 1
                break
            scores[down] += 1

        return reduce(operator.mul, scores)


if __name__ == "__main__":
    import sys

    with open(sys.argv[1], 'r') as f:
        height_map = [line.strip() for line in f.readlines()]
        tree_map = Grid(height_map)
        best_score = 0

        for idx, row in enumerate(tree_map.grid):
            for idx2, tree in enumerate(row):
                scenic_score = tree_map.scenic_score(idx, idx2)
                if scenic_score > best_score:
                    best_score = scenic_score

        print(f"The best tree has a scenic score of {best_score}.")
