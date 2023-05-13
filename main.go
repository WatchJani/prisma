package main

import (
	"fmt"
	"root/prisma"

	_ "github.com/lib/pq"
)

func main() {

	// user=janko dbname=company password=JankoKondic72621@ sslmode=disable
	DBConfig := prisma.DBConfig{
		Username: "janko",
		DBName:   "company",
		Password: "JankoKondic72621@",
		SSLmode:  "disable",
	}

	db, err := prisma.NewDB("postgres", DBConfig)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	user := prisma.Select(
		prisma.Field("name"),
		prisma.Table("user"),
	)

	fmt.Println(user)

}
