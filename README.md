# advent-of-code

This repository contains my solutions to the [Advent of Code](https://adventofcode.com/) challenges.

## Structure

Each year has its own folder, and each day of the challenge has its own folder within that.

## Useful scripts

### scripts/aoc-fetcher

This script fetches puzzles and inputs while creating the new challenge's directory structure, and generates some boilerplate code.

```bash
# fetches the puzzle and input for today's challenge (assuming we're in an active challenge)
./scripts/aoc-fetcher

# fetches the puzzle and input for day 4 of this year's challenge
./scripts/aoc-fetcher 4

# fetches the puzzle and input for 12/25/2015
./scripts/aoc-fetcher 25 2015
```
