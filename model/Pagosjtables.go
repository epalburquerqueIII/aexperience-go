package model

// PagosRecords estructura para comunicar con jtable
type PagosRecords struct {
	Result           string `json:"Result"`
	TotalRecordCount int
	Records          []Tpagos
}

// PagosRecord estructura para comunicar con jtable
type PagosRecord struct {
	Result string `json:"Result"`
	Record Tpagos
}
