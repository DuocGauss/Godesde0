package iteraciones

import (
	"fmt"
)

func Iteracion() {
	for i := 0; i <= 100; i += 5 {
		if i == 95 {
			continue
		}
		fmt.Println(i)
	}

}

//Con continue sirve para continuar la iteración dependiendo de la condición que se utilizo para el continue
//Por ende ignora esa condición la cual utiliza el continue y la iteración continua.
//Como también si se remplaza el continue por break la iteración acabara cuando la variable i sea igual a 95
//Con continue, ignora a i cuando alcanza 95 y continua el ciclo hasta que sea mayor a 100
//Para detener una interación se utiliza la función break
