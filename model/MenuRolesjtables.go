package model

//MenuRolesRecords estructura para comunicar con jtable
type MenuRolesRecords struct {
	Result           string `json:"Result"`
	TotalRecordCount int
	Records          []Tmenuroles
}

//MenuRolesRecord estructura para comunicar con jtable
type MenuRolesRecord struct {
	Result string `json:"Result"`
	Record Tmenuroles
}
