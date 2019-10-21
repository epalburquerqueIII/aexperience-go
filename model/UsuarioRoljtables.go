package model

// UsuarioRolRecords estructura para comunicar con jtable
type UsuarioRolRecords struct {
	Result           string `json:"Result"`
	TotalRecordCount int
	Records          []TusuarioRol
}

// UsuarioRolRecord estructura para comunicar con jtable
type UsuarioRolRecord struct {
	Result string `json:"Result"`
	Record TusuarioRol
}
