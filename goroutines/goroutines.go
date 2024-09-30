package goroutines

import (
	"fmt"
	"strings"
	"time"
)

// Ejemplo de rutina asincrona en Go y channels
func MiNombreLentoo(nombre string, canal1 chan bool) {
	//Asigno el parametro nombre en una variable el cual es un vector(letras) y uso el paquete de strings y su
	//función de Split() para separar el nombre en si por una cadena de texto vacio, lo que separara todas las letras
	//del nombre por un espacio vacio de la cadena de texto. El parametro canal1 es el channel
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
	//Le asigno valor a la variable de parametro la cual es un channel. El valor asignado es un true debido
	//a que es un booleano. Con esto mando un mensaje(Un true) por el channel que la Goroutine, su logica, finalizo.
	//La función main en main.go recibe ese true y cuando lo recibe significa que la goroutine termino de ejecutarse
	//Y eso sirve para que la función principal de main.go no finalice antes que la go routine, gracias al channel
	//La goroutine siempre finalizaran antes que la función principal en main.go y no al reves. Ese es su proposito del channel en las gorotuines, sincronizar.
	canal1 <- true
}
