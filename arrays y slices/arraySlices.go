//como trabajan los arreglos y slices en go
//cuando definimos un arrgelo en go tenemos que definirle un tamaño fijo
package main
import "fmt"

func main() {
/*arrgelos*/
var listaFrutas= [5]string{"Pera", "naranja"} //definimos un arreglo de 5 elementos
fmt.Println("Arreglo de frutas: ", listaFrutas) //imprimimos el arreglo vacio
fmt.Println(listaFrutas[0])

listarPaises := [4]string{"Colombia", "Peru", "Chile"} //definimos un arreglo de 3 elementos
listarPaises[3] = "Argentina" //modificamos el primer elemento del arreglo
fmt.Println("Arreglo de paises: ", listarPaises) //imprimimos el arreglo de paises

//los slices son arreglos dinamicos, no tienen un tamaño definido
var listaColores = []string{"Rojo", "Verde", "Azul"}
fmt.Println("Slice de colores: ", listaColores)
//agrenagdo datos a los arreglos
listaColores= append(listaColores, "Amarillo") //agregamos un elemento al slice
fmt.Println("Slice de colores despues de agregar un elemento: ", listaColores) //imprimimos el slice de colores



//como obtener el tamaño de un arreglo o slice
listarPaises2 := listarPaises[1:3]
fmt.Println("Slice de paises: ", listarPaises2) //imprimimos el slice de paises
listarPaises3 := listarPaises[1:] //desde uno hasta el final
fmt.Println("Slice de paises desde el segundo elemento: ", listarPaises3) 
listarPaises4 := listarPaises[:3] //desde el inicio hasta el tercer elemento
fmt.Println("Slice de paises hasta el tercer elemento: ", listarPaises4)
}