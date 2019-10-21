package model

// ReservasRecords estructura para comunicar con jtable
type ReservasRecords struct {
	Result           string `json:"Result"`
	TotalRecordCount int
	Records          []Treserva
}

// ReservasRecord estructura para comunicar con jtable
type ReservasRecord struct {
	Result string `json:"Result"`
	Record Treserva
}
