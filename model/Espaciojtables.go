package model

//EspacioRecords estructura para comunicar con jtable
type EspacioRecords struct {
	Result           string `json:"Result"`
	TotalRecordCount int
	Records          []Tespacio
}

// EspacioRecord estructura para comunicar con jtable
type EspacioRecord struct {
	Result string `json:"Result"`
	Record Tespacio
}
