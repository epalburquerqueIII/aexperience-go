package model

// Tmenu es la estructura para el menú simple
type Tmenu struct {
	ID          int
	Icono       string
	ParentTitle string
	Options     []Tmenudesplegable
	Despliega   int
}

// Tmenudesplegable es la estructura para el menu desplegable
type Tmenudesplegable struct {
	NomEnlace string
	Enlace    string
	Orden     int
}
