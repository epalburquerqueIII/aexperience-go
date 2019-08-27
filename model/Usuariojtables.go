package model

// UsuarioRecords estructura para comunicar con jtable
type UsuarioRecords struct {
	Result           string `json:"Result"`
	TotalRecordCount int
	Records          []Tusuario
}

// UsuarioRecord estructura para comunicar con jtable
type UsuarioRecord struct {
	Result string `json:"Result"`
	Record Tusuario
}
