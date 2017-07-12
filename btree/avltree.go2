package main

import (
	"fmt"
	//"bytes"
	//"strconv"
)

type AVLTree struct {
	Val         int
	LeftHeight  int
	RightHeight int
	Left        *AVLTree
	Right       *AVLTree
}

func Child(t *AVLTree) *AVLTree {
	child := t.Left
	if child == nil {
		child = t.Right
	}
	return child
}

func Insert(t *AVLTree, v int) *AVLTree {
	if t == nil {
		return &AVLTree{v, 0, 0, nil, nil}
	}
	if v < t.Val {
		t.Left = Insert(t.Left, v)
		t.LeftHeight++
	} else if v > t.Val {
		t.Right = Insert(t.Right, v)
		t.RightHeight++
	}
	switch t.RightHeight - t.LeftHeight {
	case 2:
		child := Child(t)
		dif := child.RightHeight - child.LeftHeight
		if dif == 0 || dif == 1 {
			t = RotateLeft(t)
		} else {
			t = RotateChildRight(t)
		}
	case -2:
		child := Child(t)
		dif := child.RightHeight - child.LeftHeight
		if dif == 0 || dif == -1 {
			t = RotateRight(t)
		} else {
			t = RotateChildLeft(t)

		}
	}
	return t
}

func RotateLeft(t *AVLTree) *AVLTree {
	c := t.Right
	c.Left = t
	c.LeftHeight = 1
	c.RightHeight = 1
	c.Right.LeftHeight = 0
	c.Right.RightHeight = 0
	c.Left.LeftHeight = 0
	c.Left.RightHeight = 0
	t.Right = nil
	t.Left = nil
	return c
}
func RotateRight(t *AVLTree) *AVLTree {
	c := t.Left
	c.Right = t
	c.RightHeight = 1
	c.LeftHeight = 1
	c.Right.LeftHeight = 0
	c.Right.RightHeight = 0
	c.Left.LeftHeight = 0
	c.Left.RightHeight = 0
	t.Right = nil
	t.Left = nil
	return c
}
func RotateChildLeft(t *AVLTree) *AVLTree {
	c := t.Left
	t.Left = c.Right
	c.Right.Left = c
	c.Right = nil
	c.Left = nil
	return RotateRight(t)
}
func RotateChildRight(t *AVLTree) *AVLTree {
	c := t.Right
	t.Right = c.Left
	c.Left.Right = c
	c.Right = nil
	c.Left = nil
	return RotateLeft(t)
}

func main() {
	var t *AVLTree
	t = Insert(t, 6)
	t = Insert(t, 3)
	t = Insert(t, 5)
	t = Insert(t, 15)
	t = Insert(t, 52)
	t = Insert(t, 35)
	t = Insert(t, 0)
	t = Insert(t, -5)
	t = Insert(t, 51)

	fmt.Println(t, t.Left, t.Right)
}
