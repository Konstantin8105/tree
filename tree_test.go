package tree_test

import (
	"bytes"
	"fmt"

	"github.com/Konstantin8105/tree"
)

func ExampleTree() {
	artist := tree.New("Pantera")
	album := artist.AddLine("Far Beyond Driven")
	album.AddLine("5 minutes Alone")
	album.AddLine("Some another")
	artist.AddLine("Power Metal")
	fmt.Println(artist)

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
	artist.Name = tree.Line(name)
	album := artist.AddLine("Поддерево\nс многострочным\nтекстом")
	album.AddLine("Лист поддерева\nзеленый")
	album.AddLine("Лист красный")
	artist.AddLine("Лист\nжелтый")
	fmt.Println(artist)

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
	tr.AddLine("Node 1\nof main tree\n\n\n")
	tr.AddLine("          Node 2 of main tree          ")

	subTr := tree.New("Sub tree")
	subTr.AddLine("Node 1 of sub tree")
	node := subTr.AddLine("Node 2 of sub tree")

	subsubTr := tree.New("Sub tree")
	subsubTr.AddLine("Node 1 of sub tree")
	subsubTr.AddLine("Node 2\nof sub tree")

	in := node.AddLine("\n\n\nIntermediant node")
	in.AddLine("some node")
	in.AddTree(subsubTr)

	node.AddLine("")  // empty name
	node.AddLine("B") // small name
	node.AddTree(subsubTr)

	tr.AddTree(subTr)

	ln := tr.AddLine("Last main node")

	var b bytes.Buffer
	b.WriteString("Some string from buffer\nwith multilines")
	ln.Add(&b)

	fmt.Println(tr)

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
	fmt.Println(tr)

	// Output:
}
