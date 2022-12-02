package crud

import (
	"fmt"
)

func DeleteRow() {
	db, err := Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("delete from cities where id = ?", 4)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Delete Success")
}
