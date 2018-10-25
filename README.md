<p align="center"><img src="https://raw.githubusercontent.com/coneking/golang/develop/images/gogo.gif" width="300" /></p>

<center> <h1>Golang</h1> </center><br>

- [Instalación](#instalacion)
- [Hola Mundo](#hola-mundo)
- [Ejecuciones](#ejecuciones)
- [Ayudas](#ayudas)
- [Variable](/Variables.md)
- [Constantes](/Constantes.md)
- [Struct](/Struct.md)

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

```go
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

# Ejecuciones

## Ejecutar el programa

Con Go se puede ejecutar un programa de dos formas; compilando y sin compilar.

### Ejecutar sin compilar

```sh
$ go run hello.go
Hola Mundo!!!
```

<br>

### Ejecutar compilando

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

# Ayudas

## Ayuda de formato

Go nos da una mano en el formateo de código, esto es por si tenemos un programa poco legible, de esta forma se estandarizará para una mejor lectura. <br>

### Ejemplo código sin formato

```go
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

***

<br>

Más información sobre paquetes, variables, tipos de datos, etc. En el [Tour de GO](https://tour.golang.org/basics/1)
