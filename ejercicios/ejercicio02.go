package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Mostrar tablas de multiplicación a partir del input introducido del usuario, se multiplica hasta el 10
func TablaMultiplicar() string {
	var numero int
	var err error
	var texto string
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
					texto += fmt.Sprintf("%d x %d = %d \n", numero, valor, numero*valor)
				}

			}
		}
		break
	}
	return texto
}

// con += se concatena los valores y se añaden, si solo uso texto = fmt.Sprintf se sobreescribiran los valores en
//la variable texto durante el ciclo, entonces si quiero añadir y no sobreescribir durante el ciclo for en la variable
//texto, debo concatenar con +=

//En otras palabras: Con =, los valores anteriores se sobrescriben en cada iteración, por lo que solo conservarás el último resultado.
//Con +=, los nuevos valores se concatenan (se agregan) al valor existente, manteniendo todos los resultados anteriores y sumando el nuevo contenido.
//En el contexto de mi código, usar += permite que se acumulen todas las líneas de la tabla de multiplicar en la variable texto

//fmt.Sprintf permite formatear y construir una cadena (string) a partir de valores de diferentes tipos, incluidos enteros, flotantes, etc.
//Por ende convierte los valores a string
