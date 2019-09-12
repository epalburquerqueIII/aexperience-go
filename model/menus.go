package model

// Tmenus es la estructura para menus
type Tmenus struct {
	Id         int64
	ParentId   int
	Orden      int
	Titulo     string
	Icono      string
	Url        string
	HanledFunc string
}
