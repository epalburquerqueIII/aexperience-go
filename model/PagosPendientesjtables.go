package model

// PagosPendientesRecords estructura para comunicar con jtable
type PagosPendientesRecords struct {
	Result           string `json:"Result"`
	TotalRecordCount int
	Records          []TpagosPendientes
}

// PagosPendientesRecord estructura para comunicar con jtable
type PagosPendientesRecord struct {
	Result string `json:"Result"`
	Record TpagosPendientes
}
