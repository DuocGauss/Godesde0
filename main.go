package main //El archivo principal de Go siempre debe importar main

import (
	"fmt"

	"github.com/DuocGauss/Godesde0/variables"
)

// Función principal de main, se debe llamar main y no debe tener nada entre sus parentesis
func main() {
	variables.MuestroEnteros()
	variables.RestoVariables()

	estado, texto := variables.RetornoTexto(100)
	fmt.Println(estado, texto)
}
