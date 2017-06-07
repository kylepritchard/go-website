package main

import "github.com/datastream/btree"

func main() {

	bt := btree.NewBtree()

	//Insert

	bt.Insert([]byte{1}, []byte{"K", "y", "l", "e"})
	bt.Insert("2", "Robyn")
	bt.Marshal("tree.dump")
}
