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

//MenuRoles - Pantalla de tratamiento
func MenuRoles(w http.ResponseWriter, r *http.Request) {
	error := tmpl.ExecuteTemplate(w, "menuroles", nil)
	if error != nil {
		fmt.Println("Error: ", error.Error)
	}
}

//MenuRolesList - json con los datos de clientes
func MenuRolesList(w http.ResponseWriter, r *http.Request) {

	var i int
	jtsort := r.URL.Query().Get("jtSorting")
	if jtsort != "" {
		fmt.Println("jtSorting" + jtsort)
		jtsort = "ORDER BY " + jtsort
	}
	db := database.DbConn()
	selDB, err := db.Query("SELECT id, idMenu, idUsuarioRoles FROM menuusuariosroles " + jtsort)
	if err != nil {
		util.ErrorApi(err.Error(), w, "Error en Select ")
	}
	menuroles := model.Tmenuroles{}
	res := []model.Tmenuroles{}
	for selDB.Next() {

		err = selDB.Scan(&menuroles.Id, &menuroles.IdMenu, &menuroles.IdUsuarioRoles)
		if err != nil {
			util.ErrorApi(err.Error(), w, "Error Cargando datos de menu roles")
		}
		res = append(res, menuroles)
		i++
	}

	var vrecords model.MenuRolesRecords
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

//MenuRolesCreate - Crear campos de MenuRoles
func MenuRolesCreate(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	menurol := model.Tmenuroles{}
	if r.Method == "POST" {
		menurol.IdMenu, _ = strconv.Atoi(r.FormValue("IdMenu"))
		menurol.IdUsuarioRoles, _ = strconv.Atoi(r.FormValue("IdUsuarioRoles"))

		insForm, err := db.Prepare("INSERT INTO menuusuariosroles(idMenu, idUsuarioRoles) VALUES(?,?)")

		if err != nil {
			util.ErrorApi(err.Error(), w, "Error Insertando datos de menu roles")
		}
		res, err1 := insForm.Exec(menurol.IdMenu, menurol.IdUsuarioRoles)
		if err1 != nil {
			panic(err1.Error())
		}
		menurol.Id, err1 = res.LastInsertId()
		log.Printf("INSERT: IdMenu: %d | idUsuarioRoles: %d\n", menurol.IdMenu, menurol.IdUsuarioRoles)

	}
	var vrecord model.MenuRolesRecord
	vrecord.Result = "OK"
	vrecord.Record = menurol
	a, _ := json.Marshal(vrecord)
	s := string(a)
	fmt.Println(s)

	w.Write(a)

	defer db.Close()
	//	http.Redirect(w, r, "/", 301)
}

//MenuRolesUpdate - Actualiza el campo de MenuRoles
func MenuRolesUpdate(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	menurol := model.Tmenuroles{}
	if r.Method == "POST" {
		i, _ := strconv.Atoi(r.FormValue("Id"))
		menurol.Id = int64(i)
		menurol.IdMenu, _ = strconv.Atoi(r.FormValue("IdMenu"))
		menurol.IdUsuarioRoles, _ = strconv.Atoi(r.FormValue("IdUsuarioRoles"))

		insForm, err := db.Prepare("UPDATE menuusuariosroles SET idMenu=?, idUsuarioRoles=? WHERE menuusuariosroles.Id=?")
		if err != nil {
			util.ErrorApi(err.Error(), w, "Error Actualizando Base de Datos")
		}

		insForm.Exec(menurol.IdMenu, menurol.IdUsuarioRoles, menurol.Id)
		log.Printf("UPDATE: IdMenu: %d | idUsuarioRoles: %d\n", menurol.IdMenu, menurol.IdUsuarioRoles)
	}
	defer db.Close()
	var vrecord model.MenuRolesRecord
	vrecord.Result = "OK"
	vrecord.Record = menurol
	a, _ := json.Marshal(vrecord)
	w.Write(a)

	//	http.Redirect(w, r, "/", 301)
}

//MenuRolesDelete - Borra de la DB
func MenuRolesDelete(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	menurol := r.FormValue("Id")
	delForm, err := db.Prepare("DELETE FROM menuusuariosroles WHERE id=?")
	log.Println(menurol)
	if err != nil {
		panic(err.Error())
	}
	_, err1 := delForm.Exec(menurol)
	if err1 != nil {
		util.ErrorApi(err.Error(), w, "Error borrando datos de menuroles")
	}
	log.Println("DELETE")
	defer db.Close()
	var vrecord model.MenuRolesRecord
	vrecord.Result = "OK"
	a, _ := json.Marshal(vrecord)
	w.Write(a)

	http.Redirect(w, r, "/", 301)
}

//MenuRolesGetoptions - Nombre menu roles
func MenuRolesGetOptions(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()
	selDB, err := db.Query("SELECT menuusuariosroles.id, idMenu, idUsuarioRoles from menuusuariosroles Order by menuusuariosroles.id")
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
