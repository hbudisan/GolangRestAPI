package crud

import (
	"fmt"
	"log"
)

type City struct {
	Id         int
	Name       string
	Population int
}

func SelectAll() {
	db, err := Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("select * from cities")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var city City
		err := rows.Scan(&city.Id, &city.Name, &city.Population)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v\n", city)
	}
}
