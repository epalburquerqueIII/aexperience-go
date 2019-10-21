package model

// Tpagopendientes es la estructura para pagos
type TpagoPendiente struct {
	Id            int64
	IdReserva     int
	FechaPago     string
	IdTipopago    int
	NumeroTarjeta string
}
