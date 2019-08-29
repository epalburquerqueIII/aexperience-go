package model

// UsuarioRecords estructura para comunicar con jtable
type AutorizadoRecords struct {
	Result           string `json:"Result"`
	TotalRecordCount int
	Records          []Tautorizado
}

// UsuarioRecord estructura para comunicar con jtable
type AutorizadoRecord struct {
	Result string `json:"Result"`
	Record Tautorizado
}
