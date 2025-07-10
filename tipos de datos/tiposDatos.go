package main

//declaracion de variables en go.
//es un lenguaje de tipado fuerte pero flexible.
import "fmt"

func main() {
	var nombrePersona string ="Juan Perez"
	var apelidoPersona  = "Lopez"
	segundoNombre := "Maria" //declaracion de variable sin tipo de dato, go infiere el tipo de dato
	fmt.Println("Hola mundo. Soy "+ nombrePersona)
	fmt.Println("Mi apellido es " + apelidoPersona)
	fmt.Println("Mi segundo nombre es " + segundoNombre)
/*Parte numerica */

var asoActual int16 = 2025
var asoReducido int8 = 127
edad:=29
fmt.Println("El año actual es ", asoActual)
fmt.Println("Mi edad es ", edad)
fmt.Println("El año reducido es ", asoReducido)
}