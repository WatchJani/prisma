package main

import (
	"root/prisma"

	_ "github.com/lib/pq"
)

func main() {

	// user=janko dbname=company password=JankoKondic72621@ sslmode=disable
	DBConfig := prisma.DBConfig{
		Username: "janko",          //UserName("janko")
		DBName:   "company",        //DBName("company")
		Password: "**************", //Password("****************")
		SSLmode:  "disable",        //SSLmode
	}

	db, err := prisma.NewDB("postgres", DBConfig)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	db.
		Field("name", "email").
		Table("user").
		Select()

}
