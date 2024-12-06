def parse_reports(lines):
    reports = []
    for line in lines:
        row = []
        for item in line.split():
            try:
                row.append(int(item))
            except ValueError:
                print(f"Error: Could not convert '{item}' to an integer")
        if row:
            reports.append(row)
    return reports


def get_report_safety(report):
    prev = None
    direction_check = None

    for level in report:
        if prev is None:
            prev = level
            diff = level
            continue

        # get difference for each level
        diff = level - prev if prev is not None else level

        # check if difference is either increasing or decreasing by 1, 2, or 3
        if abs(diff) not in [1, 2, 3]:
            return False

        # check if level if increasing or decreasing, but not both
        direction = diff > 0
        if direction_check is not None:
            if direction != direction_check:  # we changed directions
                return False
        direction_check = direction

        prev = level

    # got this far, we're good
    return True


if __name__ == "__main__":
    import os
    import sys

    if len(sys.argv) != 2:
        print(f"Usage: python {sys.argv[0]} <file>")
        sys.exit(1)

    if not os.path.exists(sys.argv[1]):
        print(f"Error: File '{sys.argv[1]}' not found")
        sys.exit(1)

    with open(sys.argv[1]) as f:
        lines = f.readlines()

    reports = parse_reports(lines)

    safe_reports = sum([1 for report in reports if get_report_safety(report)])
    print(f"Safe reports: {safe_reports}")
