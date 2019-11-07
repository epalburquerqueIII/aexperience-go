package model

// AutorizadoRecords estructura para comunicar con jtable
type AutorizadoRecords struct {
	Result           string `json:"Result"`
	TotalRecordCount int
	Records          []Tautorizado
}

// AutorizadoRecord estructura para comunicar con jtable
type AutorizadoRecord struct {
	Result string `json:"Result"`
	Record Tautorizado
}
