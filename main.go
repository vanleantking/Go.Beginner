/** For insert MySQL
*/
package main

import (
	"database/sql"
	"fmt"

	// "time"

	"./models"
	_ "github.com/go-sql-driver/mysql"
)

const (
	USERNAME = "root"
	PASSWORD = "mysql"
	HOST     = "127.0.0.1:3306"
	DBNAME   = "golang"
)

func main() {
	DB, err := sql.Open("mysql", USERNAME+":"+PASSWORD+"@"+"/"+DBNAME+"?parseTime=true")
	if err != nil {
		return
	}
	defer DB.Close()
	u, er := models.FindById(DB, "1")
	if er != nil {
		panic(er.Error())
	}
	fmt.Println("zzzzzz", u.Username)
	u.Username = "vanle"
	row, er := models.Update(DB, u)
	if er != nil {
		panic(er.Error())
	}
	if row == 0 {
		fmt.Println("have no record affected")
	} else {
		fmt.Println("update success")
		fmt.Println("zzzzzz", u.Username)
	}

	var user2 = models.User{}
	user2.Username = "oooooo"
	user2.Password = "123456"
	id, er := models.Insert(DB, user2)
	if er != nil {
		panic(er.Error())
	}
	fmt.Println(id)
	users, er := models.Find(DB, "")
	if er != nil {
		panic(er.Error())
	}
	fmt.Println(users[0].Username)

}
