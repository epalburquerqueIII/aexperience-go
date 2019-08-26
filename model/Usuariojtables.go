package model

// Records estructura para comunicar con jtable
type UsuarioRecords struct {
	Result           string `json:"Resultado"`
	TotalRecordCount int
	Records          []Tusuario
}

// Record estructura para comunicar con jtable
type UsuarioRecord struct {
	Result string `json:"Resultado"`
	Record Tusuario
}

// Option estructura para datos adiciones en jtable
type Option struct {
	Value       int
	DisplayText string
}
