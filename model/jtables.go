package model

// Option estructura para datos adiciones en jtable
type Option struct {
	Value       int
	DisplayText string
}

// Options estructura para datos adiciones en jtable
type Options struct {
	Result  string `json:"Result"`
	Options []Option
}

// Resulterror resultado de jtable
type Resulterror struct {
	Result string `json:"Result"`
	Error  string
}
