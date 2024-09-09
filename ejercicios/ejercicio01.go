package ejercicios

import (
	"strconv"
)

// Ejercicio para convertir un string a entero
func DevuelveEntero(valor string) (int, string) {
	numero, err := strconv.Atoi(valor) //Para ignorar algo, osea si una función retorna un valor que no nos importa, se puede utilizar el _
	if err != nil {
		return 0, "Hubo un error: " + err.Error()
	}
	if numero > 100 {
		return numero, "Es mayor a 100"
	} else if numero == 100 {
		return numero, "Es igual a 100"
	} else {
		return numero, "Es menor a 100"
	}

}
