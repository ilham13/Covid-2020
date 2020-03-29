package models

import (
	"fmt"

	"github.com/ilham13/Covid-2020/config"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *User) GetUser() error {
	db, err := config.Init()
	if err != nil {
		return err
	}
	statement := fmt.Sprintf("SELECT name, age FROM users WHERE id=%d", u.ID)
	return db.QueryRow(statement).Scan(&u.Name, &u.Age)
}

func (u *User) UpdateUser() error {
	db, err := config.Init()
	if err != nil {
		return err
	}
	statement := fmt.Sprintf("UPDATE users SET name='%s', age=%d WHERE id=%d", u.Name, u.Age, u.ID)
	_, errStatement := db.Exec(statement)
	return errStatement
}

func (u *User) DeleteUser() error {
	db, err := config.Init()
	if err != nil {
		return err
	}
	statement := fmt.Sprintf("DELETE FROM users WHERE id=%d", u.ID)
	_, errStatement := db.Exec(statement)
	return errStatement
}

func (u *User) CreateUser() error {
	db, err := config.Init()
	statement := fmt.Sprintf("INSERT INTO users(name, age) VALUES('%s', %d)", u.Name, u.Age)
	_, errStatement := db.Exec(statement)

	if errStatement != nil {
		return errStatement
	}

	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&u.ID)

	if err != nil {
		return err
	}

	return nil
}

func (u *User) GetUsers(start, count int) ([]User, error) {

	db, err := config.Init()
	// defer db.Close()

	statement := fmt.Sprintf("SELECT id, name, age FROM users LIMIT %d OFFSET %d", count, start)
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
