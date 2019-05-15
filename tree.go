// Package gotree create and print tree.
package tree

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	middleItem   = "├──"
	continueItem = "│  "
	emptyItem    = "   "
	lastItem     = "└──"
	NullNode     = "<< NULL >>"
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
	Node  interface{}
	nodes []*Tree
}

// New returns a new tree
func New(name string) *Tree {
	return &Tree{
		Node:  name,
		nodes: []*Tree{},
	}
}

// Add node in tree
func (t *Tree) Add(text interface{}) *Tree {
	n := new(Tree)
	n.Node = text
	t.nodes = append(t.nodes, n)
	return n
}

// Walk walking by tree Stringers
func Walk(t *Tree, f func(str interface{})) {
	f(t.Node)
	for i := range t.nodes {
		Walk(t.nodes[i], f)
	}
}

// String return string with tree view
func (t Tree) Print() (out string) {
	return t.printNode(false, []string{})
}

func isNil(i interface{}) bool {
	return i == nil || (reflect.ValueOf(i).Kind() == reflect.Ptr && reflect.ValueOf(i).IsNil())
}

func toString(i interface{}) (name string) {
	switch v := i.(type) {
	case *Tree:
		if !isNil(v) {
			name = toString(v.Node)
		} else {
			name = NullNode
		}

	case Tree:
		name = toString(v.Node)

	case Stringer:
		if !isNil(v) {
			name = v.String()
		} else {
			name = NullNode
		}

	case string:
		if v != "" {
			name = v
		} else {
			name = NullNode
		}

	case error:
		if !isNil(v) {
			name = v.Error()
		} else {
			name = NullNode
		}

	default:
		if i != nil {
			name = fmt.Sprintf("%v", i)
		} else {
			name = NullNode
		}
	}

	return name
}

func (t Tree) printNode(isLast bool, spaces []string) (out string) {
	// clean name from spaces at begin and end of string
	var name string
	name = strings.TrimSpace(toString(t))

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
