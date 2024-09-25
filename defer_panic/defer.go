package defer_panic

import (
	"fmt"
	"log"
)

// Ejemplo de Defer
func EjemploDefer() {
	defer fmt.Println("Esto se ejecutara al ultimo")
	fmt.Println("Primer print")
	fmt.Println("Segundo print")
}

//Lo que hace el defer es que siempre ejecutara al final dentro de la función la linea de codigo el cual contiene
//el defer. Por ende con defer su proposito es que siempre se ejecute algo justo al finalizar una función o metodo.

//El comportamiento de defer es útil para tareas como cerrar archivos, liberar recursos o imprimir mensajes finales,
//asegurando siempre su ejecución al cierre de la función.

//El defer por lo general acepta una instrucción o linea de codigo, pero si se quiere utilizar con mucha lógica o
//muchas lineas de codigos es necesario agregarle al defer a una función anónima.

// Ejemplo de panic y recover
func EjemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrio un error que generó Panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Se encontro el valor 1, por ende eso causa error")
	}
}

//El panic sirve para abortar el sistema con un mensaje, si ocurre alguna condición anteriormente.
//Si cumple la condición se ejecutara el panic y el sistema se abortara dejando un texto de error por consola.

//La función de recover es recupersarse de un panic al capturar el error del panic dentro de una variable.
//Por ende sirve para que el programa se recupere y no se cierre abruptamente.
//Recover tiene que funcionar con defer para que el error del panic sea capturado
