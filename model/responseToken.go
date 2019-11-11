package model

// ResponseToken Datos devueltos en la llamada de los servicios
type ResponseToken struct {
	AuthTokenString    string `json:"data0"`
	RefreshTokenString string `json:"data1"`
	CsrfSecretToken    string `json:"data2"`
	UserName           string `json:"username"`
	UserID             int    `json:"userid"`
	UserRol            string `json:"userrol"`
}
