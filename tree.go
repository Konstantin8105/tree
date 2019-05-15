// Package gotree create and print tree.
// Based on project: https://github.com/DiSiqueira/GoTree
package tree

import (
	"strings"
)

const (
	newLine      = "\n"
	emptySpace   = "    "
	middleItem   = "├── "
	continueItem = "│   "
	emptyItem    = "    "
	lastItem     = "└── "
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
	return t.printNode(0, false, []string{})
}

func (t Tree) printNode(level int, isLast bool, separator []string) (out string) {
	name := t.Name
	{
		// remove last new lines from name
		for i := len(name) - 1; i >= 0; i-- {
			if i == 0 {
				if name[0] == '\n' {
					name = ""
				}
				break
			}
			if name[i] == '\n' {
				name = name[:len(name)-1]
				continue
			}
			break
		}
	}
	lines := strings.Split(name, "\n")
	var tab [2]string
	for i := 0; i < level; i++ {
		if i > 0 {
			tab[0] += continueItem
			tab[1] += continueItem
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
		if i == 0 {
			out += tab[0] + lines[i]
		} else {
			out += tab[1] + lines[i]
		}
		out += newLine
	}
	level++
	for i := 0; i < len(t.nodes); i++ {
		out += t.nodes[i].printNode(level, i == len(t.nodes)-1, separator)
	}
	return
}
