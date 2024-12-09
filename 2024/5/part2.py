#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import sys


LINE_RULES = []


def is_page_order_valid(pages):
    for prev_page, after_page in LINE_RULES:
        if prev_page in pages and after_page in pages:
            prev_index = pages.index(prev_page)
            after_index = pages.index(after_page)
            if prev_index > after_index:
                return False
    return True


def fix_page_order(pages):
    for prev_page, after_page in LINE_RULES:
        if prev_page in pages and after_page in pages:
            prev_index = pages.index(prev_page)
            after_index = pages.index(after_page)
            if prev_index > after_index:
                pages[prev_index], pages[after_index] = pages[after_index], pages[prev_index]
                return fix_page_order(pages)
                break
    return pages


if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: $0 <input-file>")
        sys.exit(1)

    total = 0
    with open(sys.argv[1]) as f:
        for line in f:
            line = line.strip()
            if len(line) == 0:
                continue

            if '|' in line:
                prev_page, after_page = line.split('|')
                LINE_RULES.append((prev_page, after_page))

            if ',' in line:
                pages = line.split(',')
                if not is_page_order_valid(pages):
                    # find middle page
                    pages = fix_page_order(pages)
                    total += int(pages[len(pages) // 2])

    print(total)
