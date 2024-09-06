package variables

import "fmt" //Se debe importar este paquete para que Go tenga interacción con la terminal
//y se le pueda mandar texto

//Voy a crear una función global y definir 3 variables dentro de esa función
func MuestroEnteros() {
	var intcomun int
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println("Valor de intcomún =", intcomun)
	fmt.Println("Valor de intde32 =", intde32)
	fmt.Println("Valor de intde64 =", intde64)
}
