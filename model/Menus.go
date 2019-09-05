package model

// Tusuario es la estructura para usuario
type Tmenus struct {
	Id         int64
	Parent_id  int
	Orden      int
	Titulo     string
	Icono      string
	Url        string
	HanledFunc string
}
