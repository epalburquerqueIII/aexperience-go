package model

// Usuarios_rolesRecords estructura para comunicar con jtable
type Usuarios_rolesRecords struct {
	Result           string `json:"Result"`
	TotalRecordCount int
	Records          []Tusuarios_roles
}

// Usuarios_rolesRecord estructura para comunicar con jtable
type Usuarios_rolesRecord struct {
	Result string `json:"Result"`
	Record Tusuarios_roles
}
