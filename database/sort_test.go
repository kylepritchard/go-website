package main

import (
	"sort"
	"testing"
	"time"

	"github.com/tidwall/btree"
)

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

func bubbleSort(m map[int]string) (map[int]string, bool) {
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
	return m, true
}

func BenchmarkTest(b *testing.B) {
	var found bool
	b.ResetTimer()
	q := PopulateMap()
	for i := 0; i < b.N; i++ {
		if _, ok := q[i]; ok {
			found = true
		}
	}
	if !found {
		b.Fail()
	}
}

func BenchmarkTestSortValue(b *testing.B) {
	var found bool
	b.ResetTimer()
	q := PopulateMap()
	for i := 0; i < b.N; i++ {
		sortByValue(q)
		found = true
	}
	if !found {
		b.Fail()
	}
}

func BenchmarkTestSortKey(b *testing.B) {
	var found bool
	b.ResetTimer()
	q := PopulateMap()
	for i := 0; i < b.N; i++ {
		sortByKey(q)
		found = true
	}
	if !found {
		b.Fail()
	}
}

func BenchmarkTestBubbleSortValue(b *testing.B) {
	var found bool
	b.ResetTimer()
	q := PopulateMap()
	for i := 0; i < b.N; i++ {
		_, done := bubbleSort(q)
		found = done
	}
	if !found {
		b.Fail()
	}
}

func BenchmarkTestTreeSortValue(b *testing.B) {
	var found bool

	vals := btree.New(16, "vals")
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
