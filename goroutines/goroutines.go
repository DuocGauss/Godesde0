package goroutines

import (
	"fmt"
	"strings"
	"time"
)

// Ejemplo de rutina asincrona en Go
func MiNombreLentoo(nombre string) {
	//Asigno el parametro nombre en una variable el cual es un vector(letras) y uso el paquete de strings y su
	//función de Split() para separar el nombre en si por una cadena de texto vacio, lo que separara todas las letras
	//del nombre por un espacio vacio de la cadena de texto.
	letras := strings.Split(nombre, "")
	//Como asigne strings.Split a la variable letras se convierte en un slices por ende como es un vector puedo
	//utilizar el ciclo for range, pero como for range pide que se define 2 variables, en este caso se puede optar
	//por ignorar la primera variable con _ debido a que solo se quiere almacenar las letras separadas
	//en la variable letra durante el ciclo
	for _, letra := range letras {
		//Con time.Sleep() lo que hace es provocar una pausa en el sistema cuando se ejecuta la función
		//Esto provocara que se demore 1 segundo que se muestre letra por letra el nombre introducido en el parametro
		//de la función
		time.Sleep(1000 * time.Millisecond)
		//Para decirle que se demore 1 segundo en la ejecución, hay que multiplicar 1000 milisegundos con time.Millisecond
		//Si quiero que se demore más de 1 segundo, ejemplo 5 segundos, debo multiplicar 5000 milisegundos con time.Millisecond
		fmt.Println(letra)
	}
}
