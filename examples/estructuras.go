package main

import "fmt"

type pokemon struct {
	nombre string
	tipo string
	numero int
	color string
	altura float32
	peso float32
	habilidad string
	debilidad string
}

func main() {
	var charizard = pokemon{
		nombre: "Charizard",
		tipo: "Fuego",
		numero: 6,
		color: "Rojo",
		altura: 1.7,
		peso: 90.5,
		habilidad: "Mar llamas",
		debilidad: "Agua"}

	fmt.Println(charizard)
	fmt.Println(charizard.nombre + " es de tipo " + charizard.tipo + " y su número en la pokedex es el", charizard.numero)
}
