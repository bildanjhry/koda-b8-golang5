package services

import (
	"fmt"
	"koda-b8-golang5/utils"
)

func ConfirmRegister(form *AuthForm) {
	var confirm string

	resRegis := Error{
		status:  1,
		code:    "",
		message: "",
	}

	defer func() {
		if a := recover(); a != nil {
			resRegis.message = "*Email sudah digunakan"
			resRegis.code = "FAILED_CREATE_ACCOUNT"
			utils.ClearTerm(1, "*Email sudah digunakan")
			AskingRegister()
		}
	}()

	for _, val := range Accounts {
		if form.Email == val.Email {
			panic(1)
		}
	}

	fmt.Printf("\n*Apakah sudah benar?")
	fmt.Printf("\nNama depan anda:  %s", form.FirstName)
	fmt.Printf("\nNama belakang anda:  %s", form.LastName)
	fmt.Printf("\nAlamat email anda:  %s", form.Email)
	fmt.Print("\n(y/n): ")
	fmt.Scanf("%s", &confirm)

	if confirm == "y" {
		resRegis.status = 0
		resRegis.code = "SUCCESS_CREATE_ACCOUNT"
		resRegis.message = "Berhasil buat akun"
	} else {
		utils.ClearTerm(0, "")
		AskingRegister()
	}

	form.Register(&form.FirstName, &form.LastName, &form.Email, &form.Password, resRegis)
}

func AskingRegister() {
	form := AuthForm{
		FirstName: "",
		LastName:  "",
		Email:     "",
		Password:  "",
	}

	fmt.Print("Nama depan anda: ")
	fmt.Scanf("%s", &form.FirstName)
	fmt.Print("Nama belakang anda: ")
	fmt.Scanf("%s", &form.LastName)
	fmt.Print("Email anda: ")
	fmt.Scanf("%s", &form.Email)
	fmt.Print("Password anda: ")
	fmt.Scanf("%s", &form.Password)

	ConfirmRegister(&form)

}
