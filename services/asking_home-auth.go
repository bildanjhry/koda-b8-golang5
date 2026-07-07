package services

import (
	"fmt"
	"koda-b8-golang5/utils"
	"os"
)

func AskingHomeAuth() {
	utils.ClearTerm(0, "")

	var point string
	fmt.Println("- SYSTEM_AUTH -")
	fmt.Printf("\n1. Login\n2. Register\n3. Forgot Password\n\n")
	fmt.Println("0. Exit")
	fmt.Printf("\nSilahkan masukan pilihan anda: ")
	fmt.Scanf("%s", &point)

	defer func() {
		if val := recover(); val != nil {
			fmt.Printf("Input dengan %s tidak tersedia\n\n", val)
			AskingHomeAuth()
		}
	}()

	switch point {
	case "1":
		utils.ClearTerm(0, "")
		AskingLogin()
	case "2":
		utils.ClearTerm(0, "")
		AskingRegister()
	case "3":
		utils.ClearTerm(0, "")
		AskingForgotPass()
	case "0":
		utils.ClearTerm(1, "Sampai Jumpa!")
		fmt.Println("- SYSTEM_SHUTDOWN -")
		os.Exit(0)
	default:
		panic(point)
	}
}
