package users

import (
	"fmt"
	"time"

	"github.com/DuocGauss/Godesde0/modelos"
)

// Aqui la estructura la cual esta definida se almacena en la variable u, por ende oficialmente es un objeto
// Se puede utilizar el metodo AddUser porque esta ligado con la estructura User, por ende se puede llamar su metodo
// y añadirle datos que luego seran procesados y almacenados dentro de la estructura o entidad User.
func AgregarUsuarioAhoraSi() {
	u := new(modelos.User)
	u.AddUser(19, "Gustavo", time.Now(), true)
	fmt.Println(u)
}
