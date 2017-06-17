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

type Node struct {
    Value		string
    ID				string
    Left			*Node
    Right		*Node 
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
						return ERROR_VALUE_EXISTS_ALREADY
					
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

func (n *Node) Find(value string) (string, bool) {
				
				// if node is nil then return blank string and false
				if n == nil {
					return " ", false
				}
				
				switch {
						
						// found a match :)
						case value == n.Value :
							return n.ID, true
						
						// no match recurse over smaller values	to left	
						case value < n.Value:
								return n.Left.Find(value)
								
						// no match, recurse over bigger values to right
						default:
								return n.Right.Find(value)
				}
}

// `findMax` finds the maximum element in a (sub-)tree. Its value replaces the value of the
// to-be-deleted node.
// Return values: the node itself and its parent node.
func (n *Node) findMax(parent *Node) (*Node, *Node) {
		
		if n.Right == nil {
				return n, parent
		}
		
		return n.Right.findMax(n)
}

// `replaceNode` replaces the `parent`'s child pointer to `n` with a pointer to the `replacement` node.
// `parent` must not be `nil`.
func (n *Node) replaceNode(parent, replacement *Node) error {
	
	if n == nil {
		return errors.New("replaceNode() not allowed on a nil node")
	}

	if n == parent.Left {
		parent.Left = replacement
		return nil
	}
	parent.Right = replacement
	return nil
}

// `Delete` removes an element from the tree.
// It is an error to try deleting an element that does not exist.
// In order to remove an element properly, `Delete` needs to know the node's parent node.
// `parent` must not be `nil`.
func (n *Node) Delete(s string, parent *Node) error {
	
	if n == nil {
		return errors.New("Value to be deleted does not exist in the tree")
	}

	// Search the node to be deleted.
	switch {
	
		case s < n.Value:
			return n.Left.Delete(s, n)
		case s > n.Value:
			return n.Right.Delete(s, n)
		default:
			// We found the node to be deleted.
			// If the node has no children, simply remove it from its parent.
			if n.Left == nil && n.Right == nil {
				n.replaceNode(parent, nil)
				return nil
			}

			// If the node has one child: Replace the node with its child.
			if n.Left == nil {
				n.replaceNode(parent, n.Right)
				return nil
			}
			if n.Right == nil {
				n.replaceNode(parent, n.Left)
				return nil
			}
	
			// If the node has two children:
			// Find the maximum element in the left subtree...
			replacement, replParent := n.Left.findMax(n)
	
			//...and replace the node's value and data with the replacement's value and data.
			n.Value = replacement.Value
			n.ID = replacement.ID
	
			// Then remove the replacement node.
			return replacement.Delete(replacement.Value, replParent)
		}
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

func (t *Tree) Find(value string) (string, bool) {
		if t.Root == nil {
				return " ", false
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

// `Delete` has one special case: the empty tree. (And deleting from an empty tree is an error.)
// In all other cases, it calls `Node.Delete`.
func (t *Tree) Delete(s string) error {

	if t.Root == nil {
		return errors.New("Cannot delete from an empty tree")
	}

	// Call`Node.Delete`. Passing a "fake" parent node here *almost* avoids
	// having to treat the root node as a special case, with one exception.
	fakeParent := &Node{Right: t.Root}
	err := t.Root.Delete(s, fakeParent)
	if err != nil {
		return err
	}
	// If the root node is the only node in the tree, and if it is deleted,
	// then it *only* got removed from `fakeParent`. `t.Root` still points to the old node.
	// We rectify this by setting t.Root to nil.
	if fakeParent.Right == nil {
		t.Root = nil
	}
	return nil
}

// Results functions

type Results []string

func TreeTraverse(t *Tree) (Results) {
	var results Results
	t.Traverse(t.Root, func(n *Node) {
		fmt.Print(n.Value, ": ", n.ID, " | ")
		results = append(results, n.ID)
	})
	fmt.Println()
	return results
}


func main(){
	
	tree := &Tree{}
	// Set up a slice of strings.
	id := []string{"d", "b", "a", "e", "c", "f"}
	values := []string{"delta", "bravo", "alpha", "echo", "charlie", "alpha"}

	// Create a tree and fill it from the values.
	for i := 0; i < len(id); i++ {
		err := tree.Insert(values[i], id[i])
		if err != nil {
			fmt.Println(err, "- '", values[i], "'")
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

	fmt.Println(TreeTraverse(tree))
	
	err := tree.Delete("delta")
	if err != nil {
		fmt.Println("Error deleting", err)
	}
	
	fmt.Println(TreeTraverse(tree))

	gid, found := tree.Find("alpha")
	if found {
		fmt.Println(gid)	
	}
	
}