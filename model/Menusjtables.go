package model

// MenusRecords estructura para comunicar con jtable
type MenusRecords struct {
	Result           string `json:"Result"`
	TotalRecordCount int
	Records          []Tmenus
}

// MenusRecord estructura para comunicar con jtable
type MenusRecord struct {
	Result string `json:"Result"`
	Record Tmenus
}
