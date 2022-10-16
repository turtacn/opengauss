package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 对照school.sql中的定义
/*
-- 创建表user
CREATE TABLE "user"
(
    id INT PRIMARY KEY,
    username VARCHAR(50),
    first_name VARCHAR(20),
    last_name  VARCHAR(30)
);
*/

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

	gsCfg := postgres.Config{ //dbname=gaussdb # 默认
		DSN:                  "host=localhost user=gaussdb password=Enmo@123 dbname=school port=15432 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true,
	}

	gsDialector := postgres.New(gsCfg)

	db, err := gorm.Open(gsDialector, &gorm.Config{})
	if err != nil {
		fmt.Errorf(fmt.Sprintf("Failed to connect postgresDB db %v", err))
		return
	}

	db.AutoMigrate(&User{})

	for _, user := range users {
		db.Create(&user)
	}

	var users = make([]User, 0)

	result := db.Find(&users)
	if result.Error == nil {
		fmt.Println(fmt.Sprintf("[before] total count is %d", result.RowsAffected))
	} else {
		fmt.Errorf(fmt.Sprintf("[before] query failed by error %v", result.Error))
	}

	// Delete incoming data with filter
	db.Where(&User{Username: "bgates"}).Delete(&User{})

	users = make([]User, 0)
	result = db.Find(&users)
	if result.Error == nil {
		fmt.Println(fmt.Sprintf("[after] total count is %d", result.RowsAffected))
	} else {
		fmt.Errorf(fmt.Sprintf("[after] query failed by error %v", result.Error))
	}

	fmt.Println("done.")
}
