package model

// ConsumoBonosRecords estructura para comunicar con jtable
type ConsumoBonosRecords struct {
	Result           string `json:"Result"`
	TotalRecordCount int
	Records          []Tconsumo
}

// ConsumoBonosRecord estructura para comunicar con jtable
type ConsumoBonosRecord struct {
	Result string `json:"Result"`
	Record Tconsumo
}
