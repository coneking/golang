# Definiendo Variables

Para definir una variable usamos la palabra reservada `var` seguido de un nombre para la variable y el tipo de dato. <br>
Creamos un archivo `variable.go` con el siguiente contenido

```go
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
```
<br>

Ejecutamos:

```sh
$ go run variables.go 
Tengo que ser siempre el mejor, mejor que nadie más
4
true 10.1 Variables 
La suma de 2 + 2 es: 4
```
<br>

## Explicación:

<br>

- Variable `pokemon`, el tipo de dato se declara como `string` con su respectivo valor.
- Variable `suma` el tipo de dato se decalra como `int` con un valor numérico (una suma de dos números).
- Variables `a, b, c` se declaran sin la palabra reservada `var` pero se le agrega `:=`, con esto Go adivinará el tipo de dato de la variable sin especificarlo.
	- Variable **a**: Go intuye que es de tipo `bool`
	- Variable **b**: Go intuye que se de tipo `float`
	- Variable **c**: Go intuye que se de tipo `string`
- Variable `concat` se compone de una concatenación entre un texto y una variable de tipo `int` y su salida será un string.

<br>

# Tipos de Datos

Los tipos de datos en Go son casi lo mismos ya conocidos de otros lenguajes de programación como Python o Java.
Los más comunes son los siguientes:

- string
- bool
- int  int8  int16  int32  int64
- uint uint8 uint16 uint32 uint64 uintptr
- byte // alias for uint8
- rune // alias for int32 
- float32 float64

<br>

***

<br>

Más información sobre paquetes, variables, tipos de datos, etc. En el [Tour de GO](https://tour.golang.org/basics/1)

<br>

[Inicio](/README.md)