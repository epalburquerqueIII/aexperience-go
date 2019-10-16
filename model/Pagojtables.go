package model

// PagoRecords estructura para comunicar con jtable
type PagoRecords struct {
	Result           string `json:"Result"`
	TotalRecordCount int
	Records          []Tpago
}

// PagoRecord estructura para comunicar con jtable
type PagoRecord struct {
	Result string `json:"Result"`
	Record Tpago
}
