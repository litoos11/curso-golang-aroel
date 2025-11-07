package models

import "goapirest/db"

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Users []User

const UserSchema string = `CREATE TABLE users (
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(100) NOT NULL,
	email VARCHAR(100),
	create_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)`

func NewUser(username, password, email string) *User {
	user := &User{
		Username: username,
		Password: password,
		Email:    email,
	}
	return user
}

func CreateUser(username, password, email string) *User {
	user := NewUser(username, password, email)
	user.Save()
	return user
}

func (u *User) insert() {
	query := "INSERT INTO users (username, password, email) VALUES (?, ?, ?)"
	result, _ := db.Exec(query, u.Username, u.Password, u.Email)
	u.Id, _ = result.LastInsertId()
}

func ListUsers() (Users, error) {
	query := "SELECT id, username, password, email FROM users"
	users := Users{}
	rows, err := db.Query(query)

	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
		users = append(users, user)
	}
	return users, err
}

func GetUser(id int) (*User, error) {
	user := NewUser("", "", "")

	query := "SELECT id, username, password, email FROM users WHERE id = ?"
	if rows, err := db.Query(query, id); err != nil {
		return nil, err
	} else {
		for rows.Next() {
			rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
		}
		return user, nil
	}
}

func (u *User) update() {
	query := "UPDATE users SET username=?, password=?, email=? WHERE id=?"
	db.Exec(query, u.Username, u.Password, u.Email, u.Id)
}

func (user *User) Save() {
	if user.Id == 0 {
		user.insert()
	} else {
		user.update()
	}
}

func (user *User) Delete() {
	query := "DELETE FROM users WHERE id=?"
	db.Exec(query, user.Id)
}
