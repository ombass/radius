package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Hello Word")

	var (
		login    string
		password string
		key      string
		baza     [2]string
		status   bool
	)
	baza[0] = "admin:admin"
	baza[1] = "root:root"

	fmt.Print("Login:")
	fmt.Scanln(&login)
	fmt.Print("password:")
	fmt.Scan(&password)
	key = (login + ":" + password)
	for i := 0; i < len(baza); i++ {
		if key == baza[i] {
			status = true
		}
	}
	if status == true {
		println("вошёл")
	} else {
		println("неверный логин или пароль")
	}
}
