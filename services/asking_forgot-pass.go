package services

import (
	"crypto/md5"
	"fmt"
	"koda-b8-golang5/utils"
)

func AskingForgotPass() {
	var email string
	var newPassword string

	defer func() {
		if val := recover(); val != nil {
			utils.ClearTerm(1, "Email tidak ditemukan")
			AskingForgotPass()
		}
	}()

	fmt.Print("Masukan email anda: ")
	fmt.Scanf("%s", &email)
	for x := range Accounts {
		if email == Accounts[x].Email {
			fmt.Print("Silahkan buat password baru: ")
			fmt.Scanf("%s", &newPassword)
			encPass := md5.Sum([]byte(newPassword))
			Accounts[x].Password = encPass
			utils.ClearTerm(1, "Berhasil ubah password, silahkan login")
			AskingLogin()
		}
	}
	panic(1)
}
