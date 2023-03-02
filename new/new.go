package main

import "fmt"

type Node struct { // folder
	id       string
	parentID string
}

type Tree struct {
	node  *Node
	child []*Tree
}

func recursive(folders []Node, tree *Tree) {
	for _, folder := range folders {

		if folder.parentID == tree.node.id {
			tree.child = append(tree.child, &Tree{
				node: &Node{
					id:       folder.id,
					parentID: folder.parentID,
				},
			})
			recursive(folders, tree.child[len(tree.child)-1])
		}
	}
}

func main() {
	tree := Tree{}
	folders := []Node{
		{
			id:       "0",
			parentID: "",
		},
		{
			id:       "1",
			parentID: "0",
		},
		{
			id:       "2",
			parentID: "0",
		},
		{
			id:       "3",
			parentID: "0",
		},
		{
			id:       "4",
			parentID: "2",
		},
		{
			id:       "6",
			parentID: "5",
		},
		{
			id:       "5",
			parentID: "3",
		},
	}

	for _, folder := range folders {
		if folder.parentID == "" {
			tree = Tree{
				node: &Node{
					id:       folder.id,
					parentID: folder.parentID,
				},
			}
		}
	}

	recursive(folders, &tree)
	printTr(&tree, 0)
}

func printTr(tree *Tree, index int32) {
	fmt.Println("Index: ", index, " Value: ", tree.node)
	for _, elem := range tree.child {
		printTr(elem, index+1)
	}
}
