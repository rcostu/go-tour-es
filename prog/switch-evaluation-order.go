package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("¿Cuándo es Sábado?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Hoy.")
	case today + 1:
		fmt.Println("Mañana.")
	case today + 2:
		fmt.Println("Pasado mañana.")
	default:
		fmt.Println("Demasiado tarde.")
	}
}
