package modelos

import (
	"time"
)

// Aqui se esta definiendo la entidad con las estructuras, debido a que con las estructuras Go maneja la programación
// orientadas a objetos y la creación de clases o entidades.
// Ejemplo de Estructura
type User struct {
	Id        int
	Name      string
	CreatedAt time.Time
	Status    bool
}

// Con este método se hace referencia a la entidad dentro de la estructura en su primer parametro, luego los datos
// recibido en su segundo parametros son agregados dentro de los atributos de la entidad User el cual esta dentro
// de su estructura, para que asi los registros sean almacenados dentro de la entidad
func (usuario *User) AddUser(id int, name string, createdat time.Time, status bool) {
	usuario.Id = id
	usuario.Name = name
	usuario.CreatedAt = createdat
	usuario.Status = status
}
