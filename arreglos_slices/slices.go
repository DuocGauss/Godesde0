package arreglos_slices

import "fmt"

//Ejemplo de Slice, a diferencia de un arreglo común no se define cuantos valores tendra el vector
var tablaSlice []int = []int{19, 22, 25}

var arreglo [10]int = [10]int{40, 11, 23, 33, 55, 88, 66, 21, 64, 76}

func MuestroSlices() {
	fmt.Println(tablaSlice)

	//Ejemplo de la funcionalidad de un slice debido a que se puede crear desde un arreglo común
	porcion := arreglo[3:]   //Slice creado desde un vector, desde 3 hasta el final del arreglo
	porcion2 := arreglo[:5]  //Slice creado desde un vector, desde 0 hasta la posición 5 del arreglo
	porcion3 := arreglo[6:7] //Slice creado desde un vector, desde la posición 6 hasta la 7 del arreglo

	//Cuando no se define ningun valor de los lados va tomar el primer valor si es por la parte izquierda,
	//Como en la variable "porcion2" y en la derecha contara hasta el ultimo valor
	//el cual es el 76 como en la variable "porcion"

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

//Lo potente de los slices es que también se puede definir inicialmente su largo como los arreglos pero se puede
//aumentar su capacidad del tamaño de ese slice si se desea y indicarle a Go que me guarde en memoria esos espacios
//reservados para esa capacidad que fue definida

func Capacidad() {
	//Con make definino la capacidad inicial del slice en el segundo argumento (5) para luego decirle a go que
	//me guarde en memoria y que tiene capacidad hasta 20 elementos, por ende el slice que por ahora tiene 5 elementos
	//tiene reservado memoria que puede llegar hasta los 20 elementos en su capacidad.
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	//Ejemplo de otro slice el cual no se define un tamaño inicial ni reservar capacidad
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
		fmt.Println(nums[i])
	}

	//Como realice un ciclo for y le hice un append del contador del ciclo hacia el slice, go fue aumentando la
	//capacidad del slice de forma automatica dependiendo del limite que tiene el ciclo el cual parara cuando
	//el contador no supere a 100 pero de todas forma go aumenta la capacidad del slice nums para tener espacio de sobra
	//si se requiere seguir añadiendo datos al slice.
	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))
}
