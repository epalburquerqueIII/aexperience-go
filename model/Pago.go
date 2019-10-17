package model

// Tpago es la estructura para pagos
type Tpago struct {
	Id            int64
	IdReserva     int
	FechaPago     string
	IdTipopago    int
	NumeroTarjeta string
}
