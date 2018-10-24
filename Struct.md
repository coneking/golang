# Creando Tipos de Datos (Struct)

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
