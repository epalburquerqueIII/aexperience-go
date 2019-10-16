package model

// Treservas es la estructura para usuario
type Treserva struct {
	Id        int64
	Fecha     string
	FechaPago string
	Hora      int
	IdUsuario int

	IdEspacio int

	IdAutorizado int
}
