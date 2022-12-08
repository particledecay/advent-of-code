class Crate:
    def __init__(self, label):
        self.label = label


class CrateStacks:
    def __init__(self, raw_stacks: list):
        self.stacks = self.create_stacks(raw_stacks)

    def create_stacks(self, raw_stacks: list) -> list:
        number_of_stacks = raw_stacks[-1].split()[-1]

        return [number_of_stacks]


if __name__ == "__main__":
    import sys

    raw_stacks = []
    instructions = []
    with open(sys.argv[1], 'r') as f:
        is_instruction = False
        for line in f.readlines():
            if line.strip() == "":
                is_instruction = True
            if is_instruction:
                instructions.append(line)
            else:
                raw_stacks.append(line)

    print(f"raw stacks = {raw_stacks}")
    stacks = CrateStacks(raw_stacks)
    print(stacks.stacks)
