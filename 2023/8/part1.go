package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	Name  string
	Left  *Node
	Right *Node
}

type Movement struct {
	Directions []string
	Modifier   int
}

var nodes map[string]*Node

// Next increments the movement to access the next direction
func (m *Movement) Next() string {
	movement := m.Directions[m.Modifier]
	m.Modifier++
	if m.Modifier >= len(m.Directions) {
		m.Modifier = 0
	}
	return movement
}

func contains(s string, c byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return true
		}
	}
	return false
}

// getOrCreateNode returns an existing node or creates a new one
func getOrCreateNode(name string) *Node {
	if _, ok := nodes[name]; ok {
		return nodes[name]
	}
	return &Node{Name: name}
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

	nodes = make(map[string]*Node)
	instructions := &Movement{}

	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		if !contains(line, '=') {
			// parse the directions
			instructions.Directions = strings.Split(strings.TrimSpace(line), "")
			continue
		} else {
			// these are all the node definitions AAA = (BBB, CCC)
			// split on the equals sign
			parts := strings.Split(line, "=")

			// the first part is the node name
			name := strings.TrimSpace(parts[0])
			// check the nodes for an existing node
			node := getOrCreateNode(name)

			// the second part is the children
			children := strings.TrimSpace(parts[1])
			// remove the parens
			children = strings.Trim(children, "()")
			// split on the comma
			childNames := strings.Split(children, ",")
			// add it to the map
			nodes[name] = node
			// if there are children, add them to the node
			if len(childNames) > 0 {
				childName := strings.TrimSpace(childNames[0])
				childNode := getOrCreateNode(childName)
				node.Left = childNode
				nodes[childName] = node.Left
			}
			if len(childNames) > 1 {
				childName := strings.TrimSpace(childNames[1])
				childNode := getOrCreateNode(childName)
				node.Right = childNode
				nodes[childName] = node.Right
			}
		}
	}

	// follow the instructions until we get to "ZZZ"
	// every "R" gets the "Right" node, every "L" gets the "Left" node
	// if we get to a node with a Name of "ZZZ", we're done
	node := nodes["AAA"] // always start here
	var steps int
	for {
		direction := instructions.Next()
		steps++
		if direction == "R" {
			node = node.Right
		} else {
			node = node.Left
		}
		if node.Name == "ZZZ" {
			break
		}
	}

	fmt.Println("Steps:", steps)
}
