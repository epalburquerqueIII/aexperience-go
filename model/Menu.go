package model

// TMenus es la estructura para usuario
type Tmenu struct {
	Id         int64
	ParentId   int
	MenuParent string
	Orden      int
	Titulo     string
	Icono      string
	Url        string
	HandleFunc string
}
