package model

// Records estructura para comunicar con jtable
type Records struct {
	Result           string `json:"Resultado"`
	TotalRecordCount int
	Records          []Tusuario
}
// Record estructura para comunicar con jtable
type Record struct {
	Result string `json:"Resultado"`
	Record Tusuario
}
// Option estructura para datos adiciones en jtable
type Option struct {
	Value       int
	DisplayText string
}
// Options estructura para datos adiciones en jtable
type Options struct {
	Result  string `json:"Resultado"`
	Options []Option
}
// Resulterror resultado de jtable
type Resulterror struct {
	Result string `json:"Resultado"`
	Error  string
}
