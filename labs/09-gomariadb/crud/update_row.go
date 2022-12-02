package crud

import (
	"fmt"
)

func UpdateRow() {
	db, err := Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("update cities set name = ?, population = ? where id = ?", "Garut", 12345, 4)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Update Success")
}
