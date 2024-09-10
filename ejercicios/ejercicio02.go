package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Mostrar tablas de multiplicación a partir del input introducido del usuario, se multiplica hasta el 10
func Multiplicar() {
	var numero int
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Ingrse un número: ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Hubo un error, introduzca un valor númerico")
				continue
			} else {
				for valor := 0; valor < 11; valor++ {
					fmt.Println(numero, "X", valor, "=", numero*valor)
				}
			}
		}
		break
	}
}
