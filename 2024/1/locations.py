class Locations:
    def __init__(self, list_of_ids):
        self.left_list, self.right_list = self.parse_ids(list_of_ids)

    def parse_ids(self, list_of_ids):
        left_list = []
        right_list = []
        for line in list_of_ids:
            left, right = line.split()
            left_list.append(int(left))
            right_list.append(int(right))
        return left_list, right_list

    def get_distances(self):
        """Return a list of distances between the left and right lists."""
        sorted_left = sorted(self.left_list)
        sorted_right = sorted(self.right_list)

        if len(sorted_left) != len(sorted_right):
            raise ValueError("Lists are not the same length")

        distances = []
        for i in range(len(sorted_left)):
            distances.append(abs(sorted_left[i] - sorted_right[i]))

        return distances

    def get_sum_of_distances(self):
        return sum(self.get_distances())

    def get_similarity_score(self):
        """Return a score by summing all the left numbers after
        multiplying each by how often it appears in the right list."""
        right_dict = {}
        for right in self.right_list:
            if right in right_dict:
                right_dict[right] += 1
            else:
                right_dict[right] = 1

        score = 0
        for left in self.left_list:
            score += left * right_dict.get(left, 0)

        return score


if __name__ == "__main__":
    import os
    import sys

    if len(sys.argv) != 2:
        print("Usage: python distances.py <file>")
        sys.exit(1)

    if not os.path.exists(sys.argv[1]):
        print(f"Error: File '{sys.argv[1]}' not found")
        sys.exit(1)

    with open(sys.argv[1]) as f:
        list_of_ids = f.readlines()

    locations = Locations(list_of_ids)

    print("Distances: ", locations.get_sum_of_distances())
    print("Similarity Score: ", locations.get_similarity_score())
