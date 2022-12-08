def find_marker(text: str, uniques: int = 4) -> int:
    chars = []
    for idx, char in enumerate(text):
        if len(chars) < uniques:
            chars.append(char)
        else:
            if len(set(chars)) < len(chars):
                chars.pop(0)
                chars.append(char)
            else:
                return idx


if __name__ == "__main__":
    import sys

    with open(sys.argv[1], 'r') as f:
        target = f.read().strip()
        index = find_marker(target)

        print(index)

        message_idx = find_marker(target, 14)

        print(message_idx)
