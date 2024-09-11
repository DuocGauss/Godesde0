package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/DuocGauss/Godesde0/ejercicios"
)

var ruta string = "./files/txt/tabla.txt"

// Con este método se crea y guarda en un archivo txt la tabla de multiplicar ingresada por consola por el usuario
func GrabarTabla() {
	var texto string = ejercicios.TablaMultiplicar()
	archivo, err := os.Create(ruta)
	if err != nil {
		fmt.Println("Hubo un error: " + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

//Con Fprintln agrego el contenido de la variable texto al archivo creado en la variable archivo el cual es tabla.txt

func ConcatenarTabla() {
	var texto string = ejercicios.TablaMultiplicar()
	if Append(texto) == false {
		fmt.Println("Hubo un error al concatenar las tablas")
	}

}

func Append(texto string) bool {
	arch, err := os.OpenFile(ruta, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Hubo un error en el Append " + err.Error())
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Hubo un error al escribir el texto " + err.Error())
		return false
	}
	arch.Close()
	return true

}

func LeerTablaConsola() {
	arch, err := os.Open(ruta)
	if err != nil {
		fmt.Println("Hubo un error al leer el archivo " + err.Error())
	}
	leer := bufio.NewScanner(arch)
	for leer.Scan() {
		texto := leer.Text()
		fmt.Println("> " + texto)
	}
	arch.Close()

}
