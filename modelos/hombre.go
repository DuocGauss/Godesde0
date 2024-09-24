package modelos

type Hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	vivo       bool
}

func (h *Hombre) Respirar() {
	h.respirando = true
}

func (h *Hombre) Comer() {
	h.comiendo = true
}

func (h *Hombre) Pensar() {
	h.pensando = true
}

func (h *Hombre) Sexo() string {
	return "Hombre"
}

func (h *Hombre) EstaVivo() bool {
	h.vivo = true
	return true
}

//Como los metodos y funcion de hombre tienen el mismo nombre que las funciones definidas en la interfaz humano
//Go automaticamente hace la conexión de ambos, de los metodos creados en hombre.go y los definidos en humano.go
