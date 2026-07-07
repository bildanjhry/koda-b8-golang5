package main

import (
	"fmt"
	"os"
)

var accounts = []User{}

type User struct {
	id        string
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func (u User) Register(FirstName *string, LastName *string, Email *string, Password *string) {
	form := User{
		id:        *FirstName + *Password,
		FirstName: *FirstName,
		LastName:  *LastName,
		Email:     *Email,
		Password:  *Password,
	}
	var back string
	accounts = append(accounts, form)
	fmt.Print("\nTekan Enter untuk kembali ")
	fmt.Scanf("%s", &back)
	fmt.Println(back)
	defer main()
}

func (u User) ConcatName(FirstName string, LastName string) string {
	return FirstName + " " + LastName
}

func (u User) Login(Email *string, Password *string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%s\n\n", r)
			askingLogin()
		}
	}()

	if len(accounts) < 1 {
		fmt.Printf("\n*Akun masih kosong*\n\n")
		defer main()
	}

	for x := range accounts {
		fmt.Println(accounts[x].Email)
		if accounts[x].Email == *Email && accounts[x].Password == *Password {
			fmt.Println("\n*Login Berhasil")
			successLogin(accounts[x])
		} else {
			panic("*Email atau password salah")
		}
	}
}

func homeMenu() string {
	var point string
	fmt.Println("***** SELAMAT DATANG ******")
	fmt.Printf("\n1. Login\n2. Register\n3. Forgot Password\n\n")
	fmt.Println("0. Exit")
	fmt.Printf("\nSilahkan masukan pilihan anda: ")
	fmt.Scanf("%s", &point)

	return point
}

func confirmRegister(form *User) int {
	var confirm string

	fmt.Printf("\n*Apakah sudah benar?")
	fmt.Printf("\nNama depan kamu:  %s", form.FirstName)
	fmt.Printf("\nNama belakang kamu:  %s", form.LastName)
	fmt.Printf("\nAlamat email kamu:  %s", form.Email)
	fmt.Print("\n(y/n): ")
	fmt.Scanf("%s", &confirm)

	if confirm == "y" {
		return 1
	}

	return 0
}

func askingRegister() {
	form := User{
		FirstName: "",
		LastName:  "",
		Email:     "",
		Password:  "",
	}

	fmt.Print("Masukan nama depan anda: ")
	fmt.Scanf("%s", &form.FirstName)
	fmt.Print("Masukan belakang anda: ")
	fmt.Scanf("%s", &form.LastName)
	fmt.Print("Masukan email anda: ")
	fmt.Scanf("%s", &form.Email)
	fmt.Print("Masukan password anda: ")
	fmt.Scanf("%s", &form.Password)

	if c := confirmRegister(&form); c == 1 {
		defer User.Register(form, &form.FirstName, &form.LastName, &form.Email, &form.Password)
	} else {
		askingRegister()
	}

}

func successLogin(user User) {
	var conLgOut string

	fmt.Println("\n==========================")
	fmt.Printf("\nNama: %s", User.ConcatName(user, user.FirstName, user.LastName))
	fmt.Printf("\nEmail: %s", user.Email)
	fmt.Println("\n\n==========================")
	fmt.Printf("\n1. Logout")
	fmt.Print("\n\nPilih Aksi: ")
	fmt.Scanf("%s", &conLgOut)
	if conLgOut == "1" {
		defer main()
	}

}

func askingLogin() {

	form := User{
		Email:    "",
		Password: "",
	}

	fmt.Print("Masukan email: ")
	fmt.Scanf("%s", &form.Email)
	fmt.Print("Masukan Password: ")
	fmt.Scanf("%s", &form.Password)

	defer User.Login(form, &form.Email, &form.Password)
}

func askingForgotPass() {
	fmt.Println("Hello")
}

func main() {

	defer func() {
		if val := recover(); val != nil {
			fmt.Printf("Input dengan %s tidak tersedia\n\n", val)
			main()
		}
	}()

	switch value := homeMenu(); value {
	case "1":
		askingLogin()
	case "2":
		askingRegister()
	case "3":
		askingForgotPass()
	case "0":
		fmt.Println("Terimakasih sudah berkunjung")
		os.Exit(0)
	default:
		panic(value)
	}
}
