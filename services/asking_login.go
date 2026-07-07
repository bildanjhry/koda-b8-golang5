package services

import (
	"fmt"
)

func AskingLogin() {

	form := AuthForm{
		Email:    "",
		Password: "",
	}

	fmt.Print("Masukan email: ")
	fmt.Scanf("%s", &form.Email)
	fmt.Print("Masukan Password: ")
	fmt.Scanf("%s", &form.Password)
	form.Login(&form.Email, &form.Password)
}
