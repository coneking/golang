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

