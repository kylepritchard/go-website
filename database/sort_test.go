package main

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
	"time"

	"github.com/tidwall/btree"
)

//tree

type Post struct {
	title    string
	content  string
	id       string
	postdate time.Time
}

func CreatePost(title string, content string, store map[string]Post) bool {
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
	// fmt.Println(result.value)
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
	value   int
	id      string
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

////////////////////////

type sortedMap struct {
	m map[int]string
	s []int
}

func (sm *sortedMap) Len() int {
	return len(sm.m)
}

func (sm *sortedMap) Less(i, j int) bool {
	return sm.m[sm.s[i]] > sm.m[sm.s[j]]
}

func (sm *sortedMap) Swap(i, j int) {
	sm.s[i], sm.s[j] = sm.s[j], sm.s[i]
}

func sortByKey(m map[int]string) []int {
	// timeSplit := time.Now()
	sm := new(sortedMap)
	sm.m = m
	sm.s = make([]int, len(m))
	i := 0
	for key, _ := range m {
		sm.s[i] = key
		i++
	}
	sort.Sort(sm)
	// fmt.Println("Sort by Key:", time.Since(timeSplit))
	return sm.s
}

func sortByValue(m map[int]string) PairList {
	// timeSplit := time.Now()
	pl := make(PairList, len(m))
	i := 0
	for k, v := range m {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(pl)
	// fmt.Println("Sort by Value:", time.Since(timeSplit))
	return pl
}

type Pair struct {
	Key   int
	Value string
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

var export map[int]string

type Item struct {
	Key int
	Val string
}

func (i1 *Item) Less(item btree.Item, ctx interface{}) bool {
	i2 := item.(*Item)
	switch tag := ctx.(type) {
	case string:
		if tag == "vals" {
			if i1.Val < i2.Val {
				return true
			} else if i1.Val > i2.Val {
				return false
			}
			// Both vals are equal so we should fall though
			// and let the key comparison take over.
		}
	}
	return i1.Key < i2.Key
}

// func main() {

// 	for i := 0; i < 100; i++ {

// 		// var slice []string
// 		m := make(map[int]string)

// 		for i := 0; i < 100; i++ {
// 			// slice = append(slice, randomdata.Address())
// 			m[i] = time.Now().String()

// 		}

// 		// timeStart := time.Now()
// 		// sort.Strings(slice)
// 		// fmt.Println("Sort Slice by Key:", time.Since(timeStart))

// 		// timeSplit := time.Now()
// 		sortedKeys(m)
// 		sortByValue(m)

// 		if i == 99 {
// 			export = m
// 		}
// 	}

// 	// fmt.Println(sortedKeys(export))
// 	// fmt.Println(sortByValue(export))
// 	timeStart := time.Now()
// 	x := export[13]
// 	fmt.Println(x, time.Since(timeStart))

// }

func PopulateMap() map[int]string {
	m := make(map[int]string)
	for i := 0; i < 100; i++ {
		// slice = append(slice, randomdata.Address())
		m[i] = time.Now().String()
	}
	return m
}

func bubbleSort(m map[int]string) map[int]string {
	// n is the number of items in our list
	n := len(m)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n-1; i++ {
			if m[i-1] > m[i] {
				// fmt.Println("Swapping")
				// swap values using Go's tuple assignment
				m[i], m[i-1] = m[i-1], m[i]
				swapped = true
			}
		}
	}
	return m
}

// func BenchmarkTestTreeReadValue(b *testing.B) {
// 	var found bool

// 	vals := btree.New(32, "vals")
// 	for i := 0; i < 100; i++ {
// 		vals.ReplaceOrInsert(&Item{Key: i, Val: "value"})
// 	}

// 	b.ResetTimer()
// 	// q := PopulateMap()
// 	for i := 0; i < b.N; i++ {
// 		vals.Get({Key: i, Val: "value"})
// 		found = true
// 	}
// 	if !found {
// 		b.Fail()
// 	}
// }

func BenchmarkTestReadMap(b *testing.B) {
	var found bool
	b.ResetTimer()
	q := PopulateMap()
	for i := 0; i < b.N; i++ {
		_, ok := q[i]
		if ok {
			found = true
		}

		// if q[i] != "" {
		// 	found = true
		// }
	}
	if !found {
		b.Fail()
	}
}

func BenchmarkTestInsertMap(b *testing.B) {
	// var found bool
	m := make(map[int]string)
	b.ResetTimer()
	// q := PopulateMap()
	for i := 0; i < b.N; i++ {
		m[i] = "abc"
	}
	// if !found {
	// 	b.Fail()
	// }
}

func BenchmarkTestDeleteMap(b *testing.B) {
	// var found bool
	// m := make(map[int]string)

	b.ResetTimer()
	q := PopulateMap()
	for i := 0; i < b.N; i++ {
		delete(q, i)
	}
	// if !found {
	// 	b.Fail()
	// }
}

func BenchmarkTestMapSortKey(b *testing.B) {
	q := PopulateMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sortByKey(q)
	}
}

func BenchmarkTestSortByValue(b *testing.B) {
	// var found bool
	q := PopulateMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sortByValue(q)
		// found = true
	}
	// if !found {
	// 	b.Fail()
	// }
}

func BenchmarkTestBubbleSortValue(b *testing.B) {
	// var found bool
	q := PopulateMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bubbleSort(q)
		// found = done
	}
	// if !found {
	// 	b.Fail()
	// }
}

func BenchmarkTestTreeInsertValue(b *testing.B) {
	var found bool

	vals := btree.New(32, "vals")
	// for i := 0; i < 100; i++ {
	// 	vals.ReplaceOrInsert(&Item{Key: i, Val: time.Now().String()})
	// }

	b.ResetTimer()
	// q := PopulateMap()
	for i := 0; i < b.N; i++ {
		vals.ReplaceOrInsert(&Item{Key: i, Val: time.Now().String()})
		found = true
	}
	if !found {
		b.Fail()
	}
}

func BenchmarkTestTreeSortValue(b *testing.B) {
	var found bool

	vals := btree.New(32, "vals")
	for i := 0; i < 100; i++ {
		vals.ReplaceOrInsert(&Item{Key: i, Val: time.Now().String()})
	}

	b.ResetTimer()
	// q := PopulateMap()
	for i := 0; i < b.N; i++ {
		vals.Ascend(func(item btree.Item) bool {
			kvi := item.(*Item)
			// fmt.Printf("%s %s\n", kvi.Key, kvi.Val)
			if kvi != nil {
				found = true
			}
			return true
		})
	}
	if !found {
		b.Fail()
	}
}

var AVL Tree

func BenchmarkTestAVLTreeInsertValue(b *testing.B) {
	var tree Tree
	b.ResetTimer()
	// q := PopulateMap()
	for i := 0; i < b.N; i++ {
		tree.Insert(i)

		if i == b.N-1 {
			// fmt.Println(b.N)
			AVL = tree
		}
	}
	// if !found {
	// 	b.Fail()
	// }
}

func BenchmarkTestAVLTreeReadValue(b *testing.B) {

	var tree Tree
	for i := 0; i < b.N; i++ {
		tree.Insert(i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tree.Exists(i)
	}

}
