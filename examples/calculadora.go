package main

import "fmt"

func main(){
	var num1 float32 = 8
	var num2 float32 = 5

	fmt.Println("Sumando", num1, "+", num2, "=", sumar(num1, num2))
	fmt.Println("Restando", num1, "-", num2, "=", restar(num1, num2))
	fmt.Println("Dividiendo", num1, "/", num2, "=", dividir(num1, num2))
	fmt.Println("Multiplicando", num1, "*", num2, "=", multiplicar(num1, num2))

}


func sumar(numero1 float32, numero2 float32) float32{
	return numero1 + numero2
}

func restar(numero1 float32, numero2 float32) float32{
	return numero1 - numero2
}

func dividir(numero1 float32, numero2 float32) float32{
	return numero1 / numero2
}

func multiplicar(numero1 float32, numero2 float32) float32{
	return numero1 * numero2
}
