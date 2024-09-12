package funciones

import "fmt"

//Ejemplo de recursión
func Exponencia(valor int) {
	if valor > 100000000 {
		return
	}

	fmt.Println(valor)
	Exponencia(valor * 2)
}

//Para que una función sea recursiva o de recursión simplemente debe llamarse a si misma dentro de la función, eso
//desencadenará que se comience un ciclo implicito dentro de la recursión sin necesidad de usar el ciclo for
//También es fundamental definir al comienzo una condicional para detener el ciclo, si no se hace, el ciclo seria
//infinito y no se detendria.
