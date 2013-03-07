package main

import "fmt"

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)
	fmt.Println("p[1:4] ==", p[1:4])

	// Un valor de inicio omitido implica 0
	fmt.Println("p[:3] ==", p[:3])

	// Un valor de fin omitido implica len(s)
	fmt.Println("p[4:] ==", p[4:])
}
