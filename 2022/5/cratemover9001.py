import re


class Crate:
    def __init__(self, label_text: str):
        self.label = self.clean_label(label_text)

    def __str__(self):
        return f"[{self.label}]"

    def __repr__(self):
        return str(self)

    def clean_label(self, label_text: str) -> str:
        label = ''.join(x for x in label_text if x.isalpha())
        return label


class CrateStacks:
    def __init__(self, raw_stacks: list):
        self.stacks = self.create_stacks(raw_stacks)

    def create_stacks(self, raw_stacks: list) -> list:
        number_of_stacks = int(raw_stacks[-1].split()[-1])
        stacks = []

        for row_idx in range(len(raw_stacks) - 1):
            stack_row = []

            for stack_num in range(number_of_stacks):
                stack_idx = stack_num * 4
                crate_text = raw_stacks[row_idx][stack_idx:stack_idx + 3]
                crate = ""
                if crate_text.strip():
                    crate = Crate(crate_text)
                stack_row.append(crate)

            stacks.append(stack_row)

        return stacks

    def do_instructions(self, instructions: list):
        move_re = re.compile(r"move (?P<amount>[\d]+) from (?P<from>[\d]+) to (?P<to>[\d]+)")
        for instruction in instructions:
            move = move_re.match(instruction)
            if move:
                amt = int(move.group("amount"))
                src = int(move.group("from"))
                dst = int(move.group("to"))
                self.move_crates(amt, src, dst)

    def move_crates(self, amt: int, src: int, dst: int):
        crates = []
        col = src - 1
        for row_idx, row in enumerate(self.stacks):
            # move down stack until we're not empty
            if not row[col]:
                continue

            # remove crates
            for i in range(amt):
                crates.append(self.stacks[row_idx][col])
                self.stacks[row_idx][col] = ""
                row_idx += 1
            break

        col = dst - 1
        dst_row = None
        for row_idx in reversed(range(len(self.stacks))):
            # move up from bottom of stack until we're not empty
            if self.stacks[row_idx][col]:
                continue

            dst_row = row_idx
            break

        for i in range(amt):
            if dst_row is None:  # we need to add a row
                new_row = ["" for i in self.stacks[0]]
                new_row[col] = crates.pop()
                self.stacks.insert(0, new_row)
            else:
                self.stacks[dst_row][col] = crates.pop()
                dst_row -= 1
                if dst_row < 0:
                    dst_row = None

    def get_top_crates(self):
        tops = []
        for col in range(len(self.stacks[0])):
            for row_idx in range(len(self.stacks)):
                crate = self.stacks[row_idx][col]
                if crate:
                    tops.append(crate)
                    break

        return "".join([c.label for c in tops if c])


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
                instructions.append(line.strip())
            else:
                raw_stacks.append(line)

    stacks = CrateStacks(raw_stacks)
    stacks.do_instructions(instructions)
    print(stacks.get_top_crates())
