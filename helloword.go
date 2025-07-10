package main //paquete principal de go para que reconozca el archivo como un ejecutable

import "fmt" //importamos el paquete fmt para poder usar la funcion println

 func main(){
	fmt.Println("Hello World!") 
 }
//para ejecutar esto ponemos go run helloworld.go
//para compilarlo y generar el ejecutable ponemos go build helloworld.go
//para ejecutar el ejecutable generado ponemos ./helloworld
//para compilarlo y generar el ejecutable con un nombre especifico ponemos go build -o helloworld
//para compilarlo y generar el ejecutable con un nombre especifico y ejecutarlo en una sola linea ponemos go run -o helloworld helloworld.go