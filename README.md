<p align="center"><img src="https://raw.githubusercontent.com/coneking/golang/develop/images/gogo.gif" width="300" /></p>

# Golang

## Instalación

Para realizar la instalación de Go en Linux seguimos las siguientes [Instrucciones](https://golang.org/doc/install?download=go1.11.1.linux-amd64.tar.gz) para su versión más estable (1.11.1).
<br>

## Hola mundo

Una vez instalado y configurado, vamos a crear nuestro primer código: <br>

Creamos nuestro directorio de trabajo

```sh
mkdir -p $HOME/go/src/prueba
cd $HOME/go/src/prueba
```
<br>

Creamos un archivo `hello.go` con la siguiente información

```sh
package main

import "fmt"

func main() {
	fmt.Println("Hola Mundo!!!")
}
```

- **package**: hace referencia al nombre del paquete donde vivirá el código, en Go todo programa comienza su ejecución en el paquete `main`.
- **import**: como lo dice su nombre importa un paquete que ya existe, en este caso estamos importando el paquete `fmt` que trae funciones para darle formato a las entradas y salidas.
- **func**: es la declaración de una función, en este caso es una funcion main() que posterior a la función init() es la que sigue en su ejecución.
- **fmt.Println**: llama al método Println (homogolo al print de python o al System.Out.Println de Java) que imprime el contenido en la salida estándar del sistema (stdout).

<br>

## Ejecutar el programa

Con Go se puede ejecutar un programa de dos formas; compilando y sin compilar.

### Ejecutar sin compilar

```sh
$ go run hello.go
Hola Mundo!!!
```

<br>

### Ejeutar compilando

```sh
$ go build hello.go 
$ ls -l
total 1864
-rwxrwxr-x 1 user user 1902849 oct 15 20:53 hello
-rw-rw-r-- 1 user user      74 oct 15 20:48 hello.go
$ ./hello 
Hola Mundo!!!
```
>Esto creará un binario que posteriormente podremos ejecutar

<br>

## Definiendo Variables

Para definir una variable usamos la palabra reservada `var` seguido de un nombre para la variable y el tipo de dato. <br>
Creamos un archivo `variable.go` con el siguiente contenido

```sh
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

```
$ go run variables.go 
Tengo que ser siempre el mejor, mejor que nadie más
4
true 10.1 Variables 
La suma de 2 + 2 es: 4
```
<br>

Explicación:
<br>

- Variable `pokemon`, el tipo de dato se declara como `string` con su respectivo valor.
- Variable `suma` el tipo de dato se decalra como `int` con un valor numérico (una suma de dos números).
- Variables `a, b, c` se declaran sin la palabra reservada `var` pero se le agrega `:=`, con esto Go adivinará el tipo de dato de la variable sin especificarlo.
	- Variable **a**: Go intuye que es de tipo `bool`
	- Variable **b**: Go intuye que se de tipo `float`
	- Variable **c**: Go intuye que se de tipo `string`
- Variable `concat` se compone de una concatenación entre un texto y una variable de tipo `int` y su salida será un string.


<br>

## Ayudas sobre librerías

Los paquetes en Go son muchos y si bien existe la documentación en [internet](https://golang.org/pkg/) también podemos ayudarnos del comando `godoc` seguido de uno de los paquetes a consultar para que nos muestre un manual de uso.

```sh
$ godoc fmt

PACKAGE DOCUMENTATION

package fmt
    import "fmt"

    Package fmt implements formatted I/O with functions analogous to C's
    printf and scanf. The format 'verbs' are derived from C's but are
    simpler.


    Printing

    The verbs:

    General:

	%v	the value in a default format
...
```

<br>

## Ayuda de formato

Go nos da una mano en el formateo de código, esto es por si tenemos un programa poco legible, de esta forma se estandarizará para una mejor lectura. <br>

### Ejemplo código sin formato

```sh
package main
import "fmt"
func main() {   var pokemon string = "mejor que nadie más" 
var suma int = 2 + 2
a, b, c := true, 10.1, "Variables "
   concat := fmt.Sprintf("La suma de 2 + 2 es: %d", suma) 
    fmt.Println("Tengo que ser siempre el mejor, " + pokemon) 
fmt.Println(suma) 
            fmt.Println(a, b, c)
            fmt.Println(concat)
}
```
Como pueden ver este ejemplo a pesar de funcionar, no es muy amigable a la vista pero ejecutando el comando `go fmt` seguido del archivo a dar formato se verá de la siguiente forma.

<br>

```sh
$ go fmt sin_formato.go
sin_formato.go
$ cat sin_formato.go 
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

## Creando Tipos de Datos (Struct)

En Go es posible crear un tipo de dato personalizado, esto quiere decir que podemos crear un "objeto" que es una colecciones de campos con tipos específicos. Se inicializa con la palabra reservada `type` seguido del nombre del tipo de dato a crear y la palabra reservada `struct`. <br>
En su interior deben ir los campos con su respectivo tipo de dato. <br>

Ejemplo: 

```
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
```
Posteriormente podemos utilizarlo como un archivo json de las siguiente forma. <br>

### Nombrando cada campo y añadiendo un valor
```
func main(){
    fmt.Println(pokemon{tipo: "Fuego", numero: 6, color: "Rojo", nombre: "Charizard", altura: 1.70, peso: 90.5, habilidad: "Mar llamas", debilidad: "Agua"})
}

$ go run estructuras.go 
{Charizard Fuego 6 Rojo 1.7 90.5 Mar llamas Agua}
```
>**Nota:** El orden que se ingrese el campo con su valor no importa, esto porque retornará el orden en como se definieron los campos en el struct.

<br>

### Sólo añadiendo el valor de cada campo

```
func main(){
    fmt.Println(pokemon{"Charizard", "Fuego", 6, Rojo", 1.70, 90.5, "Mar llamas", "Agua"})
}

$ go run estructuras.go 
{Charizard Fuego 6 Rojo 1.7 90.5 Mar llamas Agua}
```
>**Nota:** Esto es válido si se conoce el orden de cada campo, ya que, si se asignara un valor de tipo `int` al primer campo daría un error porque se declaró como `string`. 

<br>

### Usando variables y nomenclatura de punto

```
func main(){
var pokemon = pokemon {nombre: "Charizard", tipo: "Fuego", numero: 6, color: "Rojo", altura: 1.70, debilidad: "Agua"}
        
        fmt.Println(pokemon)
        fmt.Println(pokemon.nombre + " es de tipo " + pokemon.tipo + " y su número en la pokedex es el", pokemon.numero)
}

$ go run prueba.go 
{Charizard Fuego 6 Rojo 1.7 0  Agua}
Charizard es de tipo Fuego y su número en la pokedex es el 6
```
>**Nota:** Los campos omitidos devolverán un valor cero o nulo.
