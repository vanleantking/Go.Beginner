package models

import (
	"database/sql"
	// "time"
)

type User struct {
	Id         int
	Username   string
	Password   string
	Dob        sql.NullString
	Create_int sql.NullString
}

func FindById(db *sql.DB, id string) (User, error) {
	find := "SELECT * FROM user u where u.id = ?"
	stmt, er := db.Prepare(find)
	if er != nil {
		return User{}, er
	}
	var user User

	stmt.QueryRow(id).Scan(&user.Id, &user.Username, &user.Password, &user.Dob, &user.Create_int)
	return user, nil
}

func Find(db *sql.DB, where string) ([]User, error) {
	find := ""
	if where == "" {
		find = "SELECT * FROM user u"
	} else {
		find = "SELECT * FROM user WHERE " + where
	}
	var users []User

	stmt, er := db.Prepare(find)
	if er != nil {
		return users, er
	}
	user, er := stmt.Query()
	if er != nil {
		return users, er
	}
	for user.Next() {
		var u User
		er := user.Scan(&u.Id, &u.Username, &u.Password, &u.Dob, &u.Create_int)
		if er != nil {
			return users, er
		}
		users = append(users, u)
	}
	return users, nil
}

func Update(db *sql.DB, user User) (int64, error) {
	update := "UPDATE user SET username = ? where id = ?"
	re, er := db.Exec(update, user.Username, user.Id)
	if er != nil {
		return 0, er
	}
	row, er := re.RowsAffected()
	if er != nil {
		return 0, er
	}
	return row, nil
}

func Insert(db *sql.DB, user User) (int64, error) {
	insert := "INSERT INTO user(username, password) VALUES (?, ?)"
	re, er := db.Exec(insert, user.Username, user.Password)

	if er != nil {
		return 0, er
	}
	id, er := re.LastInsertId()
	if er != nil {
		return 0, er
	}
	return id, nil
}
