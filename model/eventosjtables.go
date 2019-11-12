package model

// EventosRecords estructura para comunicar con jtable
type EventosRecords struct {
	Result           string `json:"Result"`
	TotalRecordCount int
	Records          []Tevento
}

// EventosRecord estructura para comunicar con jtable
type EventosRecord struct {
	Result string `json:"Result"`
	Record Tevento
}
