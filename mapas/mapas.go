package mapas

import "fmt"

//Un mapa es como un diccionario en Python debido a que su sitaxis es similar, el mapa contiene una llave y valor
//Como los diccionarios en Python o objetos en JavaScript
func MostrarMapa() {
	//Se define el mapa y entre corchetes se define el tipo de datos que tendra la llave el cual esta encerrado con []
	//Y también se debe definir el tipo de dato del valor, el cual es string.
	paises := make(map[string]string)

	paises["Chile"] = "Concepción"
	paises["Argentina"] = "Buenos Aires"

	fmt.Println(paises)
	//Asi se hace llamar un valor especifico de un mapa, en el print solo aparecera su valor y no la llave el cual
	//es Chile.
	fmt.Println(paises["Chile"])

	//Otra forma de realizar un mapa sin usar la función make()
	campeonato := map[string]int{
		"Colo-Colo":   38,
		"U de Chile":  40,
		"Real Madrid": 70,
		"Naval":       76,
	}
	fmt.Println(campeonato)
	//Recorrer un mapa con un ciclo for
	for equipo, puntaje := range campeonato {
		//%s cuando la variable es string y %d cuando la variable es entero
		fmt.Printf("Equipo %s tiene puntaje de: %d \n", equipo, puntaje)
	}

	//Como eliminar un elemento de un mapa
	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	//Como realizar una consulta a un mapa para verificar si un dato existe dentro del mapa o no
	puntaje, existe := campeonato["Naval"]
	//%t es cuando se esta utilizando booleanos en el print para que aparezca en consola
	fmt.Printf("El puntaje capturado es de %d y el equipo existe = %t \n", puntaje, existe)
}
