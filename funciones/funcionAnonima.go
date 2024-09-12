package funciones

import (
	"fmt"
)

func Calculos() {
	//Ejemplo de una función anónima en Go
	var numero3 int = 250
	var numero4 int = 50

	calculo := func(numero1 int, numero2 int) int {
		return numero1 + numero2 + numero3
	}
	fmt.Println("Suma: ", calculo(20, 40))

	calculo = func(numero1 int, numero2 int) int {
		return numero1 * numero2 * numero3 * numero4
	}
	fmt.Println("Multiplicación: ", calculo(25, 3))
}
