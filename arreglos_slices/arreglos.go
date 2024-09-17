package arreglos_slices

import "fmt"

//Ejemplo de arreglo estatico en Go
var tabla [10]int = [10]int{11, 19, 69}

//Hacer una matriz en Go
//Una matriz en terminos generales es un arreglo de arreglos
var matriz [10][5]int

func MuestroArreglos() {
	tabla[4] = 45
	tabla[9] = 100
	fmt.Println(tabla)
	tabla2 := [10]string{"Gustavo", "Pablo", "Arthur"}
	tabla2[3] = "Morgan"
	fmt.Println(tabla2)
	//recorrer arreglo con ciclo for
	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[7][4] = 69
	fmt.Println(matriz)
}
