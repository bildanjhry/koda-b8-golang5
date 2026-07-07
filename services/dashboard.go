package services

import (
	"fmt"
	"koda-b8-golang5/utils"
)

func Dashboard(user User) {
	utils.ClearTerm(0, "")
	var conLgOut string
	fmt.Println("******* WELCOME ON BOARD, CAPTAIN! ********")
	fmt.Println("\n==========================")
	fmt.Printf("Nama: %s", User.ConcatName(user, user.FirstName, user.LastName))
	fmt.Printf("\nEmail: %s", user.Email)
	fmt.Printf("\nPassword: %x", user.Password)
	fmt.Println("\n==========================")
	fmt.Println("\nTodo List :")
	fmt.Println("- Kosong -")
	fmt.Println("\n==========================")
	fmt.Printf("\n1. Logout")
	fmt.Print("\n\nPilih Aksi: ")
	fmt.Scanf("%s", &conLgOut)

	if conLgOut == "1" {
		utils.ClearTerm(0, "")
		AskingHomeAuth()
	}

}
