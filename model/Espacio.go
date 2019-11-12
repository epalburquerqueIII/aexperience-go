package model

// Tespacio es la estructura para espacios
type Tespacio struct {
	ID                   int64
	Descripcion          string
	Estado               int
	Modo                 int
	Precio               int
	IDTipoevento         int
	Aforo                int
	Fecha                string
	NumeroReservaslimite int
}
