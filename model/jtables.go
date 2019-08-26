package model

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
