package model

// TiposPagoRecords estructura para comunicar con jtable
type TiposPagoRecords struct {
	Result           string `json:"Result"`
	TotalRecordCount int
	Records          []TtiposPago
}

// TiposPagoRecords estructura para comunicar con jtable
type TiposPagoRecord struct {
	Result string `json:"Result"`
	Record TtiposPago
}
