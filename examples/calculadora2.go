package main

import "fmt"

func main(){
    var num1 float32 = 8
	var num2 float32 = 5
	var num3 float32 = 20
	var num4 float32 = 2

	salida(num1, num2)
	fmt.Println("\n")
	salida(num3, num4)
}


func calcular(numero1 float32, numero2 float32, opcion string) float32{
	if(opcion == "+"){
		return numero1 + numero2
	}else if(opcion == "-"){
		return numero1 - numero2
	}else if(opcion == "/"){
		return numero1 / numero2
	}else{
		return numero1 * numero2
	}
}


func salida(num1 float32, num2 float32){
	fmt.Println("Sumando", num1, "+", num2, "=", calcular(num1, num2, "+"))
	fmt.Println("Restando", num1, "-", num2, "=", calcular(num1, num2, "-"))
	fmt.Println("Dividiendo", num1, "/", num2, "=", calcular(num1, num2, "/"))
	fmt.Println("Multiplicando", num1, "*", num2, "=", calcular(num1, num2, "*"))
}
