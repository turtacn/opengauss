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

	gsdb, err := gorm.Open(gsDialector, &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect opengauss db %v", err)
		return
	}

	if gsdb != nil {
		fmt.Println("DB struct is %v", gsdb)
	}

}
