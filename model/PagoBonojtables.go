package model

// PagoBonoRecords estructura para comunicar con jtable
type PagoBonoRecords struct {
	Result           string `json:"Result"`
	TotalRecordCount int
	Records          []Tpagobono
}

// PagoBonoRecord estructura para comunicar con jtable
type PagoBonoRecord struct {
	Result string `json:"Result"`
	Record Tpagobono
}
