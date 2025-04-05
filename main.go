package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseLine(line string) *Node {
	// line can't be empty
	if len(line) == 0 {
		return nil
	}
	runes := []Rune(line)
	keyword := []Rune{}

	node := &Node{}
	if runes[len(runes)-1].IsSlash() {

		node.IsFolder = true
	}
	i := 0
	// for space parsing
	for {
		if !runes[i].IsSpace() || i >= len(runes)-1 {
			break
		}
		i++

		node.Level++
	}

	// store the keyword
	for {
		if i == len(runes)-1 && runes[i].IsSlash() {
			break
		}

		keyword = append(keyword, runes[i])
		i++
		if i >= len(runes) {
			break
		}
	}
	node.Level = node.Level / 2
	node.Name = string(keyword)
	return node
}

func main() {
	file, err := os.Open("folder-structure.yaml")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	root := &Node{}

	for scanner.Scan() {
		line := strings.TrimRight(scanner.Text(), " ")
		if len(line) == 0 {
			continue
		}
		node := parseLine(line)
		root.InsertNode(node)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error:", err)
	}
	fmt.Println("RESULT>>>>>>>>>>>>>>")
	root.PrintNode()

}
