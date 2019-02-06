# Funciones

Hasta el momento en los ejemplos que hemos visto, solamente hemos trabajado con la funcion principa (main) y esto porque nuestros programas son muy pequeños y sus ejecuciones son secuenciales. En algún momento tendremos que crear un programa más extenso y ejecutar todo desde la `funcion main` no será conveniente. Para este caso crearemos `funciones` que contendrán ciertas tareas y se ejecutaran solo cuando sean llamadas desde la función principal.<br>
En Go las funciones son declaradas con la palabra reservada `func` seguido del nombre de la función, sus parámetros (si es que son necesarios) y uno o más valores para retornar. <br>

### Ejemplo

```go
func saludoInicial(){
	fmt.Println("Hola Mundo")
}
```
Este bloque de código no será ejecutado hasta que lo llamemos en nuestra función main. <br>

<br>

```go
package main

import "fmt"

func main (){
	saludoInicial()
}


func saludoInicial(){
	fmt.Println("Hola Mundo")
}
```
Al ejecutarse mostrará el mensaje "Hola Mundo"

```sh
$ go run funciones.go 
Hola Mundo
```
>**Nota:** Si una función es llamada más de una vez dentro de la funcion main, está devolverá su resultado las mismas veces que es invocada.

<br>

---

# Programa básico con funciones

En este ejemplo crearemos un programa en base a funciones el cual permita sumar, restar, dividir y multiplicar dos números (posteriormente se dará una explicación de lo que hace la función).<br>
Lo primero será crear una funcion para sumar.

```go
func sumar(numero1 float32, numero2 float32) float32{
	return(numero1 + numero2)
}
```
### Explicación 1

Creamos la función con la palabra reservada `func` y le damos el nombre `sumar`, entre paréntesis indicamos que esta función recibirá dos parámetros de tipo `float32`, los que se declaran con las variables `numero1` y `numero2`, posterior a los paréntesis indicamos que la función devolverá un valor de tipo `float32`. Finalmente con la palabra reservada `return` se indica que la función retornará la suma de las variables numero1 y numero2. <br>
Haremos el mismo ejercicio para declarar las funciones restar, dividir y multiplicar.

```go
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
```
<br>

Ahora en nuestra funcion main agregaremos lo siguiente.

```go
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
```

### Explicación 2

En la función main se declaran dos variables llamadas `num1` y `num2` de tipo `float32` con su respectivo valor. Se crean cuatro print los que llaman a cada una de las funciones (sumar, restar, dividir, multiplicar), y se les pasan los parámetros `num1` y `num2`. <br>

## Resultado

```go
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

```

```sh
$ go run calculadora.go 
Sumando 8 + 5 = 13
Restando 8 - 5 = 3
Dividiendo 8 / 5 = 1.6
Multiplicando 8 * 5 = 40
```
<br>

---

# Mejorando el programa

Podemos mejorar el programa para leerlo de una manera más sencilla con menos funciones.

## Resultado

```go
package main

import "fmt"

func main(){
    var num1 float32 = 8
	var num2 float32 = 5

	salida(num1, num2)
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
```

```sh
$ go run calculadora.go 
Sumando 8 + 5 = 13
Restando 8 - 5 = 3
Dividiendo 8 / 5 = 1.6
Multiplicando 8 * 5 = 40
```

### Explicación

En este ejemplo se crean dos funciones, la primera llamada `calcular` recibe tres parámetros, dos de tipo `float32` y un `string` y retornará un valor tipo `float32`. En su interior evaluará si el parámetro `opcion` es resta (-), suma (+), división (/) o multiplicación (\*) y en base a un condicional (if) retornará, la suma, resta, división o multiplicación de los valores. <br>
La segunda función llamada `salida` recibe dos parámetros de tipo `float32` y no tiene retorno. Los parámetros que reciba los pasará a la función `calcular` más el símbolo de la operación (+,-,/,\*). <br>
Finalmente en la función principal se declaran las mismas dos variables de tipo float32 con su respectivo valor y se llama a la función `salida`, pasando los dos parámetros que solicita.

---

# Reutilización de código

Una de las principales ventajas de las funciones es que se pueden reutilizar sin necesidad de tener que volver a crear código que haga un mismo trabajo. En el ejemplo anterior de la calculadora, si fuese necesario calcular otros dos número bastaría con volver a llamar la función `salida` y pasarle otros parámetros.

```go
package main

import "fmt"

func main(){
    var num1 float32 = 8
	var num2 float32 = 5
	var num3 float32 = 20
	var num4 floar32 = 2

	salida(num1, num2)
	fmt.Println("\n")
	salida(num3, num4)
}
```

Con esto obtendremos dos resultados sin necesidad de volver a crear nuevas funciones.

```sh
$ go run calculadora2.go 
Sumando 8 + 5 = 13
Restando 8 - 5 = 3
Dividiendo 8 / 5 = 1.6
Multiplicando 8 * 5 = 40


Sumando 20 + 2 = 22
Restando 20 - 2 = 18
Dividiendo 20 / 2 = 10
Multiplicando 20 * 2 = 40
```
