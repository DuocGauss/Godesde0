package main //El archivo principal de Go siempre debe importar main

import (
	/*"fmt"
	"runtime"

	"github.com/DuocGauss/Godesde0/ejercicios"
	"github.com/DuocGauss/Godesde0/iteraciones"
	"github.com/DuocGauss/Godesde0/teclado"
	"github.com/DuocGauss/Godesde0/variables"
	"github.com/DuocGauss/Godesde0/files"
	"github.com/DuocGauss/Godesde0/funciones"
	"github.com/DuocGauss/Godesde0/arreglos_slices"
	"github.com/DuocGauss/Godesde0/mapas"*/
	"github.com/DuocGauss/Godesde0/users"
)

//runtime permite hacer un escaneo de las especificaciones de tu computadora, es bueno esta libreria para
//especificar cual son las especificaciones de tu pc lo cual es fundamental si compilamos archivos go y se lo
//entregamos a otra persona, asi el cliente tendra noción de cuales eran las especificaciones de la computadora
//que compilo el archivo .go

// Función principal de main, se debe llamar main y no debe tener nada entre sus parentesis
func main() {
	/*variables.MuestroEnteros()
	variables.RestoVariables()

	estado, texto := variables.RetornoTexto(100)
	fmt.Println(estado, texto)

	//Ejemplo de condicionales en Go
	//con runtime.GOOS permite conocer cuál es tu sistema operativo de tu sistema
	// "=" para asignación, "==" para comparación
	// Se puede hacer una asignación a una variable dentro de la condicional
	// y se separa con ; para luego hacer la comparación de la variable
	// El || es un or
	if os := runtime.GOOS; os == "linux" || os == "os x" {
		fmt.Println("Esto no es Windows, es", os)
	} else {
		fmt.Println("Esto es", os)
	}

	//Alternativa de condicional con Switch
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Su PC es Linux, se utilizo condicionales con switch")
	case "os x":
		fmt.Println("Su PC es OS X, se utilizo condicionales con switch")
	default:
		fmt.Println("Su PC es", os, "Se utilizo condicionales con switch")
	}

	//Ejercicio 01
	numero, texto := ejercicios.DevuelveEntero("101")
	fmt.Println(numero, texto)

	//Ingresos de datos desde la terminal (input)
	teclado.IngreseNumeros()

	//Ciclos utilizando for
	iteraciones.Iteracion()

	//Ejercicio 02
	fmt.Println(ejercicios.TablaMultiplicar()) */

	//files
	/*files.GrabarTabla()
	files.ConcatenarTabla()
	files.LeerTablaConsola()*/

	//Funciones anónimas
	//funciones.Calculos()
	//Función clousure
	//funciones.LlamarClousure()
	//Función de recursión
	//funciones.Exponencia(2)

	//Arreglos y slices
	//arreglos_slices.MuestroArreglos()
	//arreglos_slices.MuestroSlices()
	//arreglos_slices.Capacidad()

	//Mapas
	//mapas.MostrarMapa()

	//Estructuras(Programación orientadas a objetos, entidades)
	users.AgregarUsuarioAhoraSi()

}
