package model

// TiposeventoRecords estructura para comunicar con jtable
type TiposeventoRecords struct {
	Result           string `json:"Result"`
	TotalRecordCount int
	Records          []Ttiposevento
}

// TiposeventoRecord estructura para comunicar con jtable
type TiposeventoRecord struct {
	Result string `json:"Result"`
	Record Ttiposevento
}
