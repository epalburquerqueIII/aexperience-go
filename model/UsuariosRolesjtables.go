package model

// UsuariosRolesRecords estructura para comunicar con jtable
type UsuariosRolesRecords struct {
	Result           string `json:"Result"`
	TotalRecordCount int
	Records          []TusuariosRoles
}

// UsuariosRolesRecord estructura para comunicar con jtable
type UsuariosRolesRecord struct {
	Result string `json:"Result"`
	Record TusuariosRoles
}
