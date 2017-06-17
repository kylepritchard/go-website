package kyledb

import (
	"fmt"
	"strconv"
	"time"
	"os"
	"encoding/gob"
)


type Post struct {
	Title				string
	Content			string
	ID							string
	PostDate		time.Time
}

type PostStore struct {
	Store map[string]Post
}

type IndexStore struct {
	Idx	map[string]*Tree
}

type DB struct {
	PostStore PostStore
	Indexstore IndexStore
}

func InitDB(path string) bool{
	db := &DB{}
	
	file, err := os.Create(path)
	defer file.Close()
	if err == nil {
		enc := gob.NewEncoder(file)
		err := enc.Encode(db)
		if err != nil {
			fmt.Println(err)
			return false
		}
		return true
	}	
	return false
}

func Load(path string) (*DB, error) {
	db := &DB{}
	
	file, err := os.Open(path)
	defer file.Close()
	if err == nil {
		dec := gob.NewDecoder(file)
		err := dec.Decode(&db)
		if err != nil {
			return nil, err
		}
		return db, nil
	}
	return nil, err	
}

func (db *DB) Close() {
	fmt.Println("Does fuck all")
}


/*
file, err := os.Create("tree.idx")
	defer file.Close()
	if err == nil {
		enc := gob.NewEncoder(file)
		err := enc.Encode(tree)
 		if err != nil {
        fmt.Println(err)
   } 
	}
	
	tree := &Tree{}
	file, err := os.Open("tree.idx")
	defer file.Close()
	if err == nil {
		dec := gob.NewDecoder(file)
		err := dec.Decode(&tree)
		if err != nil {
			fmt.Println(err)
		}
	}
	*/

func (db *DB) MakeAndSave() bool {
	file, err := os.Open("database.kyled")
	defer file.Close()
	if err == nil {
		dec := gob.NewEncoder(file)
		err := dec.Encode(db)
		if err != nil {
			fmt.Println(err)
			return false
		}
		return true
	}
	return false
}












func CreatePost(title string, content string, store map[string]Post) bool{
	//initialise a Post struct
	var post Post
	// make an id
	id := "1234"
	// make postTime
	time := time.Now()
	
	//build the Post struct
	post.Title = title
	post.Content = content
	post.PostDate = time
	post.ID = id
	
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
	Root *Node
}

// Exists to see if this is in the tree or not.
func (tree Tree) Exists(value int) bool {
	var result *Node
	if tree.Root == nil {
		return false
	}
	
	result = tree.Root.Find(value) 
	if result == nil {
		return false
	}
	fmt.Println(result.Value)
	return true
}

// Insert an element in the tree
func (tree *Tree) Insert(value int) {
	tree.Root, _ = tree.Root.Insert(value)
}

// Delete return true if successful
func (tree *Tree) Delete(value int) bool {
	var change int

	if tree.Root == nil {
		return false
	}

	tree.Root, change = tree.Root.Delete(value)

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
	if tree.Root == nil {
		fmt.Printf("--- AVL Tree:\n    EMPTY\n")
		return
	}

	fmt.Printf("--- AVL Tree:\n")
	tree.Root.Print(1)
}

/*
   AVL Tree Nodes
*/

// Node in the tree
type Node struct {
	Value int
	ID			string
	Balance int
	Left    *Node
	Right   *Node
}

// Compare two values,
func (node *Node) Compare(value int) int {
	return node.Value - value
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

	if node.Left != nil {
		result = node.Left.Find(value)
	}

	if result == nil && node.Right != nil {
		result = node.Right.Find(value)
	}

	return result
}

// RotateLeft the tree 
func (node *Node) RotateLeft() *Node {
	if node == nil {
		return nil
	}

	if node.Right == nil {
		return node
	}

	var result = node.Right

	node.Right = result.Left
	result.Left = node

	var LeftBalance = node.Balance
	var Balance = result.Balance

	node.Balance = LeftBalance - 1 - max(Balance, 0)
	result.Balance = min(LeftBalance-2, Balance+LeftBalance-2, Balance-1)

	return result
}

// RotateRight the tree 
func (node *Node) RotateRight() *Node {
	if node == nil {
		return nil
	}

	if node.Left == nil {
		return node
	}

	var result = node.Left
	node.Left = result.Right
	result.Right = node

	var RightBalance = node.Balance
	var Balance = result.Balance

	node.Balance = RightBalance + 1 - min(Balance, 0)
	result.Balance = max(RightBalance+2, Balance+RightBalance+2, Balance+1)

	return result
}

// Insert a node
func (node *Node) Insert(value int) (*Node, int) {

	// Terminal Condition, create this node
	if node == nil {
		return &Node{Value: value, Balance: 0}, 1
	}

	var change int

	// Descend to the children
	diff := node.Compare(value)
	switch {

	case diff == 0:
		// Ignore duplicates
		fmt.Println("cannot have 2 values the same")

	case diff > 0:
		node.Left, change = node.Left.Insert(value)
		change *= -1

	case diff < 0:
		node.Right, change = node.Right.Insert(value)
	}

	node.Balance += change

	// ReBalance at the parents or grandparents
	var insert int

	if node.Balance != 0 && change != 0 {
		switch {

		case node.Balance < -1:
			if node.Left.Balance >= 0 {
				node.Left = node.Left.RotateLeft()
			}
			node = node.RotateRight()
			insert = 0

		case node.Balance > 1:
			if node.Right.Balance <= 0 {
				node.Right = node.Right.RotateRight()
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
		node.Left, change = node.Left.Delete(value)

	case diff < 0:
		node.Right, change = node.Right.Delete(value)
		change *= -1

	case diff == 0:
		switch {
		case node.Left == nil:
			return node.Right, 1

		case node.Right == nil:
			return node.Left, 1

		default:
			// Pick the heavier of the two...
			if -1*node.Left.Balance < node.Right.Balance {
				node = node.RotateLeft()
				node.Left, change = node.Left.Delete(value)

			} else {
				node = node.RotateRight()
				node.Right, change = node.Right.Delete(value)
				change *= -1
			}
		}
	}

	// Update the Balance
	if change != 0 {

		if node.Balance != change {
			node.Balance += change
		}

		switch {
		case node.Balance < -1:
			if node.Left.Balance >= 0 {
				node.Left = node.Left.RotateLeft()
			}
			node = node.RotateRight()

		case node.Balance > 1:
			if node.Right.Balance <= 0 {
				node.Right = node.Right.RotateRight()
			}
			node = node.RotateLeft()
		}
	}

	return node, change
}

// Print the current nodes, rotated 90 degrees, in-order traversal.
func (node Node) Print(depth int) {

	if node.Right != nil {
		node.Right.Print(depth + 1)
	}

	padding := padding(depth)
	fmt.Printf("%s[%v] Value: %v\n", padding, symbols(node.Balance), node.Value)

	if node.Left != nil {
		node.Left.Print(depth + 1)
	}
}

// Slightly more decorative: shift 1/-1 -> +/-
func symbols(Balance int) string {
	switch {
	case Balance == 1:
		return "+"
	case Balance == -1:
		return "-"
	}
	return strconv.Itoa(Balance)
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
	t.Traverse(n.Left, f)
	f(n)
	t.Traverse(n.Right, f)
}


/*
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
	
	//tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.value, " | ") })
	
	for i := 50; i < 75; i++ {
		tree.Delete(i)
	}
	
	//tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.value, " | ") })
	
	//tree.Exists(55)
	
	//tree.Print()
	//indexStore["test"] = tree
	fmt.Println(indexStore["test"].Exists(60))
	fmt.Println(indexStore["test2"].Exists(10))
}

*/