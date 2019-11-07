package model

// PagoPendienteRecords estructura para comunicar con jtable
type PagoPendienteRecords struct {
	Result           string `json:"Result"`
	TotalRecordCount int
	Records          []TpagoPendiente
}

// PagoPendienteRecord estructura para comunicar con jtable
type PagoPendienteRecord struct {
	Result string `json:"Result"`
	Record TpagoPendiente
}
