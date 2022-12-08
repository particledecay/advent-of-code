#!/usr/bin/env python
import os


ALL_DIRS = {}


def get_size(dirname):
    all_files = set()
    for root, dirs, files in os.walk(dirname):
        for name in files:
            size, filename = name.split('-')
            all_files.add(name)

        for name in dirs:
            files = get_size(os.path.join(root, name))
            if files:
                all_files.union(files)

    ALL_DIRS[dirname] = sum([int(name.split("-")[0]) for name in all_files])
    return all_files


if __name__ == "__main__":
    import sys

    get_size(sys.argv[1])

    total_disk = 70000000
    available_disk = total_disk - ALL_DIRS[sys.argv[1]]
    print(f"We have {available_disk} bytes available")
    need_disk = 30000000 - available_disk
    print(f"We need {need_disk} bytes to have enough space")
    best_option = total_disk
    for dirname, dirsize in ALL_DIRS.items():
        if dirsize >= need_disk and dirsize < best_option:
            best_option = dirsize

    print(f"deleted size = {best_option}")

