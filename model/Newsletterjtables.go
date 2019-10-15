package model

// NewsletterRecords estructura para comunicar con jtable
type NewsletterRecords struct {
	Result           string `json:"Result"`
	TotalRecordCount int
	Records          []Tnewsletter
}

// NewsletterRecord estructura para comunicar con jtable
type NewsletterRecord struct {
	Result string `json:"Result"`
	Record Tnewsletter
}
