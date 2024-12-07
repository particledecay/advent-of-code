import re


MUL_REGEX = re.compile(r'mul\((\d{1,3}),(\d{1,3})\)')
ENABLED_MUL_REGEX = re.compile(
    r'(do\(\)|don\'t\(\)|mul\((\d{1,3}),(\d{1,3})\))')


def extract_sum(line):
    total = 0

    for matched in MUL_REGEX.finditer(line):
        first, second = matched.groups()
        total += int(first) * int(second)

    return total


def extract_enabled_sum(lines):
    total = 0
    enabled = True

    for line in lines:
        for matched in ENABLED_MUL_REGEX.finditer(line):
            instruction = matched.group(1)

            if instruction == "do()":
                enabled = True
            elif instruction == "don't()":
                enabled = False
            elif enabled and instruction.startswith("mul"):
                first, second = matched.group(2), matched.group(3)
                total += int(first) * int(second)

    return total


if __name__ == "__main__":
    import os
    import sys

    if len(sys.argv) < 2:
        print(f"Usage: python3 {sys.argv[0]} <file>")
        sys.exit(1)

    filename = sys.argv[1]

    if not os.path.isfile(filename):
        print(f"{filename} does not exist")
        sys.exit(1)

    total = 0
    enabled_total = 0
    with open(filename) as f:
        lines = f.readlines()

        for i, line in enumerate(lines):
            total += extract_sum(line)

        enabled_total += extract_enabled_sum(lines)

    print(f"Sum: {total}")
    print(f"Enabled Sum: {enabled_total}")
