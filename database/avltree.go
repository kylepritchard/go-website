package kyledb

import (
	"fmt"
	"strconv"
	"time"
)

type Post struct {
	title				string
	content			string
	id							string
	postdate		time.Time
}

func CreatePost(title string, content string, store map[string]Post) bool{
	//initialise a Post struct
	var post Post
	// make an id
	id := "1234"
	// make postTime
	time := time.Now()
	
	//build the Post struct
	post.title = title
	post.content = content
	post.postdate = time
	post.id = id
	
	store[id] = post
	
	// start indexing?????
	return true
}

//func createIndexs(post Post, indexMap *map[string]Tree) bool {
//	fmt.Println("map ====> ", indexMap)
//	indexMap["test"].Insert(101)
//	return true
//}

// Tree Structure
type Tree struct {
	root *Node
}

// Exists to see if this is in the tree or not.
func (tree Tree) Exists(value int) bool {
	var result *Node
	if tree.root == nil {
		return false
	}
	
	result = tree.root.Find(value) 
	if result == nil {
		return false
	}
	fmt.Println(result.value)
	return true
}

// Insert an element in the tree
func (tree *Tree) Insert(value int) {
	tree.root, _ = tree.root.Insert(value)
}

// Delete return true if successful
func (tree *Tree) Delete(value int) bool {
	var change int

	if tree.root == nil {
		return false
	}

	tree.root, change = tree.root.Delete(value)

	if change == 0 {
		return false
	}

	return true
}

// Update the key with a new one
func (tree *Tree) Update(orig int, value int) {
	tree.Delete(orig)
	tree.Insert(value)
}

// Print the tree
func (tree Tree) Print() {
	if tree.root == nil {
		fmt.Printf("--- AVL Tree:\n    EMPTY\n")
		return
	}

	fmt.Printf("--- AVL Tree:\n")
	tree.root.Print(1)
}

/*
   AVL Tree Nodes
*/

// Node in the tree
type Node struct {
	value int
	id				string
	balance int
	left    *Node
	right   *Node
}

// Compare two values,
func (node *Node) Compare(value int) int {
	return node.value - value
}

// Find a value in the tree, searching
func (node *Node) Find(value int) *Node {
	if node == nil {
		return nil
	}

	if node.Compare(value) == 0 {
		return node
	}

	var result *Node

	if node.left != nil {
		result = node.left.Find(value)
	}

	if result == nil && node.right != nil {
		result = node.right.Find(value)
	}

	return result
}

// RotateLeft the tree 
func (node *Node) RotateLeft() *Node {
	if node == nil {
		return nil
	}

	if node.right == nil {
		return node
	}

	var result = node.right

	node.right = result.left
	result.left = node

	var leftBalance = node.balance
	var balance = result.balance

	node.balance = leftBalance - 1 - max(balance, 0)
	result.balance = min(leftBalance-2, balance+leftBalance-2, balance-1)

	return result
}

// RotateRight the tree 
func (node *Node) RotateRight() *Node {
	if node == nil {
		return nil
	}

	if node.left == nil {
		return node
	}

	var result = node.left
	node.left = result.right
	result.right = node

	var rightBalance = node.balance
	var balance = result.balance

	node.balance = rightBalance + 1 - min(balance, 0)
	result.balance = max(rightBalance+2, balance+rightBalance+2, balance+1)

	return result
}

// Insert a node
func (node *Node) Insert(value int) (*Node, int) {

	// Terminal Condition, create this node
	if node == nil {
		return &Node{value: value, balance: 0}, 1
	}

	var change int

	// Descend to the children
	diff := node.Compare(value)
	switch {

	case diff == 0:
		// Ignore duplicates
		fmt.Println("cannot have 2 values the same")

	case diff > 0:
		node.left, change = node.left.Insert(value)
		change *= -1

	case diff < 0:
		node.right, change = node.right.Insert(value)
	}

	node.balance += change

	// Rebalance at the parents or grandparents
	var insert int

	if node.balance != 0 && change != 0 {
		switch {

		case node.balance < -1:
			if node.left.balance >= 0 {
				node.left = node.left.RotateLeft()
			}
			node = node.RotateRight()
			insert = 0

		case node.balance > 1:
			if node.right.balance <= 0 {
				node.right = node.right.RotateRight()
			}
			node = node.RotateLeft()
			insert = 0

		default:
			insert = 1
		}
	} else if change != 0 {
		insert = 0
	}

	return node, insert
}

// Delete a node
func (node *Node) Delete(value int) (*Node, int) {
	var change int

	if node == nil {
		return nil, change
	}

	diff := node.Compare(value)
	switch {
	case diff > 0:
		node.left, change = node.left.Delete(value)

	case diff < 0:
		node.right, change = node.right.Delete(value)
		change *= -1

	case diff == 0:
		switch {
		case node.left == nil:
			return node.right, 1

		case node.right == nil:
			return node.left, 1

		default:
			// Pick the heavier of the two...
			if -1*node.left.balance < node.right.balance {
				node = node.RotateLeft()
				node.left, change = node.left.Delete(value)

			} else {
				node = node.RotateRight()
				node.right, change = node.right.Delete(value)
				change *= -1
			}
		}
	}

	// Update the balance
	if change != 0 {

		if node.balance != change {
			node.balance += change
		}

		switch {
		case node.balance < -1:
			if node.left.balance >= 0 {
				node.left = node.left.RotateLeft()
			}
			node = node.RotateRight()

		case node.balance > 1:
			if node.right.balance <= 0 {
				node.right = node.right.RotateRight()
			}
			node = node.RotateLeft()
		}
	}

	return node, change
}

// Print the current nodes, rotated 90 degrees, in-order traversal.
func (node Node) Print(depth int) {

	if node.right != nil {
		node.right.Print(depth + 1)
	}

	padding := padding(depth)
	fmt.Printf("%s[%v] Value: %v\n", padding, symbols(node.balance), node.value)

	if node.left != nil {
		node.left.Print(depth + 1)
	}
}

// Slightly more decorative: shift 1/-1 -> +/-
func symbols(balance int) string {
	switch {
	case balance == 1:
		return "+"
	case balance == -1:
		return "-"
	}
	return strconv.Itoa(balance)
}

// Convert depth into spaces, 4 per
func padding(size int) string {
	result := ""
	for i := 0; i < size; i++ {
		result += "    "
	}
	return result
}

func max(values ...int) int {
	var total = values[0]
	for _, value := range values {
		if value > total {
			total = value
		}
	}
	return total
}

func min(values ...int) int {
	var total = values[0]
	for _, value := range values {
		if value < total {
			total = value
		}
	}
	return total
}

func (t *Tree) Traverse(n *Node, f func(*Node)) {
	if n == nil {
		return
	}
	t.Traverse(n.left, f)
	f(n)
	t.Traverse(n.right, f)
}

func main(){
	
	//Create a map to store all the posts with ID as key
	postStore := make(map[string]Post)
	
	// Create a map to store pointers to the index trees created
	indexStore := make(map[string]*Tree)
	
	
	ok := CreatePost("test title", "content ...", postStore)
	if ok {
		fmt.Println(postStore)
		fmt.Println(postStore["1234"])
		fmt.Println(postStore["1234"].postdate.String())
		fmt.Println(postStore["1234"].title)
		fmt.Println(postStore["1234"].content)
		fmt.Println(postStore["1234"].id)
	} else {
		fmt.Println("didnt post :( ")
	}
	
	//tree := &Tree{}
	var tree Tree
	var tree2 Tree
	indexStore["test"] = &tree
	indexStore["test2"] = &tree2
	for i := 0; i < 100; i++ {
		tree.Insert(i)
		tree2.Insert(i)
	}
	
	//tree.Insert(2)
	//tree.Print()
	
	
	tree.Exists(55)
	
	//tree.Print()
	
	//tree.Traverse(tree.root, func(n *Node) { fmt.Print(n.value, " | ") })
	
	for i := 50; i < 75; i++ {
		tree.Delete(i)
	}
	
	//tree.Traverse(tree.root, func(n *Node) { fmt.Print(n.value, " | ") })
	
	//tree.Exists(55)
	
	//tree.Print()
	//indexStore["test"] = tree
	fmt.Println(indexStore["test"].Exists(60))
	fmt.Println(indexStore["test2"].Exists(10))
}