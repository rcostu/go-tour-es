package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Respuesta"] = 42
	fmt.Println("Valor:", m["Respuesta"])

	m["Respuesta"] = 48
	fmt.Println("Valor:", m["Respuesta"])

	delete(m, "Respuesta")
	fmt.Println("Valor:", m["Respuesta"])

	v, ok := m["Respuesta"]
	fmt.Println("Valor:", v, "¿Existe?", ok)
}
