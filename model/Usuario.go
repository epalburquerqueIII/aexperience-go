package model

// Tusuario es la estructura para usuario
type Tusuario struct {
	ID               int64
	Nombre           string
	Nif              string
	Email            string
	Tipo             int
	Telefono         string
	SesionesBonos    int
	NewsletterNombre string //variable para mostrar en la tabla sí o no
	Newsletter       int
	FechaBaja        string
}
