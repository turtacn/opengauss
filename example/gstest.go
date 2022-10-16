package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        uint
	Username  string
	FirstName string
	LastName  string
}

var users []User = []User{
	User{Username: "mzuckerberg", FirstName: "Mark", LastName: "Zuckerberg"},
	User{Username: "bgates", FirstName: "Bill", LastName: "Gates"},
	User{Username: "sjobs", FirstName: "Steve", LastName: "Jobs"},
	User{Username: "lellison", FirstName: "Larry", LastName: "Ellison"},
}

func main() {

	gsCfg := postgres.Config{ //dbname=postgres
		DSN:                  "host=localhost user=gaussdb password=Enmo@123 port=15432 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true,
	}

	gsDialector := postgres.New(gsCfg)

	postgresDB, err := gorm.Open(gsDialector, &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect postgresDB db %v", err)
		return
	}

	if postgresDB == nil {
		fmt.Println("DB instance is null")
		return
	}

	if testDB := postgresDB.Exec("CREATE DATABASE testdb"); testDB.Error != nil {
		fmt.Println("create databse test01db failed, error=%v", testDB.Error)
		return
	}

}
