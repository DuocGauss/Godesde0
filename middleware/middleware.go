package middleware

import "fmt"

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func MiMiddleWare() {
	fmt.Println("Inicio")
	//Tiene 2 parametros debido que en la función middleware primero recibe como parametro una función(sumar) y retorna una
	//función anonima que como parametro recibe 2 valores enteros, entonces le envio esos datos(5, 5) a la
	//función anonima la cual esta en operacionesMidd, por eso tiene: operacionesMidd(sumar)(5,5). Luego de recibir
	//esos valores (5, 5) dentro de la función anonima, el middleware agrega nueva logica el cual es un println y le envia
	//esos valores a la función sumar con f(a, b) debido a que esa función esta guardada en esa variable.
	result := operacionesMidd(sumar)(5, 5)
	fmt.Println(result)

	result = operacionesMidd(restar)(5, 3)
	fmt.Println(result)

	result = operacionesMidd(multiplicar)(6, 6)
	fmt.Println(result)
}

//Ejemplo de una función middleware
func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de la operación")
		return f(a, b)
	}
}

//Un middleware es una función intermediaria que sirve para agregar una logica extra a funciones
//ya definidas anteriormente. La logica extra en mi ejemplo fue el Println("Inicio de la operación")
