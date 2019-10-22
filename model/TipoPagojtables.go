package model

// TiposPagoRecords estructura para comunicar con jtable
type TipoPagoRecords struct {
	Result           string `json:"Result"`
	TotalRecordCount int
	Records          []TtipoPago
}

// TiposPagoRecords estructura para comunicar con jtable
type TipoPagoRecord struct {
	Result string `json:"Result"`
	Record TtipoPago
}
