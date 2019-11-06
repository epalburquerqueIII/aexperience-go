package model

// TpagoPendiente es la estructura para pagos
type TpagoPendiente struct {
	Id             int64
	IdReserva      int
	ReservaNombre  string
	FechaPago      string
	IdTipopago     int
	TipopagoNombre string
	NumeroTarjeta  string
	Importe        float64
	Referencia     string
}
