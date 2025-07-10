package main

import "fmt"

//mapas diccionarios en go
//los mapas son estructuras de datos que nos permiten almacenar pares clave-valor
//cuando utilizamos los mapas, debemos definir el tipo de dato de la clave y el tipo de dato del valor
//tambien si vamos a utilizar los := no tenemos que poner el tipo de dato, go infiere el tipo de dato
//go no garantiza el orden de los elementos en un mapa, por lo que no debemos depender del orden de los elementos
//los mapas son mutables, es decir, podemos agregar, eliminar o modificar elementos en un mapa

func main() {
	//clave string valor string
	miMapa := map[string]string{
		"Colombia":  "Bogota",
		"Peru":      "Lima",
		"Chile":     "Santiago",
		"Argentina": "Buenos Aires",
	}

	fmt.Println("Este es mi mapa:", miMapa)
	fmt.Println("El valor de la clave Colombia es:", miMapa["Colombia"])
	miMapa["Mexico"] = "Ciudad de Mexico" //agregamos un elemento al mapa
	fmt.Println("Mapa despues de agregar un elemento:", miMapa)
	delete(miMapa, "Chile") //eliminamos un elemento del mapa
	fmt.Println("Mapa despues de eliminar un elemento:", miMapa)
}
