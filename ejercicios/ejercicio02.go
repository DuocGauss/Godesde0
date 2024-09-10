package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Multiplicar() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrse un número: ")
	if scanner.Scan() {
		numero, err := strconv.Atoi(scanner.Text())
		for valor := 0; valor < 11; valor++ {
			if err != nil {
				fmt.Println("Hubo un error porfavor ingrese un número")
				continue
			} else {
				fmt.Println(numero, "X", valor, "=", numero*valor)
			}
		}

	}
}
