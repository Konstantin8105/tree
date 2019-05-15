package tree_test

import (
	"fmt"

	"github.com/Konstantin8105/tree"
)

func ExampleTree() {
	artist := tree.New("Pantera")
	album := artist.Add("Far Beyond Driven")
	album.Add("5 minutes Alone")
	album.Add("Some another")
	artist.Add("Power Metal")
	fmt.Println(artist)

	// Output:
	// Pantera
	// ├── Far Beyond Driven
	// │   ├── 5 minutes Alone
	// │   └── Some another
	// └── Power Metal
}

func ExampleTreeMultiline() {
	const name string = "Дерево"
	artist := tree.Tree{}
	artist.Name = name
	album := artist.Add("Поддерево\nс многострочным\nтекстом")
	album.Add("Лист поддерева\nзеленый")
	album.Add("Лист красный")
	artist.Add("Лист\nжелтый")
	fmt.Println(artist)

	// Output:
	// Дерево
	// ├── Поддерево
	// │   с многострочным
	// │   текстом
	// │   ├── Лист поддерева
	// │   │   зеленый
	// │   └── Лист красный
	// └── Лист
	//     желтый
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
	subsubTr.Add("Node 2 of sub tree")

	in := node.Add("Intermediant node")
	in.Add("some node")
	in.AddTree(subsubTr)

	node.Add("")  // empty name
	node.Add("B") // small name
	node.AddTree(subsubTr)

	tr.AddTree(subTr)
	fmt.Println(tr)

	// Output:
	// Main tree
	// ├── Node 1 of main tree
	// ├── Node 2 of main tree
	// └── Sub tree
	//     ├── Node 1 of sub tree
	//     └── Node 2 of sub tree
	//         ├── Intermediant node
	//         │   ├── some node
	//         │   └── Sub tree
	//         │       ├── Node 1 of sub tree
	//         │       └── Node 2 of sub tree
	//         ├──
	//         ├── B
	//         └── Sub tree
	//             ├── Node 1 of sub tree
	//             └── Node 2 of sub tree
}
