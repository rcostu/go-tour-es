package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Ahora tienes %g problemas.",
		math.Nextafter(2, 3))
}
