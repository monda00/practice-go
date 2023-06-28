package main

import "fmt"

var x int = 10

// x int := 10

func main() {
	fmt.Println(x)

	var y = [3]int{1, 2, 3}
	fmt.Println(y)

	z := []int{1, 2, 3}
	z = append(z, 10)
	fmt.Println(z)

	s := "Hello"
	var st string = "hello"
	fmt.Println(s)
	fmt.Println(st)

	totalWins := map[string]int{}
	fmt.Println(totalWins["abc"])
	totalWins["abc"] = 3
	fmt.Println(totalWins["abc"])

	type persion struct {
		name string
		age  int
		pet  string
	}
	tom := persion{
		"tom",
		40,
		"cat",
	}
	fmt.Println(tom)
}
