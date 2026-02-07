package entity

import (
	"ToDo/storage"
	"ToDo/utils"
	"bufio"
	"fmt"
	"os"
)

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

func NewUser() (User, error) {
	user := User{}
	fileHandler := storage.FileHandler{}
	_, userCount, _ := fileHandler.Load("User")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter your name: ")
	scanner.Scan()
	user.Name = scanner.Text()
	fmt.Println("Please enter your email: ")
	scanner.Scan()
	user.Email = scanner.Text()
	fmt.Println("Please enter your password: ")
	scanner.Scan()
	rawPassword := scanner.Text()
	password, _ := utils.HashPasswordMD5(rawPassword)
	user.Password = password
	user.Id = userCount + 1
	return user, nil
}
