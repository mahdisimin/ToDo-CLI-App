package user

import (
	"ToDo/utils"
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

var userList []User

const path = "user//userList.txt"

func (user User) RegisterUserMethod() {
	serializationMode := flag.String("serializationMode", "json", "how to serialize the user data")
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	userCount, _ := GetUserList()
	println("Please enter your name: ")
	scanner.Scan()
	user.Name = scanner.Text()
	println("Please enter your email: ")
	scanner.Scan()
	user.Email = scanner.Text()
	println("Please enter your password: ")
	scanner.Scan()
	rawPassword := scanner.Text()
	//passwordBase64 := []byte(rawPassword)
	password, _ := utils.HashPasswordMD5(rawPassword)
	user.Password = password
	user.Id = userCount
	var data []byte
	var err error

	if *serializationMode == "simple" {
		data = []byte(fmt.Sprintf("Id : %v , Name : %v , Email : %v\n", user.Id, user.Name, user.Email))
	} else if *serializationMode == "json" {
		data, err = json.Marshal(user)
		if err != nil {
			log.Fatal("error on Json serializing user : ", err)

			return
		}

		data = append(data, '\n')

	} else {
		fmt.Println("Unsupported serialization mode")
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("error on opening file : ", err)

		return
	}

	_, wErr := file.Write(data)
	if wErr != nil {
		log.Fatal("error on writing to file : ", wErr)
	}

}

func (user User) LoginUser() (error, *User) {
	var authenticatedUser *User
	_, userList := GetUserList()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter your name: ")
	scanner.Scan()
	name := scanner.Text()
	println("Please enter your password: ")
	scanner.Scan()
	password := scanner.Text()
	for _, userItem := range userList {
		if userItem.Name == name && userItem.Password == password {
			authenticatedUser = &userItem
			fmt.Printf("Welcome %v\n", userItem.Name)
			return nil, authenticatedUser
		}
	}
	return errors.New("User not found"), authenticatedUser
}

func GetUserList() (userCount int, userlist []User) {
	var data = make([]byte, 500, 500)
	var user = User{}
	var userList []User
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("error on opening file : ", err)
	}
	byteread, rErr := file.Read(data)
	if rErr != nil {
		log.Fatal("error on reading file : ", err)
	}
	var dataString = string(data)
	dataString = dataString[:byteread-1]
	userSlice := strings.Split(dataString, "\n")
	for _, userItem := range userSlice {
		jdErr := json.Unmarshal([]byte(userItem), &user)
		if jdErr != nil {
			log.Fatal("error on unmarshalling user : ", jdErr)
		}
		userList = append(userList, user)
	}

	userCount = len(userList)

	return userCount, userList
}
