package entity

import (
	"ToDo/contracts"
	"ToDo/utils"
	"bufio"
	"encoding/json"
	"flag"
	"os"
)

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

const path = "user//userList.txt"

func (user User) RegisterUserMethod(storage contracts.Storage) {
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	userCount := user.load
	println("Please enter your name: ")
	scanner.Scan()
	user.Name = scanner.Text()
	println("Please enter your email: ")
	scanner.Scan()
	user.Email = scanner.Text()
	println("Please enter your password: ")
	scanner.Scan()
	rawPassword := scanner.Text()
	password, _ := utils.HashPasswordMD5(rawPassword)
	user.Password = password
	user.Id = userCount

	jDataBite, _ := json.Marshal(user)
	jDataString := string(jDataBite)
	storage.Save(jDataString)
}
