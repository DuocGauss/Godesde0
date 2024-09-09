package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//La libreria bufio permite escanear y capturar datos introducidos desde la terminal en Go
//por ende sirven como inputs para la terminal

var numero1 int
var numero2 int
var leyenda string
var err error

func IngreseNumeros() {
	//Lee y captura lo que se escribe desde la terminal
	scanner := bufio.NewScanner(os.Stdin) //os.Stadin es el input por defecto que viene en la libreria os

	fmt.Println("Ingrse el numero 1: ")
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Hubo un error: " + err.Error())
		}
	}
	//Con panic() el programa se detiene por completo si detecta un error
	fmt.Println("Ingrse el numero 2: ")
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Hubo un error: " + err.Error())
		}
	}
	fmt.Println("Ingrse la leyenda: ")
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	fmt.Println(leyenda, ":", numero1*numero2)

}
