package model

// Treservas es la estructura para usuario
type Treserva struct {
	Id               int64
	Fecha            string
	FechaPago        string
	Hora             int
	IdUsuario        int
	UsuarioNombre    string
	IdEspacio        int
	EspacioNombre    string
	IdAutorizado     int
	AutorizadoNombre string
}
