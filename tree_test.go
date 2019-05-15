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
	fmt.Fprintf(os.Stdout, "%s\n", artist.Print())
	tree.Walk(artist, checkWalk)

	// Output:
	// Pantera
	// ├──Far Beyond Driven
	// │  ├──5 minutes Alone
	// │  └──Some another
	// └──Power Metal
	//
	// Node: string                    |Pantera
	// Node: string                    |Far Beyond Driven
	// Node: string                    |5 minutes Alone
	// Node: string                    |Some another
	// Node: string                    |Power Metal
}

func ExampleTreeMultiline() {
	const name string = "Дерево"
	artist := tree.New(name)
	album := artist.Add("Поддерево\nс многострочным\nтекстом")
	album.Add("Лист поддерева\nзеленый")
	album.Add("Лист красный")
	artist.Add("Лист\nжелтый")
	fmt.Fprintf(os.Stdout, "%s\n", artist.Print())
	tree.Walk(artist, checkWalk)

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
	//
	// Node: string                    |Дерево
	// Node: string                    |Поддерево << BreakLine >> с многострочным << BreakLine >> текстом
	// Node: string                    |Лист поддерева << BreakLine >> зеленый
	// Node: string                    |Лист красный
	// Node: string                    |Лист << BreakLine >> желтый
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

	fmt.Fprintf(os.Stdout, "%s\n", tr.Print())
	tree.Walk(tr, checkWalk)

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
	//
	// Node: string                    |Main tree
	// Node: string                    |Node 1 << BreakLine >> of main tree
	// Node: string                    |Node 2 of main tree
	// Node: string                    |Sub tree
	// Node: string                    |Node 1 of sub tree
	// Node: string                    |Node 2 of sub tree
	// Node: string                    |Intermediant node
	// Node: string                    |some node
	// Node: string                    |Sub tree
	// Node: string                    |Node 1 of sub tree
	// Node: string                    |Node 2 << BreakLine >> of sub tree
	// Node: string                    |
	// Node: string                    |B
	// Node: string                    |Sub tree
	// Node: string                    |Node 1 of sub tree
	// Node: string                    |Node 2 << BreakLine >> of sub tree
	// Node: string                    |Last main node
	// Node: *bytes.Buffer             |Some string from buffer << BreakLine >> with multilines
}

func ExampleEmptyTree() {
	tr := tree.Tree{}
	fmt.Fprintf(os.Stdout, "%s\n", tr.Print())
	tree.Walk(&tr, checkWalk)

	tr2 := (*tree.Tree)(nil)
	tree.Walk(tr2, checkWalk)

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
	str.Add(nil)
	str.Add(nil)

	tr.Add(&str)

	tr.Add(nil)
	tr.Add(nil)
	tr.Add((*tree.Tree)(nil))
	tr.Add((*TempStruct)(nil))

	fmt.Fprintf(os.Stdout, "%s\n", tr.Print())
	tree.Walk(&tr, checkWalk)

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
	//
	// Node: <nil>                     |<nil>
	// Node: string                    |
	// Node: <nil>                     |<nil>
	// Node: <nil>                     |<nil>
	// Node: <nil>                     |<nil>
	// Node: <nil>                     |<nil>
	// Node: *tree.Tree                |<nil>
	// Node: *tree_test.TempStruct     |<nil>
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
	vt := valueTree.Add("Value tree")
	vt.Add("value tree node")

	tr.Add(valueTree)

	fmt.Fprintf(os.Stdout, "%s\n", tr.Print())
	tree.Walk(tr, checkWalk)

	// Output:
	// Main tree
	// ├──Node 1
	// │  of main tree
	// ├──Node 2 of main tree
	// ├──Some string from buffer
	// │  with multilines
	// ├──Some error
	// ├──{42 23 3.1415927}
	// └──Sub tree
	//    └──Node 1 of sub tree
	//
	// Node: string                    |Main tree
	// Node: string                    |Node 1 << BreakLine >> of main tree
	// Node: string                    |Node 2 of main tree
	// Node: *bytes.Buffer             |Some string from buffer << BreakLine >> with multilines
	// Node: *errors.errorString       |Some error
	// Node: tree_test.TempStruct      |{42 23 3.1415927}
	// Node: string                    |Sub tree
	// Node: string                    |Node 1 of sub tree
}

func checkWalk(str interface{}) {
	name := fmt.Sprintf("%v", str)
	name = strings.TrimSpace(name)
	name = strings.ReplaceAll(name, "\n", " << BreakLine >> ")
	fmt.Fprintf(os.Stdout, "Node: %-25s |%s\n", fmt.Sprintf("%T", str), name)
}
