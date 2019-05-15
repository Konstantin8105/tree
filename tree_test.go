package tree_test

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/Konstantin8105/tree"
)

func ExampleTree() {
	artist := tree.New("Pantera")
	album := artist.Add("Far Beyond Driven")
	album.Add("5 minutes Alone")
	album.Add("Some another")
	artist.Add("Power Metal")
	fmt.Println(artist.Print())

	// Output:
	// Pantera
	// ├──Far Beyond Driven
	// │  ├──5 minutes Alone
	// │  └──Some another
	// └──Power Metal
}

func ExampleTreeMultiline() {
	const name string = "Дерево"
	artist := tree.Tree{}
	artist.Node = name
	album := artist.Add("Поддерево\nс многострочным\nтекстом")
	album.Add("Лист поддерева\nзеленый")
	album.Add("Лист красный")
	artist.Add("Лист\nжелтый")
	fmt.Println(artist.Print())

	// Output:
	// Дерево
	// ├──Поддерево
	// │  с многострочным
	// │  текстом
	// │  ├──Лист поддерева
	// │  │  зеленый
	// │  └──Лист красный
	// └──Лист
	//    желтый
}

func ExampleSubTree() {
	tr := tree.New("     Main tree")
	tr.Add("Node 1\nof main tree\n\n\n")
	tr.Add("          Node 2 of main tree          ")

	subTr := tree.New("Sub tree")
	subTr.Add("Node 1 of sub tree")
	node := subTr.Add("Node 2 of sub tree")

	subsubTr := tree.New("Sub tree")
	subsubTr.Add("Node 1 of sub tree")
	subsubTr.Add("Node 2\nof sub tree")

	in := node.Add("\n\n\nIntermediant node")
	in.Add("some node")
	in.Add(subsubTr)

	node.Add("")  // empty name
	node.Add("B") // small name
	node.Add(subsubTr)

	tr.Add(subTr)

	ln := tr.Add("Last main node")

	var b bytes.Buffer
	b.WriteString("Some string from buffer\nwith multilines")
	ln.Add(&b)

	fmt.Println(tr.Print())

	// Output:
	// Main tree
	// ├──Node 1
	// │  of main tree
	// ├──Node 2 of main tree
	// ├──Sub tree
	// │  ├──Node 1 of sub tree
	// │  └──Node 2 of sub tree
	// │     ├──Intermediant node
	// │     │  ├──some node
	// │     │  └──Sub tree
	// │     │     ├──Node 1 of sub tree
	// │     │     └──Node 2
	// │     │        of sub tree
	// │     ├──
	// │     ├──B
	// │     └──Sub tree
	// │        ├──Node 1 of sub tree
	// │        └──Node 2
	// │           of sub tree
	// └──Last main node
	//    └──Some string from buffer
	//       with multilines
}

func ExampleEmptyTree() {
	tr := tree.Tree{}
	fmt.Println(tr.Print())

	// Output:
}

func ExampleEmptySubTree() {
	var (
		tr  = tree.Tree{}
		str = tree.Tree{}
	)
	tr.Node = nil
	str.Add(nil)
	str.Add("")
	str.Add(nil)
	str.Add(nil)
	tr.Add(&str)
	tr.Add(nil)
	tr.Add(nil)
	tr.Add((*tree.Tree)(nil))
	tr.Add((*TempStruct)(nil))
	fmt.Println(tr.Print())

	// Output:
	// ├──
	// │  ├──<< NULL >>
	// │  ├──
	// │  ├──<< NULL >>
	// │  └──<< NULL >>
	// ├──<< NULL >>
	// ├──<< NULL >>
	// └──<< NULL >>
}

type TempStruct struct {
	a, b int
	f    float32
}

func ExampleWalk() {
	tr := tree.New("     Main tree")
	tr.Add("Node 1\nof main tree\n\n\n")
	tr.Add("          Node 2 of main tree          ")

	// bytes buffer
	var b bytes.Buffer
	b.WriteString("Some string from buffer\nwith multilines")
	tr.Add(&b)

	// error
	err := fmt.Errorf("Some error")
	tr.Add(err)

	// struct
	ts := TempStruct{
		a: 42,
		b: 23,
		f: math.Pi,
	}
	tr.Add(ts)

	subTr := tree.New("Sub tree")
	subTr.Add("Node 1 of sub tree")

	tr.Add(subTr)

	fmt.Fprintf(os.Stdout, "%s", tr.Print())

	tree.Walk(tr, func(str interface{}) {
		name := fmt.Sprintf("%v", str)
		name = strings.TrimSpace(name)
		name = strings.ReplaceAll(name, "\n", " << BreakLine >> ")
		fmt.Fprintf(os.Stdout, "Node: %-20s %s\n", fmt.Sprintf("%T", str), name)
	})

	// Output:
	// Main tree
	// ├──Node 1
	// │  of main tree
	// ├──Node 2 of main tree
	// ├──Some string from buffer
	// │  with multilines
	// └──Sub tree
	//    └──Node 1 of sub tree
	//
	// Node: tree.Line            Main tree
	// Node: tree.Line            Node 1 << BreakLine >> of main tree
	// Node: tree.Line            Node 2 of main tree
	// Node: *bytes.Buffer        Some string from buffer << BreakLine >> with multilines
	// Node: tree.Line            Sub tree
	// Node: tree.Line            Node 1 of sub tree
}
