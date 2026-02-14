package main

import (
	"ToDo/app"
	"ToDo/entity"
	"ToDo/storage"
	"bufio"
	"fmt"
	"os"
)

func main() {
	application := app.App{
		Storage: storage.FileHandler{},
	}
	command := getCommand()
	switch command {
	case "register":
		user, _ := entity.NewUser()
		application.PersistEntity(user)
	case "add-category":
		category, _ := entity.CreateCategory()
		application.PersistEntity(category)
	}

}

func getCommand() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter command: ")
	scanner.Scan()
	command := scanner.Text()
	return command
}
