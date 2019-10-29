package model

// MenusRecords estructura para comunicar con jtable
type TipoNoticiaRecords struct {
	Result           string `json:"Result"`
	TotalRecordCount int
	Records          []TtipoNoticia
}

// MenusRecord estructura para comunicar con jtable
type TipoNoticiaRecord struct {
	Result string `json:"Result"`
	Record TtipoNoticia
}
