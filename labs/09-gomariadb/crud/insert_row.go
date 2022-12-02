package crud

import (
	"fmt"
)

func InsertRow() {
	db, err := Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("insert into cities(id, name, population) values(?,?,?)", 4, "Denpasar", 100000)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Insert Success")
}
