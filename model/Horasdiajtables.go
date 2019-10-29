package model

// HorariosRecords estructura para comunicar con jtable
type HorariodiaRecords struct {
	Result           string `json:"Result"`
	TotalRecordCount int
	Records          []THorasdia
}

// HorariosRecord estructura para comunicar con jtable
type HorariosdiasRecord struct {
	Result string `json:"Result"`
	Record THorasdia
}
