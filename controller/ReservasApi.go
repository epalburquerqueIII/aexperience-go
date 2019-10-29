package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"../model"
	"../model/database"
	"../util"
)

// Reservas Pantalla de tratamiento de Reservas
func Reservas(w http.ResponseWriter, r *http.Request) {
	menu := util.Menus(usertype)
	error := tmpl.ExecuteTemplate(w, "reservas", &menu)
	if error != nil {
		fmt.Println("Error ", error.Error)
	}
}

// ReservasList - json con los datos de clientes
func ReservasList(w http.ResponseWriter, r *http.Request) {

	var i int = 0
	jtsort := r.URL.Query().Get("jtSorting")
	if jtsort != "" {
		fmt.Println("jtSorting" + jtsort)
		jtsort = "ORDER BY " + jtsort
	}
	db := database.DbConn()
	selDB, err := db.Query("SELECT reservas.id, reservas.fecha, reservas.hora, usuarios.id, usuarios.nombre, espacios.id, espacios.descripcion, autorizados.id,autorizados.nombreAutorizado FROM reservas LEFT OUTER JOIN usuarios ON (usuarios.id = reservas.idUsuario) LEFT OUTER JOIN espacios ON (espacios.id = reservas.idEspacio) LEFT OUTER JOIN autorizados ON (autorizados.id = reservas.idAutorizado) " + jtsort)
	if err != nil {
		util.ErrorApi(err.Error(), w, "Error en Select ")
	}
	reser := model.Treserva{}
	res := []model.Treserva{}
	for selDB.Next() {

		err = selDB.Scan(&reser.Id, &reser.Fecha, &reser.Hora, &reser.IdUsuario, &reser.UsuarioNombre, &reser.IdEspacio, &reser.EspacioNombre, &reser.IdAutorizado, &reser.AutorizadoNombre)
		if err != nil {
			util.ErrorApi(err.Error(), w, "Error Cargando registros de Reservas")
		}
		res = append(res, reser)
		i++
	}

	var vrecords model.ReservasRecords
	vrecords.Result = "OK"
	vrecords.TotalRecordCount = i
	vrecords.Records = res
	// create json response from struct
	a, err := json.Marshal(vrecords)
	// Visualza
	s := string(a)
	fmt.Println(s)
	w.Write(a)
	defer db.Close()
}

// ReservasCreate Crear un Usuario
func ReservasCreate(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	reser := model.Treserva{}
	if r.Method == "POST" {
		reser.Fecha = util.DateSql(r.FormValue("Fecha"))
		//reser.Sesiones, _ = strconv.Atoi(r.FormValue("Sesiones"))
		reser.Hora, _ = strconv.Atoi(r.FormValue("Hora"))
		reser.IdUsuario, _ = strconv.Atoi(r.FormValue("IdUsuario"))
		reser.IdEspacio, _ = strconv.Atoi(r.FormValue("IdEspacio"))
		reser.IdAutorizado, _ = strconv.Atoi(r.FormValue("IdAutorizado"))

		insForm, err := db.Prepare("INSERT INTO reservas(fecha, hora, idUsuario, idEspacio, idAutorizado) VALUES(?,?,?,?,?)")

		if err != nil {
			util.ErrorApi(err.Error(), w, "Error Insertando Pago")
		}

		res, err1 := insForm.Exec(reser.Fecha, reser.Hora, reser.IdUsuario, reser.IdEspacio, reser.IdAutorizado)

		if err1 != nil {
			//panic(err1.Error())
			util.ErrorApi(err.Error(), w, "")
		}
		reser.Id, err1 = res.LastInsertId()
		log.Printf("INSERT: fecha: %s | hora:  %d\n ", reser.Fecha, reser.Hora)

	}
	var vrecord model.ReservasRecord
	vrecord.Result = "OK"
	vrecord.Record = reser
	a, _ := json.Marshal(vrecord)
	s := string(a)
	fmt.Println(s)

	w.Write(a)

	defer db.Close()
	//	http.Redirect(w, r, "/", 301)
}

// ReservasUpdate Actualiza las reservas
func ReservasUpdate(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	reser := model.Treserva{}
	if r.Method == "POST" {
		i, _ := strconv.Atoi(r.FormValue("Id"))
		reser.Id = int64(i)
		reser.Fecha = util.DateSql(r.FormValue("Fecha"))

		/* // convertir de español a fecha
		format := "02-01-2006"
		t, _ := time.Parse(format, reser.Fecha)
		// format date to string en ingles para sql
		format = "2006-01-02"
		reser.Fecha = t.Format(format) */

		//reser.Sesiones, _ = strconv.Atoi(r.FormValue("Sesiones"))
		reser.Hora, _ = strconv.Atoi(r.FormValue("Hora"))
		reser.IdUsuario, _ = strconv.Atoi(r.FormValue("IdUsuario"))
		reser.IdEspacio, _ = strconv.Atoi(r.FormValue("IdEspacio"))
		reser.IdAutorizado, _ = strconv.Atoi(r.FormValue("IdAutorizado"))

		insForm, err := db.Prepare("UPDATE reservas SET fecha=?, hora=?, idUsuario=?, idEspacio =?, idAutorizado=? WHERE id=?")
		if err != nil {
			util.ErrorApi(err.Error(), w, "Error Actualizando Base de Datos")
		}

		insForm.Exec(reser.Fecha, reser.Hora, reser.IdUsuario, reser.IdEspacio, reser.IdAutorizado, reser.Id)
		log.Printf("UPDATE: fecha: %s  | hora:  %d\n", reser.Fecha, reser.Hora)

	}
	defer db.Close()
	var vrecord model.ReservasRecord
	vrecord.Result = "OK"
	vrecord.Record = reser
	a, _ := json.Marshal(vrecord)
	w.Write(a)

	//	http.Redirect(w, r, "/", 301)
}

// ReservasDelete Borra reservas de la DB
func ReservasDelete(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	reser := r.FormValue("Id")
	delForm, err := db.Prepare("DELETE FROM reservas WHERE id=?")
	if err != nil {
		panic(err.Error())
		util.ErrorApi(err.Error(), w, "")
	}
	_, err1 := delForm.Exec(reser)
	if err1 != nil {
		util.ErrorApi(err.Error(), w, "Error Borrando reserva")
	}
	log.Println("Borrar")
	defer db.Close()
	var vrecord model.ReservasRecord
	vrecord.Result = "OK"
	a, _ := json.Marshal(vrecord)
	w.Write(a)

	// 	// 	// 	http.Redirect(w, r, "/", 301)
}

// Reservasgetoptions saca las id de las reservas para la tabla pagos
func Reservasgetoptions(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	selDB, err := db.Query("SELECT reservas.id, reservas.fecha from reservas Order by reservas.id")
	if err != nil {
		panic(err.Error())
	}
	elem := model.Option{}
	vtabla := []model.Option{}
	for selDB.Next() {
		err = selDB.Scan(&elem.Value, &elem.DisplayText)
		if err != nil {
			panic(err.Error())
		}
		vtabla = append(vtabla, elem)
	}

	var vtab model.Options
	vtab.Result = "OK"
	vtab.Options = vtabla
	// create json response from struct
	a, err := json.Marshal(vtab)
	// Visualza
	s := string(a)
	fmt.Println(s)
	w.Write(a)
	defer db.Close()
}

// ReservarBono Crear una reserva de Bono
func ReservarBono(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()

	reser := model.Treserva{}
	res := []model.Treserva{}
	if r.Method == "POST" {
		//Carga de datos de Reservas
		reser.Fecha = util.DateSql(r.FormValue("Fecha"))
		reser.Sesiones, _ = strconv.Atoi(r.FormValue("Sesiones"))
		reser.Hora = 0
		reser.IdUsuario = 18
		reser.IdEspacio = 1
		reser.IdAutorizado = 3

		insForm, err := db.Prepare("INSERT INTO reservas(fecha, sesiones, hora, idUsuario, idEspacio, idAutorizado) VALUES(CURDATE(),?,?,?,?,?)")

		if err != nil {
			util.ErrorApi(err.Error(), w, "Error Insertando Reserva de Bono")
		}

		res, err1 := insForm.Exec(reser.Fecha, reser.Sesiones, reser.Hora, reser.IdUsuario, reser.IdEspacio, reser.IdAutorizado)

		if err1 != nil {
			//panic(err1.Error())
			util.ErrorApi(err.Error(), w, "")
		}
		reser.Id, err1 = res.LastInsertId()
		log.Printf("INSERT: sesiones: %d | idUsuario:  %d\n ", reser.Sesiones, reser.IdUsuario)

	}

	pago := model.Tpago{}
	// Obtiene el ID de la reserva del bono en curso
	//Verificación de usuario
	selDB, err2 := db.Query("SELECT reservas.id FROM reservas WHERE idusuario=? ORDER BY reservas.fecha DESC LIMIT 1")
	if err2 != nil {
		util.ErrorApi(err2.Error(), w, "Error en Select ")
	}

	//Almacena ID en la estructura de Reservas
	err2 = selDB.Scan(&reser.Id)
	if err2 != nil {
		util.ErrorApi(err2.Error(), w, "Error Cargando registros de Reservas")
	}
	res = append(res, reser)

	//Pago en efectivo
	if pago.IdTipopago == 1 {
		//INSERCION DE DATOS
		if r.Method == "POST" {
			//Carga de datos para pagos
			i, _ := strconv.Atoi(r.FormValue("IdReserva"))
			reser.Id = int64(i)
			pago.FechaReserva = util.DateSql(r.FormValue("FechaReserva"))
			pago.FechaPago = util.DateSql(r.FormValue("FechaPago"))
			pago.IdTipopago = 1
			pago.Importe, _ = strconv.ParseFloat(r.FormValue("importe"), 64)

			insForm1, err1 := db.Prepare("INSERT INTO pagos(idReserva, fechaReserva, fechaPago, idTipopago, importe) VALUES(?,CURDATE(),CURDATE(),?,?)")

			if err1 != nil {
				util.ErrorApi(err1.Error(), w, "Error Insertando Pago")
			}
			res, err2 := insForm1.Exec(reser.Id, pago.FechaReserva, pago.FechaPago, pago.IdTipopago)
			if err2 != nil {
				panic(err1.Error())
			}
			pago.Id, err2 = res.LastInsertId()
			log.Printf("INSERT: idReserva: %d | idTipopago:  %d\n", pago.IdReserva, pago.IdTipopago)

			//Obtención de Sesiones del Usuario
			// TODO: PASAR ID DEL USUARIO A TRAVÉS DEL LOGIN
			selDB, err3 := db.Query("SELECT sesionesbonos FROM usuarios WHERE usuarios.id = ?")
			if err3 != nil {
				util.ErrorApi(err3.Error(), w, "Error en Select ")
			}

			usu := model.Tusuario{}
			reser := model.Treserva{}

			err3 = selDB.Scan(&usu.SesionesBonos)

			var sesionesUsuario int = 0

			sesionesUsuario = usu.SesionesBonos + reser.Sesiones

			//Update ususarios
			if r.Method == "POST" {
				i, _ := strconv.Atoi(r.FormValue("ID"))
				usu.ID = int64(i)
				sesionesUsuario, _ = strconv.Atoi(r.FormValue("SesionesBonos"))

				insForm, err4 := db.Prepare("UPDATE usuarios SET sesionesbonos=? WHERE usuarios.id=?")
				if err4 != nil {
					util.ErrorApi(err4.Error(), w, "Error Actualizando Base de Datos")
				}
				insForm.Exec(usu.SesionesBonos, usu.ID)
				log.Printf("UPDATE: ID: %d  | Sesiones:  %d\n", usu.ID, usu.SesionesBonos)
			}

		} else {
			//Pago pendiente

		}

		var vrecord model.ReservasRecord
		vrecord.Result = "OK"
		vrecord.Record = reser
		a, _ := json.Marshal(vrecord)
		s := string(a)
		fmt.Println(s)

		w.Write(a)

		defer db.Close()
		//	http.Redirect(w, r, "/", 301)
	}
}
