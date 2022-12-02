package main

import (
	"example.com/gomariadb/crud"
)

func main() {
	/*
		db, err := crud.Connect()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer db.Close()
	*/
	crud.InsertRow()
	crud.SelectAll()

	crud.UpdateRow()
	crud.SelectAll()

	crud.DeleteRow()
	crud.SelectAll()
}
