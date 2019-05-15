// Package gotree create and print tree.
package tree

import (
	"fmt"
	"runtime/debug"
	"strings"
)

const (
	middleItem   = "├──"
	continueItem = "│  "
	emptyItem    = "   "
	lastItem     = "└──"
)

// Line is string type
type Line string

// String return string of Line
func (l Line) String() string {
	return string(l)
}

// Stringer base interface
type Stringer interface {
	String() string
}

// Tree struct of tree
type Tree struct {
	Name  Stringer
	nodes []*Tree
}

// New returns a new tree
func New(name string) *Tree {
	return &Tree{
		Name:  Line(name),
		nodes: []*Tree{},
	}
}

// Add node in tree
func (t *Tree) Add(text Stringer) *Tree {
	n := new(Tree)
	n.Name = text
	t.nodes = append(t.nodes, n)
	return n
}

// AddLine add node with string
func (t *Tree) AddLine(text string) *Tree {
	n := new(Tree)
	n.Name = Line(text)
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
	// panic free
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()

	// clean name from spaces at begin and end of string
	var name string
	if t.Name != nil {
		name = strings.TrimSpace(t.Name.String())
	}

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
		out += "\n"
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
