package model

// Tpago es la estructura para un pago
type Tpago struct {
	Id            int64
	IdReserva     int
	FechaReserva  string
	FechaPago     string
	IdTipopago    int
	TipoPago      string
	NumeroTarjeta string
	Importe       float64
	Referencia    string
}
