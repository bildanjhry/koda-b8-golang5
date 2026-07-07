package services

var Accounts = []User{}

type User struct {
	id        string
	FirstName string
	LastName  string
	Email     string
	Password  [16]byte
}

type Error struct {
	status  int
	code    string
	message string
}

type AuthForm struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}
