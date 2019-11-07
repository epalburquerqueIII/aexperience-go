package model

// Tusuario es la estructura para usuario
type Tusuario struct {
	ID              int64
	Nombre          string
	Nif             string
	Email           string
	FechaNacimiento string
	IDUsuarioRol    int
	Telefono        string
	Password        string
	SesionesBonos   int
	Newsletter      int
	FechaBaja       string
}
