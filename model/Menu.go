package model

// TMenus es la estructura para usuario
type Tmenu struct {
	Id         int64
	ParentId   int
	Orden      int
	Titulo     string
	Icono      string
	Url        string
	HandleFunc string
}
