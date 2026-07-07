package services

import (
	"crypto/md5"
	"fmt"
	"koda-b8-golang5/utils"
)

type Auth interface {
	Login(Email *string, Password *string)
	Register(FirstName *string, LastName *string, Email *string, Password *string)
	ConcatName() string
}

func (u User) ConcatName(FirstName string, LastName string) string {
	return FirstName + " " + LastName
}

func (u AuthForm) Register(FirstName *string, LastName *string, Email *string, Password *string, res Error) {
	encPass := md5.Sum([]byte(*Password))
	form := User{
		id:        *FirstName + *Email,
		FirstName: *FirstName,
		LastName:  *LastName,
		Email:     *Email,
		Password:  encPass,
	}

	var back string
	Accounts = append(Accounts, form)
	fmt.Printf("\n*%s\n", res.message)
	fmt.Print("\nTekan Enter untuk kembali ")
	fmt.Scanf("%s", &back)
	fmt.Println(back)
	utils.ClearTerm(0, "")
	defer AskingHomeAuth()
}

func (u AuthForm) Login(Email *string, Password *string) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%s\n\n", r)
			utils.ClearTerm(1, "*Email atau password salah")
			AskingLogin()
		}
	}()

	if len(Accounts) == 0 {
		defer AskingHomeAuth()
	}

	for x := range Accounts {
		if k := md5.Sum([]byte(*Password)); Accounts[x].Email == *Email && k == Accounts[x].Password {
			fmt.Println("\n*Login Berhasil")
			Dashboard(Accounts[x])
			return
		}
	}
	panic(1)
}
