package main

import "fmt"

type Node struct {
	Name     string
	Level    int
	IsFolder bool

	Content    []*Node
	ParentNode *Node
}

func (parent *Node) InsertNode(currentNode *Node) {
	if currentNode.IsRoot() {
		parent.IsFolder = true
		parent.Content = nil
		parent.Level = 0
		parent.Name = currentNode.Name
		parent.ParentNode = nil
		return
	}
	if parent.IsParentOf(currentNode) {
		currentNode.ParentNode = parent
		parent.Content = append(parent.Content, currentNode)
		return
	}
	if parent.IsLevelHigherThan(currentNode) {
		lastChild := parent.Content[len(parent.Content)-1]
		lastChild.InsertNode(currentNode)
		return
	}
}

func (n *Node) IsRoot() bool {
	return n.Level == 0
}
func (parent *Node) IsLevelHigherThan(n *Node) bool {
	return parent.Level < n.Level
}

func (parent *Node) IsParentOf(n *Node) bool {
	return parent.Level+1 == n.Level
}

func (n *Node) PrintNode() {
	eol := ""
	if n.IsFolder {
		eol = "/"
	}
	for range n.Level * 2 {
		fmt.Print("_")
	}

	fmt.Println(n.Name + eol)
	if !n.IsFolder || len(n.Content) == 0 {
		return
	}

	for _, child := range n.Content {
		child.PrintNode()
	}
}