package model

// Tespacio es la estructura para espacios
type Tespacios struct {
	ID                   int64
	Descripcion          string
	Estado               string
	Modo                 string
	Precio               int
	IdTipoevento         int
	Aforo                int
	Fecha                string
	NumeroReservaslimite int
}
