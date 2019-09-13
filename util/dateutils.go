package util

import (
	"time"
)


// DateSql Función para pasar a formato SQL date
func DateSql(fecha string) string {
	// convertir de español a fecha
	format := "02-01-2006"
	t, _ := time.Parse(format, fecha)
	// format date to string en ingles para sql
	format = "2006-01-02"
	return t.Format(format)
}
