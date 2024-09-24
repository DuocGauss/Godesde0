package modelos

type Mujer struct {
	Hombre
}

//En este ejemplo se realiza la herencia en programación orientada objetos, por ende la entidad Mujer, ereda todos
//los atributos de la entidad Hombre y para hacer eso simplemente dentro de la estructura hay que llamar a la
//estructura Hombre para que herede sus atributos a Mujer.

func (m *Mujer) Respirar() {
	m.respirando = true
}

func (m *Mujer) Comer() {
	m.comiendo = true
}

func (m *Mujer) Pensar() {
	m.pensando = true
}

func (m *Mujer) Sexo() string {
	return "Mujer"
}

func (m *Mujer) EstaVivo() bool {
	m.vivo = false
	return false
}
