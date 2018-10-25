package main

import "fmt"

func main() {
        const elemento string = "fuego"
        fmt.Println("Charizard es un pokemon de tipo", elemento)
	

	//Si trato de modificar el valor de la constante el código fallará
	elemento = "agua"
	fmt.Println("Blastoise es un pokemon de tipo", elemento)
	
}
