package main

import (
	"fmt"
	"sort"
	"testing"
	"time"
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

func sortedKeys(m map[int]string) []int {
	timeSplit := time.Now()
	sm := new(sortedMap)
	sm.m = m
	sm.s = make([]int, len(m))
	i := 0
	for key, _ := range m {
		sm.s[i] = key
		i++
	}
	sort.Sort(sm)
	fmt.Println("Sort by Key:", time.Since(timeSplit))
	return sm.s
}

func sortByValue(m map[int]string) PairList {
	timeSplit := time.Now()
	pl := make(PairList, len(m))
	i := 0
	for k, v := range m {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(pl)
	fmt.Println("Sort by Value:", time.Since(timeSplit))
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

func main() {

	for i := 0; i < 100; i++ {

		// var slice []string
		m := make(map[int]string)

		for i := 0; i < 100; i++ {
			// slice = append(slice, randomdata.Address())
			m[i] = time.Now().String()

		}

		// timeStart := time.Now()
		// sort.Strings(slice)
		// fmt.Println("Sort Slice by Key:", time.Since(timeStart))

		// timeSplit := time.Now()
		sortedKeys(m)
		sortByValue(m)

		if i == 99 {
			export = m
		}
	}

	// fmt.Println(sortedKeys(export))
	// fmt.Println(sortByValue(export))
	timeStart := time.Now()
	x := export[13]
	fmt.Println(x, time.Since(timeStart))

}

func BenchmarkSetIntMap(b *testing.B) {
	var found bool
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, ok := export[i]; ok {
			found = true
		}
	}
	if !found {
		b.Fail()
	}
}
