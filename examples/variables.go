package main

import "fmt"

func main() {
	var pokemon string = "mejor que nadie más"
	var suma int = 2 + 2
	a, b, c := true, 10.1, "Variables "
        concat := fmt.Sprintf("La suma de 2 + 2 es: %d", suma) 
	
	fmt.Println("Tengo que ser siempre el mejor, " + pokemon)
	fmt.Println(suma)
	fmt.Println(a, b, c)
	fmt.Println(concat)
}
