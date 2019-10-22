package model

// PagoPendientesRecords estructura para comunicar con jtable
type PagoPendienteRecords struct {
	Result           string `json:"Result"`
	TotalRecordCount int
	Records          []TpagoPendiente
}

// PagosPendientesRecord estructura para comunicar con jtable
type PagoPendienteRecord struct {
	Result string `json:"Result"`
	Record TpagoPendiente
}
