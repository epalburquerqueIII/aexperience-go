package controller

import (
	"fmt"
	"net/http"
)
	
}
func UsuariosRegister(w http.ResponseWriter, r *http.Request) {
 error := tmpl.ExecuteTemplate(w, "registro", nil)
     if error != nil {
	fmt.Println("Error ", error.Error)
	}
	db := database.DbConn()
	usu := model.Tusuario{}
	if r.Method == "POST" {
		usu.Nombre = r.FormValue("Nombre")
		usu.Nif = r.FormValue("Nif")
		usu.Email = r.FormValue("Email")
		usu.FechaNacimiento = util.DateSql(r.FormValue("FechaNacimiento"))
		usu.Telefono = r.FormValue("Telefono")
		usu.Password = r.FormValue("Password")

		usu.IDUsuarioRol = 0
		usu.SesionesBonos = 0
		usu.Newsletter = 1
		usu.FechaBaja = util.DateSql("00-00-0000")

		insForm, err := db.Prepare("INSERT INTO usuarios(nombre, nif, email, fechaNacimiento, idusuariorol, telefono, password,sesionesbonos,newsletter,fechabaja) VALUES(?,?,?,?,?,?,?,?,?,?)")
		if err != nil {
			var verror model.Resulterror
			verror.Result = "ERROR"
			verror.Error = "Error registrando Usuario"
			a, _ := json.Marshal(verror)
			w.Write(a)
			panic(err.Error())
		}
		res, err1 := insForm.Exec(usu.Nombre, usu.Nif, usu.Email, usu.FechaNacimiento, usu.IDUsuarioRol, usu.Telefono, usu.Password, usu.SesionesBonos, usu.Newsletter, usu.FechaBaja)
		if err1 != nil {
			panic(err1.Error())
		}
		usu.ID, err1 = res.LastInsertId()
		log.Println("INSERT: nombre: " + usu.Nombre + " | nif: " + usu.Nif + "| password:" + usu.Password)

	}
	var vrecord model.UsuarioRecord
	vrecord.Result = "OK"
	vrecord.Record = usu
	a, _ := json.Marshal(vrecord)
	s := string(a)
	fmt.Println(s)

	w.Write(a)

	defer db.Close()
	//	http.Redirect(w, r, "/", 301)
}
