package model

// Tpagos es la estructura para pagos
type Tpagos struct {
	Id            int64
	IdReserva     int
	FechaPago     string
	IdTipopago    int
	NumeroTarjeta string
}
