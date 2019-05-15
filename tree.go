// Package gotree create and print tree.
// Based on project: https://github.com/DiSiqueira/GoTree
package tree

import (
	"strings"
)

const (
	newLine      = "\n"
	middleItem   = "├──"
	continueItem = "│  "
	emptyItem    = "   "
	lastItem     = "└──"
)

// Tree struct of tree
type Tree struct {
	Name  string
	nodes []*Tree
}

// New returns a new tree
func New(name string) *Tree {
	return &Tree{
		Name:  name,
		nodes: []*Tree{},
	}
}

// Add node in tree
func (t *Tree) Add(text string) *Tree {
	n := New(text)
	t.nodes = append(t.nodes, n)
	return n
}

// AddTree is add tree in present tree
func (t *Tree) AddTree(at *Tree) {
	t.nodes = append(t.nodes, at)
}

// String return string with tree view
func (t Tree) String() (out string) {
	return t.printNode(false, []string{})
}

func (t Tree) printNode(isLast bool, spaces []string) (out string) {
	// clean name from spaces at begin and end of string
	name := strings.TrimSpace(t.Name)

	// split name into strings lines
	lines := strings.Split(name, "\n")

	var tab [2]string
	for i, level := 0, len(spaces); i < level; i++ {
		if i > 0 {
			tab[0] += spaces[i]
			tab[1] += spaces[i]
		}
		if i == level-1 {
			if isLast {
				tab[0] += lastItem
				tab[1] += emptyItem
			} else {
				tab[0] += middleItem
				tab[1] += continueItem
			}
		}
	}

	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
		if i == 0 {
			out += tab[0] + lines[i]
		} else {
			out += tab[1] + lines[i]
		}
		out += newLine
	}

	size := len(spaces)
	if isLast {
		spaces = append(spaces, emptyItem)
	} else {
		spaces = append(spaces, continueItem)
	}
	defer func() {
		spaces = spaces[:size]
	}()

	for i := 0; i < len(t.nodes); i++ {
		out += t.nodes[i].printNode(i == len(t.nodes)-1, spaces)
	}
	return
}
