package main

import (
	"fmt"
	"sort"
)

type fruit struct {
	avgNumSeeds int
	name        string
}

type Fruit interface {
	Name() string
	AvgNumSeeds() int
}

func (f fruit) Name() string {
	return f.name
}

func (f fruit) AvgNumSeeds() int {
	return f.avgNumSeeds
}

type Apple struct {
	fruit
	Diameter int
}

type Banana struct {
	fruit
	Length int
}

type ByNumSeeds []Fruit

func (p ByNumSeeds) Len() int {
	return len(p)
}

func (p ByNumSeeds) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p ByNumSeeds) Less(i, j int) bool {
	return p[i].AvgNumSeeds() < p[j].AvgNumSeeds()
}

func main() {
	apples := []Fruit{
		Apple{fruit: fruit{avgNumSeeds: 4, name: "Cox"}, Diameter: 10},
		Apple{fruit: fruit{avgNumSeeds: 6, name: "Granny Smith"}, Diameter: 20},
		Apple{fruit: fruit{avgNumSeeds: 5, name: "Pink Lady"}, Diameter: 21},
		Apple{fruit: fruit{avgNumSeeds: 2, name: "Russett"}, Diameter: 15},
		Apple{fruit: fruit{avgNumSeeds: 1, name: "Crab"}, Diameter: 7},
		Apple{fruit: fruit{avgNumSeeds: 7, name: "Brambley"}, Diameter: 40},
		Apple{fruit: fruit{avgNumSeeds: 3, name: "Braeburn"}, Diameter: 25},
	}

	bananas := []Fruit{
		Banana{fruit: fruit{avgNumSeeds: 40, name: "Lacatan"}, Length: 20},
		Banana{fruit: fruit{avgNumSeeds: 60, name: "Lady Finger"}, Length: 22},
		Banana{fruit: fruit{avgNumSeeds: 50, name: "Senorita"}, Length: 25},
		Banana{fruit: fruit{avgNumSeeds: 20, name: "Cavendish"}, Length: 30},
		Banana{fruit: fruit{avgNumSeeds: 10, name: "Goldfinger"}, Length: 27},
		Banana{fruit: fruit{avgNumSeeds: 70, name: "Gros Michel"}, Length: 15},
		Banana{fruit: fruit{avgNumSeeds: 30, name: "Red Dacca"}, Length: 19},
	}

	fmt.Println("Apples")
	fmt.Printf("%+v\n\n", apples)
	sort.Sort(ByNumSeeds(apples))
	fmt.Printf("%+v\n\n\n", apples)

	fmt.Println("Bananas")
	fmt.Printf("%+v\n\n", bananas)
	sort.Sort(ByNumSeeds(bananas))
	fmt.Printf("%+v\n\n", bananas)
}
