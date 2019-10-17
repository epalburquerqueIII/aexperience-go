package model

// Tpagospendientes es la estructura para pagos
type TpagosPendientes struct {
	Id            int64
	IdReserva     int
	FechaPago     string
	IdTipopago    int
	NumeroTarjeta string
}
