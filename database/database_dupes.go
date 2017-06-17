package main

import (
    "fmt"
    "errors"
    _ "log"
    _ "os"
    _ "encoding/gob"
)

var (
		ERROR_NODE_NOT_IMPLEMENTED = errors.New("Error: Node is not implemented yet")
		ERROR_INSERT_NIL_TREE = errors.New("Error: Cannot insert into a nil tree")
		ERROR_VALUE_EXISTS_ALREADY = errors.New("Error: The value exists already in the tree")
)

type DupeNode struct {
			Value		string
			ID					string
}

type Node struct {
    Value		string
    ID				string
    Left			*Node
    Right		*Node
    Dupe			bool
    Dupes		[]DupeNode 
}

// Insert method for Node
// Returns error if the node is nil
func (n *Node) Insert(value string, id string) error {
			
			// If node is nil then return the error
			if n == nil {
				return ERROR_INSERT_NIL_TREE
			}
			
			switch {
					
					// check if value exists add it to the duplicates slice			
					case value == n.Value :
						var dupes []DupeNode
						dupes = append(dupes, DupeNode{Value: value, ID: id})
						n.Dupes = dupes
						n.Dupe = true
						return nil
					
					case value < n.Value :
						if n.Left == nil {
							n.Left = &Node{Value: value, ID: id}
							return nil
						}
						return n.Left.Insert(value, id)
					
					case value > n.Value :
						if n.Right == nil {
							n.Right = &Node{Value: value, ID: id}
							return nil
						}
						return n.Right.Insert(value, id)
			
			}
			
			return nil
}


// Find method for Node
// Returns the ID and a boolean

func (n *Node) Find(value string) ([]string, bool) {
				
				var ids []string
				
				// if node is nil then return blank string and false
				if n == nil {
					return ids, false
				}
				
				switch {
						
						// found a match :)
						case value == n.Value :
								if n.Dupe == true {
										ids = append(ids, n.ID)
										for _, node := range n.Dupes {
												ids = append(ids, node.ID)
										}
										return ids, true
								}
								ids = append(ids, n.ID)
								return ids, true
						
						// no match recurse over smaller values	to left	
						case value < n.Value:
								return n.Left.Find(value)
								
						// no match, recurse over bigger values to right
						case value > n.Value:
								return n.Right.Find(value)
				}
				return ids, false
}


// TREE

type Tree struct {
		Root *Node
}

func (t *Tree) Insert(value string, id string) error {
		
		// if no root then make a root node
		if t.Root == nil {
				t.Root = &Node{Value: value, ID: id}
				return nil
		}
		//Insert data to node pointed by Root
		return t.Root.Insert(value, id)
}

func (t *Tree) Find(value string) ([]string, bool) {
		var ids []string
		if t.Root == nil {
				return ids, false
		}
		return t.Root.Find(value)
}

func (t *Tree) Traverse(n *Node, f func(*Node)) {
	if n == nil {
		return
	}
	t.Traverse(n.Left, f)
	f(n)
	t.Traverse(n.Right, f)
}

func main(){
	fmt.Println("hello")
	
	tree := &Tree{}
	// Set up a slice of strings.
	id := []string{"d", "b", "a", "e", "c", "f"}
	values := []string{"delta", "bravo", "alpha", "echo", "charlie", "alpha"}

	// Create a tree and fill it from the values.
	for i := 0; i < len(id); i++ {
		err := tree.Insert(values[i], id[i])
		if err != nil {
			fmt.Println("Error inserting value '", values[i], "': ", err)
		}
	}
	
	/*
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.ID, " | ") })
	
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
	
	tree.Traverse(tree.Root, func(n *Node) { fmt.Println(n.Value, ": ", n.ID) })
	slice, _ := tree.Find("alpha")
	for _, v := range slice {
			fmt.Println("IDs that match 'alpha'")
			fmt.Println("ID:", v)
		}

}