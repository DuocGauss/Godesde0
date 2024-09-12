package funciones

import "fmt"

//Ejemplo de una función tipo clousure, en este clousure se realiza una tabla de multiplicar
func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}

func LlamarClousure() {
	numerodel := 5
	mitabla := tabla(numerodel)
	for i := 1; i <= 10; i++ {
		fmt.Println(mitabla())
	}
}
