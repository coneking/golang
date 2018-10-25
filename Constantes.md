# Definiendo Constantes

Las constantes, distintas de las variables, son contenedores que almacenan valores que se asignan de manera fija, su valor no puede ser modificado dentro del código, ya que, como su nombre lo dice es `constante` y una variable puede tener un valor mutable.<br>
- Pueden contener los mismos tipos de datos que una variable. <br>
- No pueden declararse usando `:=` como lo hace una variable. <br>
- Par definir una constante usamos la palabra reservada `const` de la siguiente manera:

```
package main

import "fmt"

func main() {
        const elemento string = "fuego"
        fmt.Println("Charizard es un pokemon de tipo", elemento)

}
```
<br>

Ejecutamos

```
$ go run constante.go
Charizard es un pokemon de tipo fuego
```

<br>

Si intentamos modificar el valor de la constante, nuestro código fallará

```
package main

import "fmt"

func main() {
        const elemento string = "fuego"
        fmt.Println("Charizard es un pokemon de tipo", elemento)
    

        //Intento modificar el valor de la constante elemento
        elemento = "agua"
        fmt.Println("Blastoise es un pokemon de tipo", elemento)
    
}
```

<br>

Ejecutamos

```
$ go run constante.go 
# command-line-arguments
./constante.go:11:11: cannot assign to elemento
```
>Como se explicó, el valor de una constante no puede ser modificado.

<br>

***

<br>

Más información sobre paquetes, constantes, tipos de datos, etc. En el [Tour de GO](https://tour.golang.org/basics/1)

<br>

[Inicio](/README.md)