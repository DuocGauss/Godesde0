package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

// Ejemplo de método en Go
func RestoVariables() {
	Nombre = "Gustavo"
	Estado = true
	Sueldo = 800.000
	Fecha = time.Now()

	fmt.Println("Nombre =", Nombre)
	fmt.Println("Estado =", Estado)
	fmt.Println("Sueldo =", Sueldo)
	fmt.Println("Fecha =", Fecha)

}

// Ejemplo de funciones en Go
func RetornoTexto(numero int) (bool, string) {
	var texto string
	//Con Itoa() convierto el int a string
	texto = strconv.Itoa(numero)
	return true, texto
}
