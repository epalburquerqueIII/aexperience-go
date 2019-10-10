package model

// HorariosRecords estructura para comunicar con jtable
type HorariosRecords struct {
	Result           string `json:"Result"`
	TotalRecordCount int
	Records          []Thorarios
}

// HorariosRecord estructura para comunicar con jtable
type HorariosRecord struct {
	Result string `json:"Result"`
	Record Thorarios
}
