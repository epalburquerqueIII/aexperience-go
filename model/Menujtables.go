package model

// MenusRecords estructura para comunicar con jtable
type MenuRecords struct {
	Result           string `json:"Result"`
	TotalRecordCount int
	Records          []Tmenu
}

// MenusRecord estructura para comunicar con jtable
type MenuRecord struct {
	Result string `json:"Result"`
	Record Tmenu
}
