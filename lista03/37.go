package main

import "fmt"

func main() {
	arr := []interface{}{1, 2, "apple", true}
	fmt.Println(arr)

	// however, now you need to use type assertion access elements
	i := arr[0].(int)
	fmt.Printf("i: %d, i type: %T\n", i, i)

	s := arr[2].(string)
	fmt.Printf("b: %s, i type: %T\n", s, s)
}
