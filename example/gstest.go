package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	gsCfg := postgres.Config{
		DSN:                  "host=localhost user=gaussdb password=Enmo@123 dbname=postgres port=15432 sslmode=disable TimeZone=Asia/Shanghai",
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
