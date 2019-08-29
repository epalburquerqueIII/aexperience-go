package model

// UsuarioRecords estructura para comunicar con jtable
type BonoRecords struct {
	Result           string `json:"Result"`
	TotalRecordCount int
	Records          []Tbono
}

// BonoRecord estructura para comunicar con jtable
type BonoRecord struct {
	Result string `json:"Result"`
	Record Tbono
}
