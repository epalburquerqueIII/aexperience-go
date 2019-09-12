package model

// Tmenu es la estructura para el menú
type Tmenu struct {
	Icono        string
	ParentTitle  string
	Options      []Tmenudesplegable
	uniqueOption Tmenudesplegable
	Despliega    int
}

// Tmenudesplegable es la estructura para el menu desplegable
type Tmenudesplegable struct {
	NomEnlace string
	Enlace    string
}
