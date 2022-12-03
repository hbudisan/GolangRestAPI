package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	Id     int
	Name   string
	Salary int
}

func dbConn() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:Welcome1@tcp(localhost:3306)/hrd")

	if err != nil {
		panic(err.Error())
	}

	return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("select * from employees order by id")
	if err != nil {
		panic(err.Error())
	}
	employee := Employee{}
	res := []Employee{}
	for selDB.Next() {
		err = selDB.Scan(&employee.Id, &employee.Name, &employee.Salary)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, employee)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

func main() {
	fmt.Println("Server Started on http://localhost:8080")
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8080", nil)
}
