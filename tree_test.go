package tree_test

import (
	"bytes"
	"fmt"
	"math"
	"os"

	"github.com/Konstantin8105/tree"
)

func ExampleTree() {
	artist := tree.New("Pantera")
	album := tree.New("Far Beyond Driven")
	album.Add("5 minutes Alone")
	album.Add("Some another")
	artist.Add(album)
	artist.Add("Power Metal")
	fmt.Fprintf(os.Stdout, "%s\n", artist)

	// Output:
	// Pantera
	// ├──Far Beyond Driven
	// │  ├──5 minutes Alone
	// │  └──Some another
	// └──Power Metal
}

func ExampleTreeMultiline() {
	const name string = "Дерево"
	artist := tree.New(name)
	album := tree.New("Поддерево\nс многострочным\nтекстом")
	album.Add("Лист поддерева\nзеленый")
	album.Add("Лист красный")
	artist.Add(album)
	artist.Add("Лист\nжелтый")
	fmt.Fprintf(os.Stdout, "%s\n", artist)

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
	node := tree.New("Node 2 of sub tree")
	subTr.Add(node)

	subsubTr := tree.New("Sub tree")
	subsubTr.Add("Node 1 of sub tree")
	subsubTr.Add("Node 2\nof sub tree")

	in := tree.New("\n\n\nIntermediant node")
	node.Add(in)
	in.Add("some node")
	in.Add(subsubTr)

	node.Add("")  // empty name
	node.Add("B") // small name
	node.Add(subsubTr)

	tr.Add(subTr)

	ln := tree.New("Last main node")
	tr.Add(ln)

	var b bytes.Buffer
	b.WriteString("Some string from buffer\nwith multilines")
	ln.Add(&b)

	fmt.Fprintf(os.Stdout, "%s\n", tr)

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
	// │     ├──<< NULL >>
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
	fmt.Fprintf(os.Stdout, "%s\n", tr)

	// Output:
	// << NULL >>
}

func ExampleEmptySubTree() {
	var (
		tr  = tree.Tree{}
		str = tree.Tree{}
	)
	str.Add(nil)
	str.Add("")
	str.Add((*tree.Tree)(nil))
	str.Add(nil)

	tr.Add(&str)

	tr.Add(nil)
	tr.Add(nil)
	tr.Add((*tree.Tree)(nil))
	tr.Add((*TempStruct)(nil))

	fmt.Fprintf(os.Stdout, "%s\n", tr)

	// Output:
	// << NULL >>
	// ├──<< NULL >>
	// │  ├──<< NULL >>
	// │  ├──<< NULL >>
	// │  ├──<< NULL >>
	// │  └──<< NULL >>
	// ├──<< NULL >>
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

	valueTree := tree.Tree{}
	vt := tree.New("Value tree")
	vt.Add("value tree node")
	valueTree.Add(vt)
	tr.Add(valueTree)

	fmt.Fprintf(os.Stdout, "%s\n", tr)

	// Output:
	// Main tree
	// ├──Node 1
	// │  of main tree
	// ├──Node 2 of main tree
	// ├──Some string from buffer
	// │  with multilines
	// ├──Some error
	// ├──{42 23 3.1415927}
	// ├──Sub tree
	// │  └──Node 1 of sub tree
	// └──<< NULL >>
	//    └──Value tree
	//       └──value tree node
}
