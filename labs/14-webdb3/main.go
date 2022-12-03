package main

import (
	"database/sql"
	"log"
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

func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM employees WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	employee := Employee{}
	for selDB.Next() {
		err = selDB.Scan(&employee.Id, &employee.Name, &employee.Salary)
		if err != nil {
			panic(err.Error())
		}
	}
	tmpl.ExecuteTemplate(w, "Show", employee)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		salary := r.FormValue("salary")
		insForm, err := db.Prepare("INSERT INTO Employees(name, salary) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, salary)
		log.Println("INSERT: Name: " + name + " | Salary: " + salary)
	}
	defer db.Close()
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM Employees WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	for selDB.Next() {
		var id, salary int
		var name string
		err = selDB.Scan(&id, &name, &salary)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.Salary = salary
	}
	tmpl.ExecuteTemplate(w, "Edit", emp)
	defer db.Close()
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		salary := r.FormValue("salary")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE Employees SET name=?, salary=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, salary, id)
		log.Println("UPDATE: Name: " + name + " | Salary: " + salary)
	}
	defer db.Close()
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	_, err := db.Query("DELETE FROM Employees WHERE id=?", nId)
	log.Println("DELETE")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)

	http.HandleFunc("/new", New)
	http.HandleFunc("/insert", Insert)

	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/update", Update)

	http.HandleFunc("/remove", Delete)

	http.ListenAndServe(":8080", nil)
}
