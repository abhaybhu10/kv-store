package main

import "fmt"

type functor func(a, b int)

func main() {
	functions := map[string]functor{}
		
	functions["add"] = func(a, b int) {
		fmt.Printf("Sum :%d\n", a+b)
	}
	functions["mult"] = func(a, b int) {
		fmt.Printf("product :%d\n", a*b)
	}


	functions["mult"](2, 4)
	functions["add"](2, 4)
}
