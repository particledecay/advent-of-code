package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Node struct {
	Left  string
	Right string
}

func solve(element string, instructions []rune, network map[string]Node) int {
	steps := 0
	for !strings.HasSuffix(element, "Z") {
		instruction := instructions[steps%len(instructions)]
		if instruction == 'L' {
			element = network[element].Left
		} else {
			element = network[element].Right
		}
		steps++
	}
	return steps
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: $0 <input-file>")
		os.Exit(1)
	}

	// read file from arg
	buf, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer buf.Close()

	network := make(map[string]Node)
	var instructions []rune
	scanner := bufio.NewScanner(buf)

	// Read instructions
	scanner.Scan()
	instructionLine := scanner.Text()
	re := regexp.MustCompile(`([LR]+)`)
	instructions = []rune(re.FindStringSubmatch(instructionLine)[1])

	// Read network data
	nodeRegex := regexp.MustCompile(`^(\w{3}) = \((\w{3}), (\w{3})\)$`)
	for scanner.Scan() {
		line := scanner.Text()
		matches := nodeRegex.FindStringSubmatch(line)
		if matches != nil {
			network[matches[1]] = Node{Left: matches[2], Right: matches[3]}
		}
	}

	// Get solutions and find least common multiple
	var solutions []int
	for node := range network {
		if strings.HasSuffix(node, "A") {
			solutions = append(solutions, solve(node, instructions, network))
		}
	}

	result := solutions[0]
	for _, s := range solutions[1:] {
		result = lcm(result, s)
	}
	fmt.Println(result)
}
