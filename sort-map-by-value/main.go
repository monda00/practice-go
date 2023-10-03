package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func printKeyOrderByValue(targetMap map[string]int) {
	pl := make(PairList, len(targetMap))
	i := 0
	for k, v := range targetMap {
		pl[i] = Pair{k, v}
		i++
	}
	for _, v := range pl {
		fmt.Printf("%s: %d\n", v.Key, v.Value)
	}
	sort.Sort(pl)
	//sort.Sort(sort.Reverse(pl))

	for _, v := range pl {
		fmt.Printf("%s: %d\n", v.Key, v.Value)
	}
}

func main() {
	targetMap := map[string]int{
		"five":  5,
		"nine":  9,
		"three": 3,
		"four":  4,
		"six":   6,
		"two":   2,
		"ten":   10,
		"seven": 7,
		"one":   1,
		"eight": 8,
	}

	printKeyOrderByValue(targetMap)
}
