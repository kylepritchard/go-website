package main

import (
	"fmt"
	"testing"
	"time"

	randomdata "github.com/Pallinder/go-randomdata"
)

type Node struct {
	level int
	left  *Node
	right *Node
	key   int
	value string
}

var NilNode Node

type Tree struct {
	root *Node
}

//Need to make a node instead of changing a referenced node
func NewNode(key int, value string) *Node {
	// var node *Node
	// node :=
	return &Node{1, &NilNode, &NilNode, key, value}
}

//Skew function
func Skew(n *Node) *Node {
	if n.level != 0 && n.left.level == n.level {
		// fmt.Println("Skew")
		// var save = n.left
		// n.left = save.right
		// save.right = n
		// n = save

		//JS Skew
		var temp = n
		n = n.left
		temp.left = n.right
		n.right = temp
	}
	return n
}

// function split(node) {
//     if (node.right.right.level === node.level) {
//         var temp = node;
//         node = node.right;
//         temp.right = node.left;
//         node.left = temp;
//         node.level++;
//     }
//     return node;
// }

func Split(n *Node) *Node {
	// time.Sleep(time.Second * 2)
	if n.level != 0 && n.right.right.level == n.level {

		// var save = n.right
		// n.right = save.left
		// save.left = n
		// n = save
		// n.level++

		//JS Split
		var temp = n
		n = n.right
		temp.right = n.left
		n.left = temp
		n.level++

	}
	return n
}

func (n *Node) insert(key int, value string) *Node {
	if n == nil || n.level == 0 {
		return NewNode(key, value)
	}
	if n.value < value {
		n.right = n.right.insert(key, value)

	} else {
		n.left = n.left.insert(key, value)
	}
	n = Skew(n)
	n = Split(n)
	return n
}

func (t *Tree) Insert(key int, value string) {
	t.root = t.root.ItInsert(key, value)
}

func (t *Tree) Remove(value string) {
	t.root = ItRemove(t.root, value)
}

func remove(n *Node, value string) *Node {
	fmt.Println("remove", value)
	var heir *Node
	if n != &NilNode {
		if n.value == value {
			if n.left != &NilNode && n.right != &NilNode {
				heir = n.left
				for heir.right != &NilNode {
					heir = heir.right
				}
				n.key = heir.key
				n.value = heir.value
				n.left = remove(n.left, n.value)
			} else if n.left == &NilNode {
				n = n.right
			} else {
				n = n.left
			}
		} else if n.value < value {
			n.right = remove(n.right, value)
		} else {
			n.left = remove(n.left, value)
		}
	}

	if n.left.level < (n.level-1) || n.right.level < (n.level-1) {
		n.level--
		if n.right.level > n.level {
			n.right.level = n.level
		}
		n = Split(Skew(n))
	}

	return n
}

func ItRemove(root *Node, value string) *Node {
	fmt.Println("remove", value)
	if root != &NilNode {
		var it = root
		var up []*Node
		top := 0
		dir := 0
		dir2 := 0

		for {
			up = append(up, it)
			fmt.Println(up)

			// if (it == nil)
			// {
			//     return root;
			// }
			// else if (data == it->data)
			// {
			//     break;
			// }

			// dir = it->data < data;
			// it = it->link[dir];

			if it == &NilNode {
				return root
			} else if value == it.value {
				fmt.Println("found")
				break
			}

			if it.value < value {
				dir = 1
				it = it.right
			} else {
				dir = 0
				it = it.left
			}

			top++
		}

		fmt.Println(up)

		if it.left == &NilNode || it.right == &NilNode {
			if it.left == &NilNode {
				dir2 = 1
			}

			if top > 1 {
				if dir == 0 {
					if dir2 == 0 {
						up[top-1].left = it.left
					} else {
						up[top-1].left = it.right
					}
				} else {
					if dir2 == 0 {
						up[top-1].right = it.left
					} else {
						up[top-1].right = it.right
					}
				}
			} else {
				root = it.right
			}
		} else {

			var heir = it.right
			var prev = it

			for {
				if heir.left == &NilNode {
					break
				}

				up = append(up, prev)
				heir = prev
				heir = heir.left
				top++
			}

			it.value = heir.value
			if prev == it {
				prev.right = heir.right
			} else {
				prev.left = heir.right
			}
		}

		for i := top - 1; i >= 0; i-- {
			if i != 0 {
				if up[i-1].right == up[i] {
					dir = 1
				} else {
					dir = 0
				}
			}

			if up[i].left.level < up[i].level-1 || up[i].right.level < up[i].level-1 {
				if up[i].right.level > up[i].level-1 {
					up[i].right.level = up[i].level
				}

				up[i] = Skew(up[i])
				up[i].right = Skew(up[i].right)
				up[i].right.right = Skew(up[i].right.right)
				up[i] = Split(up[i])
				up[i].right = Split(up[i].right)
			}

			if i != 0 {
				if dir == 0 {
					up[i-1].left = up[i]
				} else {
					up[i-1].right = up[i]
				}
			} else {
				root = up[i]
			}
		}
	}

	return root
}

func init() {
	// Initialise Nil Node
	NilNode.level = 0
	NilNode.left = &NilNode
	NilNode.right = &NilNode
	NilNode.key = 0
	NilNode.value = ""
}

func (t *Tree) Find(key int) string {

	n := t.root

	for n.level != 0 {
		if n.key == key {
			return n.value
		}
		if n.key < key {
			n = n.right
		} else {
			n = n.left
		}
	}

	return ""
}

func (n *Node) ItInsert(key int, value string) *Node {
	// fmt.Println("insert")
	if n == nil || n.level == 0 {
		// fmt.Println("new tree")
		n = NewNode(key, value)
	} else {
		var save = n
		// var up = make([]*Node, 1)
		var up []*Node
		top := 0
		dir := 0
		for {
			up = append(up, save)
			// top++
			if save.value < value {
				// fmt.Println("right")
				// fmt.Println(save.right.level)
				dir = 1
				if save.right == &NilNode {
					// fmt.Println("break loop")
					break
				}
				// fmt.Println(save, save.right)
				save = save.right
				// fmt.Println(save)
			} else {
				// fmt.Println("left")
				dir = 0
				if save.left == &NilNode {
					break
				}
				save = save.left
			}
			top++
		}

		if dir == 0 {
			// fmt.Println("New Node left")
			save.left = NewNode(key, value)
		} else {
			// fmt.Println("New Node right")
			save.right = NewNode(key, value)
		}

		for i := top - 1; i >= 0; i-- {
			if i != 0 {
				if up[i-1].right == up[i] {
					dir = 1
				} else {
					dir = 0
				}
			}

			// up[i] = Skew(up[i])
			up[i] = Split(Skew(up[i]))

			if i != 0 {
				if dir == 0 {
					up[i-1].left = up[i]
				} else {
					up[i-1].right = up[i]
				}
			} else {
				n = up[i]
			}
		}
	}
	return n
}

// Traversing

func (t *Tree) MaxMin(max bool) (int, string) {

	n := t.root
	top := 0
	var path []*Node

	/* Build a path to work with */
	if n != nil {
		if max {
			for n.right.level != 0 {
				path = append(path, n)
				n = n.right
				top++
			}
		}
		if !max {
			for n.left.level != 0 {
				path = append(path, n)
				n = n.left
				top++
			}
		}
	}

	return n.key, n.value

}

var sl []int

func Ascend(n *Node) {
	if n.level == 0 {
		return
	}
	Ascend(n.left)
	// fmt.Println(n.value)
	sl = append(sl, n.key)
	Ascend(n.right)
}

func Descend(n *Node) {
	if n.level == 0 {
		return
	}
	Descend(n.right)
	fmt.Println(n.key)
	Descend(n.left)
}

func InorderTraversal(n *Node) []string {
	var res []string

	cur := n
	for cur != &NilNode {
		if cur.left == &NilNode {
			res = append(res, cur.value)
			// if len(res) == l {
			// 	return res
			// }
			cur = cur.right
			continue
		}

		pre := cur.left
		for pre.right != &NilNode && pre.right != cur {
			pre = pre.right
		}

		if pre.right == &NilNode {
			pre.right = cur
			cur = cur.left
		} else {
			pre.right = &NilNode
			res = append(res, cur.value)
			// if len(res) == l {
			// 	return res
			// }
			cur = cur.right
		}
	}
	return res
}

func main() {

	tree := &Tree{}
	// tree.root = &NilNode
	// fmt.Println(tree.root)
	// tree.Insert(0, "0")
	// tree.Insert(1, "1")
	// tree.Insert(2, "2")
	// tree.Insert(3, "3")
	// tree.Insert(4, "4")
	// tree.Insert(5, "5")

	letters := []string{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p"}
	timeStart := time.Now()
	for i := 0; i < 10; i++ {
		tree.Insert(i, letters[i])
	}
	fmt.Println("Tree Insertion done in:", time.Since(timeStart))

	// timeStart = time.Now()
	// str := tree.Find(4563252)
	// fmt.Println("Tree Find:", str, "in: ", time.Since(timeStart))

	// fmt.Println(tree.MaxMin(false))
	// fmt.Println(tree.MaxMin(true))
	// fmt.Println(Ascend(tree.root))
	// Ascend(tree.root)
	// Descend(tree.root)

	fmt.Println("1st", InorderTraversal(tree.root))

	tree.Remove("q")

	fmt.Println("After Removal", InorderTraversal(tree.root))

	tree.Remove("w")

	fmt.Println("After Removal", InorderTraversal(tree.root))

	tree.Remove("e")

	fmt.Println("After Removal", InorderTraversal(tree.root))

	tree.Remove("r")

	fmt.Println("After Removal", InorderTraversal(tree.root))

	var slice [200000]string
	var m = make(map[int]string)
	var a string
	// letters := []string{"a", "g", "e", "j", "h", "b", "d", "i", "f", "c"}
	for i := 0; i < 200000; i++ {
		city := randomdata.City()
		slice[i] = city
		m[i] = city
	}
	timeStart = time.Now()
	// for i := 0; i < 1; i++ {
	// 	a = slice[i]
	// }
	a = slice[12345]
	t := time.Since(timeStart)
	fmt.Println(a, t)
	timeStart = time.Now()
	// _, ok := m[1]

	a = m[12345]
	t = time.Since(timeStart)
	fmt.Println(a, t)

}

// func BenchmarkInsert(b *testing.B) {
// 	tree := &Tree{}
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		tree.Insert(i, string(i))
// 	}
// }

// func BenchmarkGet(b *testing.B) {
// 	tree := &Tree{}
// 	for i := 0; i < 1000000; i++ {
// 		tree.Insert(i, string(i))
// 	}
// 	b.ResetTimer()
// 	var sl []string
// 	for i := 0; i < b.N; i++ {
// 		s := tree.Find(i)
// 		sl = append(sl, s)
// 	}
// 	fmt.Println(len(sl))
// }

// func BenchmarkTrav5(b *testing.B) {
// 	tree := &Tree{}
// 	for i := 0; i < 100; i++ {
// 		tree.Insert(i, randomdata.StringNumber(2, "-"))
// 	}
// 	var sl []string
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		sl = InorderTraversal(tree.root)
// 	}
// 	fmt.Println(sl[:5])
// }

// func BenchmarkTrav10(b *testing.B) {
// 	tree := &Tree{}
// 	// letters := []string{"a", "g", "e", "j", "h", "b", "d", "i", "f", "c"}
// 	for i := 0; i < 100; i++ {
// 		tree.Insert(i, randomdata.StringNumber(2, "-"))
// 	}
// 	// var sl []string
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		InorderTraversal(tree.root, 10)
// 	}
// 	// fmt.Println(sl)
// }

// func BenchmarkTrav100(b *testing.B) {
// 	tree := &Tree{}
// 	// letters := []string{"a", "g", "e", "j", "h", "b", "d", "i", "f", "c"}
// 	for i := 0; i < 100; i++ {
// 		tree.Insert(i, randomdata.StringNumber(2, "-"))
// 	}
// 	// var sl []string
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		InorderTraversal(tree.root, 100)
// 	}
// 	// fmt.Println(sl)
// }

// func BenchmarkTrav1000(b *testing.B) {
// 	tree := &Tree{}
// 	// letters := []string{"a", "g", "e", "j", "h", "b", "d", "i", "f", "c"}
// 	for i := 0; i < 1000; i++ {
// 		tree.Insert(i, randomdata.StringNumber(2, "-"))
// 	}
// 	// var sl []string
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		InorderTraversal(tree.root, 1000)
// 	}
// 	// fmt.Println(sl)
// }

// func BenchmarkTrav1000BigData(b *testing.B) {
// 	tree := &Tree{}
// 	// letters := []string{"a", "g", "e", "j", "h", "b", "d", "i", "f", "c"}
// 	for i := 0; i < 1000; i++ {
// 		tree.Insert(i, randomdata.Address())
// 	}
// 	// var sl []string
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		InorderTraversal(tree.root, 1000)
// 	}
// 	// fmt.Println(sl)
// }

// func BenchmarkTravRecursive(b *testing.B) {
// 	tree := &Tree{}
// 	// letters := []string{"a", "g", "e", "j", "h", "b", "d", "i", "f", "c"}
// 	for i := 0; i < 100; i++ {
// 		tree.Insert(i, randomdata.StringNumber(2, "-"))
// 	}
// 	// var sl []string
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		Ascend(tree.root)
// 	}

// 	// fmt.Println(sl)
// }

func BenchmarkSliceLookup(b *testing.B) {
	var slice []int
	var a int
	// letters := []string{"a", "g", "e", "j", "h", "b", "d", "i", "f", "c"}
	for i := 0; i < 50000001; i++ {
		slice = append(slice, i)
	}
	fmt.Println(slice[49999999])
	// var sl []string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if a > 50000000 {
			i = 49999999
		}
		a = slice[i]
	}

	fmt.Println(a)
}

// func BenchmarkMapLookup(b *testing.B) {
// 	var m = make(map[int]int)
// 	var a int
// 	// letters := []string{"a", "g", "e", "j", "h", "b", "d", "i", "f", "c"}
// 	for i := 0; i < 50000000; i++ {
// 		m[i] = i
// 	}
// 	// var sl []string
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		a = m[i]
// 	}

// 	fmt.Println(a)
// }
