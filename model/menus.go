package model

// Tmenu es la estructura para menus
type Tmenu struct {
	ID         int64
	ParentID   int
	Orden      string
	Titulo     string
	Icono      string
	Url        string
	HandleFunc string
}
