package ejer_interfaces

import (
	"fmt"

	"github.com/DuocGauss/Godesde0/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando? = %t \n", hu.Sexo(), hu.EstaVivo())
}
