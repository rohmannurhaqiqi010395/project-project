package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	var users Users
	var arr_user []Users
	var response response

	db := connect()
	defer db.Close()

	rows, err := db.Query("select * from person")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&users.Id, &users.FirstName, &users.LastName); err != nil {
			log.Fatal(err.Error())
		} else {
			arr_user = append(arr_user, users)
		}
	}

	response.Status = 1
	response.Message = "success"
	response.Data = arr_user

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func insertUsersMultipart(w http.ResponseWriter, r *http.Request) {
	// var users Users
	// var users_arr []Users
	var response response

	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	FirstName := r.FormValue("first_name")
	LastName := r.FormValue("last_name")

	_, err = db.Exec("INSERT INTO person (first_name, last_name) value(?,?)", FirstName, LastName)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Success Add"
	// response.Data = users_arr
	// log.Print("Insert data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func updateUsersMultipart(w http.ResponseWriter, r *http.Request) {
	var response response

	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	id := r.FormValue("id")
	FirstName := r.FormValue("first_name")
	LastName := r.FormValue("last_name")

	_, err = db.Exec("update person set first_name = ?, last_name = ? where id = ?", FirstName, LastName, id)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Success update"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func deleteUsersMultipart(w http.ResponseWriter, r *http.Request) {
	var response response

	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	id := r.FormValue("id")

	_, err = db.Exec("delete from person where id = ?", id)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Success delete"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
