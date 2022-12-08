class Elf:

    def __init__(self, calories):
        self.calories = calories
        self.total_calories = sum(self.calories)


class CalorieReader:

    def __init__(self, filename):
        self._input = open(filename, 'r')
        self.raw_calories = [line.strip() for line in self._input.readlines()]

    def group_elves(self):
        calories = []
        for line in self.raw_calories:
            if line == '':
                yield Elf(calories)
                calories = []
            else:
                calories.append(int(line))

        if calories:
            yield Elf(calories)  # one last elf


if __name__ == "__main__":
    cr = CalorieReader('input.txt')

    highest = []
    for elf in cr.group_elves():
        min_calories = min(highest) if highest else 0
        if elf.total_calories > min_calories:
            if len(highest) == 3:
                highest.pop(highest.index(min_calories))
            highest.append(elf.total_calories)

    print(sum(highest))
