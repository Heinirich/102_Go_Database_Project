package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func errorCheck(err error){
	if err != nil{
		log.Fatal(err)
	}
}

type City struct {
	Id int
	Name string
	Population int
}

func main()  {
	db,err := sql.Open("mysql", "root:Darkweb360@tcp(127.0.0.1:3306)/cities")
	defer db.Close()

	errorCheck(err)
	// var version string
	// err2 := db.QueryRow("SELECT version()").Scan(&version)
	// errorCheck(err2)

	// fmt.Println(version)
	res,err := db.Query("SELECT * From cities")
	defer res.Close()
	errorCheck(err)

	for res.Next() {
		var city City
		err := res.Scan(&city.Id,&city.Name,&city.Population)
		errorCheck(err)
		fmt.Println(city)
	}


	// sql := "INSERT INTO cities(name, population) VALUES ('Moscow', 12506000)"
    // addCity, err := db.Exec(sql)
	// errorCheck(err)

	// lastId,err := addCity.LastInsertId()
	// errorCheck(err)

	// fmt.Println(lastId)


}