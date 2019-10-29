package model

// Tpago es la estructura para un pago
type Tpago struct {
	Id            int64
	IdReserva     int
	FechaReserva  string
	FechaPago     string
	IdTipopago    int
	TipoPago      string
	Importe       float64
	NumeroTarjeta string
}
