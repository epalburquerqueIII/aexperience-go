package model

// Tmenus es la estructura para menus para la web y android
type Tmenus struct {
	Id         int64
	ParentId   int
	Orden      int
	Titulo     string
	Icono      string
	Url        string
	HanledFunc string
}
