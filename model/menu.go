package model

// Tmenu es la estructura para el menú
type Tmenu struct {
	ID        int64
	Icono     string
	ParentID  int
	NomEnlace string
	Enlace    string
	Despliega int
}
