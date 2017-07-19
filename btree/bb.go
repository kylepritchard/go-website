package main

import (
	"fmt"
	"time"
)

// Store is the main slice containing all the Posts
var Store []Post

type Post struct {
	Id         int //'Unique' Key?
	Title      string
	Content    string
	Slug       string //slugified version of the title - for routing
	PostDate   time.Time
	FeatureImg string
}

type Node struct {
	Level int
	Left  *Node
	Right *Node
	Key   int
	Value string
}

type Tree struct {
	Root *Node
}

var NilNode Node

func makeNilNode() *Node {
	// Initialise Nil Node
	NilNode.Level = 0
	NilNode.Left = nil
	NilNode.Right = nil
	NilNode.Key = 0
	NilNode.Value = ""
	return &NilNode
}

func NewTree() *Tree {
	t := new(Tree)
	t.Root = &NilNode
	return t
}

// NewNode creates a new node in the tree containing the key/value. Left and Right pointers are nil
func NewNode(key int, value string) *Node {
	return &Node{1, nil, nil, key, value}
}

//Skew function
func Skew(n *Node) *Node {
	if n.Left != nil {
		if n.Level != 0 && n.Left.Level == n.Level {
			// fmt.Println("Skew")
			// var save = n.Left
			// n.Left = save.Right
			// save.Right = n
			// n = save

			//JS Skew
			var temp = n
			n = n.Left
			temp.Left = n.Right
			n.Right = temp
		}
	}
	return n
}

// function split(node) {
//     if (node.Right.Right.Level === node.Level) {
//         var temp = node;
//         node = node.Right;
//         temp.Right = node.Left;
//         node.Left = temp;
//         node.Level++;
//     }
//     return node;
// }

// Split function
func Split(n *Node) *Node {
	// time.Sleep(time.Second * 2)
	if n.Right != nil && n.Right.Right != nil {
		if n.Level != 0 && n.Right.Right.Level == n.Level {

			// var save = n.Right
			// n.Right = save.Left
			// save.Left = n
			// n = save
			// n.Level++

			//JS Split
			var temp = n
			n = n.Right
			temp.Right = n.Left
			n.Left = temp
			n.Level++

		}
	}
	return n
}

// Insert new item to the tree
func (tree *Tree) Insert(key int, value string) {
	tree.Root = tree.Root.insert(key, value)
}

// Recursive insert
// func (n *Node) insert(key int, value string) *Node {
// 	if n == nil || n.Level == 0 {
// 		return NewNode(key, value)
// 	}
// 	if n.Value < value {
// 		n.Right = n.Right.insert(key, value)

// 	} else {
// 		n.Left = n.Left.insert(key, value)
// 	}
// 	n = Skew(n)
// 	n = Split(n)
// 	return n
// }

// Iterative Insert
func (n *Node) insert(key int, value string) *Node {
	if n == nil || n.Level == 0 {
		n = NewNode(key, value)
	} else {
		var save = n
		var up []*Node
		top := 0
		dir := 0
		for {
			up = append(up, save)
			if save.Value < value {
				dir = 1
				if save.Right == nil {
					break
				}
				save = save.Right
			} else {
				dir = 0
				if save.Left == nil {
					break
				}
				save = save.Left
			}
			top++
		}

		if dir == 0 {
			save.Left = NewNode(key, value)
		} else {
			save.Right = NewNode(key, value)
		}

		for i := top - 1; i >= 0; i-- {
			if i != 0 {
				if up[i-1].Right == up[i] {
					dir = 1
				} else {
					dir = 0
				}
			}

			up[i] = Split(Skew(up[i]))

			if i != 0 {
				if dir == 0 {
					up[i-1].Left = up[i]
				} else {
					up[i-1].Right = up[i]
				}
			} else {
				n = up[i]
			}
		}
	}
	return n
}

// Tree Remove

func (tree *Tree) Remove(value string) {
	tree.Root = remove(tree.Root, value)
}

// Recursive Removal
// func remove(n *Node, value string) *Node {
// 	fmt.Println("remove", value)
// 	var heir *Node
// 	if n != nil {
// 		if n.Value == value {
// 			if n.Left != nil && n.Right != nil {
// 				heir = n.Left
// 				for heir.Right != nil {
// 					heir = heir.Right
// 				}
// 				n.Key = heir.Key
// 				n.Value = heir.Value
// 				n.Left = remove(n.Left, n.Value)
// 			} else if n.Left == nil {
// 				n = n.Right
// 			} else {
// 				n = n.Left
// 			}
// 		} else if n.Value < value {
// 			n.Right = remove(n.Right, value)
// 		} else {
// 			n.Left = remove(n.Left, value)
// 		}
// 	}

// 	if n.Left.Level < (n.Level-1) || n.Right.Level < (n.Level-1) {
// 		n.Level--
// 		if n.Right.Level > n.Level {
// 			n.Right.Level = n.Level
// 		}
// 		n = Split(Skew(n))
// 	}

// 	return n
// }

func remove(root *Node, value string) *Node {
	if root != nil {
		var it = root
		var up []*Node
		top := 0
		dir := 0
		dir2 := 0

		for {
			up = append(up, it)

			if it == nil {
				return root
			} else if value == it.Value {
				break
			}

			if it.Value < value {
				dir = 1
				it = it.Right
			} else {
				dir = 0
				it = it.Left
			}

			top++
		}

		if it.Left == nil || it.Right == nil {
			if it.Left == nil {
				dir2 = 1
			}

			if top > 1 {
				if dir == 0 {
					if dir2 == 0 {
						up[top-1].Left = it.Left
					} else {
						up[top-1].Left = it.Right
					}
				} else {
					if dir2 == 0 {
						up[top-1].Right = it.Left
					} else {
						up[top-1].Right = it.Right
					}
				}
			} else {
				root = it.Right
			}
		} else {

			var heir = it.Right
			var prev = it

			for {
				if heir.Left == nil {
					break
				}

				up = append(up, prev)
				heir = prev
				heir = heir.Left
				top++
			}

			it.Value = heir.Value
			if prev == it {
				prev.Right = heir.Right
			} else {
				prev.Left = heir.Right
			}
		}

		for i := top - 1; i >= 0; i-- {
			if i != 0 {
				if up[i-1].Right == up[i] {
					dir = 1
				} else {
					dir = 0
				}
			}

			if up[i].Left.Level < up[i].Level-1 || up[i].Right.Level < up[i].Level-1 {
				if up[i].Right.Level > up[i].Level-1 {
					up[i].Right.Level = up[i].Level
				}

				up[i] = Skew(up[i])
				up[i].Right = Skew(up[i].Right)
				up[i].Right.Right = Skew(up[i].Right.Right)
				up[i] = Split(up[i])
				up[i].Right = Split(up[i].Right)
			}

			if i != 0 {
				if dir == 0 {
					up[i-1].Left = up[i]
				} else {
					up[i-1].Right = up[i]
				}
			} else {
				root = up[i]
			}
		}
	}

	return root
}

// Tree Find

func (tree *Tree) Find(key int) string {

	n := tree.Root

	for n.Level != 0 {
		if n.Key == key {
			return n.Value
		}
		if n.Key < key {
			n = n.Right
		} else {
			n = n.Left
		}
	}

	return ""
}

// Traversing

// TREE TRAVERSAL
func (tree *Tree) InOrderTraversal() []string {
	t := inorderTraversal(tree.Root)
	return t
}

func inorderTraversal(n *Node) []string {
	var res []string

	cur := n
	for cur != nil {
		if cur.Left == nil {
			res = append(res, cur.Value)
			// if len(res) == l {
			// 	return res
			// }
			cur = cur.Right
			continue
		}

		pre := cur.Left
		for pre.Right != nil && pre.Right != cur {
			pre = pre.Right
		}

		if pre.Right == nil {
			pre.Right = cur
			cur = cur.Left
		} else {
			pre.Right = nil
			res = append(res, cur.Value)
			// if len(res) == l {
			// 	return res
			// }
			cur = cur.Right
		}
	}
	return res
}

var marker = "§"

type serialData struct {
	Value string
	index int
}

// SERIALIZING THE TREE
func (tree *Tree) Serialize() {
	// If current node is NULL, store marker
	var n = tree.Root
	serialize(n)
}

func serialize(n *Node) {
	if n == nil {
		fmt.Print(marker)
		return
	}

	// Else, store current node and recur for its children

	fmt.Print(serialData{n.Value, n.Key})
	serialize(n.Left)
	serialize(n.Right)
}

// func DeSerialize(t *Tree, s string)
// {
//     // Read next item from file. If theere are no more items or next
//     // item is marker, then return
//     int val;
//     if !fscanf(fp, "%d ", &val) || val == MARKER {
// 	   return
// 	}

//     // Else create node with this item and recur for children
// 	root = newNode(val);

//     DeSerialize(root->left, fp);
//     DeSerialize(root->right, fp);
// }

func main() {

	// var tree Tree

	// // tree := new(Tree)
	// tree.Root = &NilNode

	tree := NewTree()

	letters := []string{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p"}
	for i := 0; i < 10; i++ {
		tree.Insert(i, letters[i])
	}

	fmt.Println(tree.InOrderTraversal())

	// 	markdown := []byte(`
	// # New Features!

	// 	- Import a HTML file and watch it magically convert to Markdown
	// 	- Drag and drop images (requires your Dropbox account be linked)
	// `)

	// 	formatted := blackfriday.MarkdownCommon(markdown)

	// 	fmt.Println(string(formatted))

	// Serialize(tree.Root)

	// var m = map[int]string{}
	// var size []string
	// size1000 := make([]string, 1000)
	// // size100000 := make([]string, 100000)

	// for i := 0; i < 1000; i++ {
	// 	num := randomdata.Address()
	// 	size = append(size, num)
	// 	size1000 = append(size1000, num)
	// 	// size100000 = append(size100000, num)
	// 	m[i] = num
	// }

}
