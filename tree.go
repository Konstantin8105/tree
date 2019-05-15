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
	nodes []interface{}
}

// New returns a new tree
func New(node interface{}) (tr *Tree) {
	if tr, ok := node.(Tree); ok {
		return &tr
	}
	tr = new(Tree)
	tr.nodes = append(tr.nodes, node)
	return tr
}

// Add node in tree
func (t *Tree) Add(node interface{}) *Tree {
	if tr, ok := node.(Tree); ok {
		node = &tr
	}
	if tr, ok := node.(*Tree); ok {
		t.nodes = append(t.nodes, tr)
		return tr
	}
	n := new(Tree)
	n.nodes = append(n.nodes, node)
	t.nodes = append(t.nodes, n)
	return n
}

// Walk walking by tree Stringers
func Walk(t *Tree, f func(str interface{})) {
	if t == (*Tree)(nil) {
		return
	}
	for i := range t.nodes {
		if tr, ok := t.nodes[i].(*Tree); ok && tr != (*Tree)(nil) {
			Walk(tr, f)
			continue
		}
		f(t.nodes[i])
	}
}

// String return string with tree view
func (t Tree) Print() (out string) {
	return t.printNode(false, []string{})
}

func isNil(i interface{}) bool {
	return i == nil || (reflect.ValueOf(i).Kind() == reflect.Ptr && reflect.ValueOf(i).IsNil())
}

func toString(i interface{}) (out string) {
	out = NullNode
	if isNil(i) {
		return
	}

	switch v := i.(type) {
	case *Tree:
		if len(v.nodes) > 0 {
			out = toString(v.nodes[0])
		}

	case Tree:
		panic("not acceptable use Tree. Use *Tree insteand of Tree")

	case Stringer:
		if !isNil(v) {
			out = v.String()
		}

	case string:
		if v != "" {
			out = v
		}

	case error:
		if !isNil(v) {
			out = v.Error()
		}

	default:
		if i != nil {
			out = fmt.Sprintf("%v", i)
		}
	}

	return out
}

func (t *Tree) printNode(isLast bool, spaces []string) (out string) {
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

	if t != (*Tree)(nil) {
		for i := 0; i < len(t.nodes); i++ {
			node := (t.nodes[i])
			if tr, ok := node.(*Tree); ok {
				out += tr.printNode(i == len(t.nodes)-1, spaces)
			}
		}
	}

	return
}
