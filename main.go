package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"os/exec"
)

var accounts = []User{}

type User struct {
	id        string
	FirstName string
	LastName  string
	Email     string
	Password  [16]byte
}

type AuthForm struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type Auth interface {
	Login(Email *string, Password *string)
	Register(FirstName *string, LastName *string, Email *string, Password *string)
	ConcatName() string
}

func (u AuthForm) Register(FirstName *string, LastName *string, Email *string, Password *string) {

	encPass := md5.Sum([]byte(*Password))
	form := User{
		id:        *FirstName + *Password,
		FirstName: *FirstName,
		LastName:  *LastName,
		Email:     *Email,
		Password:  encPass,
	}
	var back string
	accounts = append(accounts, form)
	fmt.Print("\nTekan Enter untuk kembali ")
	fmt.Scanf("%s", &back)
	fmt.Println(back)

	defer main()
}

func clearTerm() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func (u User) ConcatName(FirstName string, LastName string) string {
	return FirstName + " " + LastName
}

func (u AuthForm) Login(Email *string, Password *string) {
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
		if k := md5.Sum([]byte(*Password)); accounts[x].Email == *Email && k == accounts[x].Password {
			fmt.Println("\n*Login Berhasil")
			successLogin(accounts[x])
		}
	}
	panic("*Email atau password salah")
}

func homeMenu() string {
	clearTerm()
	var point string
	fmt.Println("- SYSTEM_AUTH -")
	fmt.Printf("\n1. Login\n2. Register\n3. Forgot Password\n\n")
	fmt.Println("0. Exit")
	fmt.Printf("\nSilahkan masukan pilihan anda: ")
	fmt.Scanf("%s", &point)

	return point
}

func confirmRegister(form *AuthForm) int {
	var confirm string
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()

	for _, val := range accounts {
		if form.Email == val.Email {
			panic("Email sudah digunakan")
		}
	}

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
	clearTerm()
	form := AuthForm{
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
		form.Register(&form.FirstName, &form.LastName, &form.Email, &form.Password)
	} else {
		askingRegister()
	}

}

func successLogin(user User) {
	clearTerm()
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
		defer main()
	}

}

func askingLogin() {
	clearTerm()

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
		clearTerm()
		fmt.Println("Sampai jumpa!")
		fmt.Println("- SYSTEM_SHUTDOWN -")
		os.Exit(0)
	default:
		panic(value)
	}
}
