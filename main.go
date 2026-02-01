package main

import (
	cat "ToDo/category"
	task "ToDo/task"
	user "ToDo/user"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Welcome to TODO Application")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Please Login or register")
		for true {
			scanner.Scan()
			command := scanner.Text()
			if command == "login" || command == "register" {
				getCommand(command)
			} else {
				fmt.Println("Please Login or register")
			}

		}

	}
}

func getCommand(command string) {
	switch command {
	case "register":
		user.User{}.RegisterUserMethod()
	case "create-category":
		cat.Category{}.CreateCategory()
	case "login":
		lErr, _ := user.User{}.LoginUser()
		if lErr != nil {
			fmt.Println(lErr)
			log.Fatalf(lErr.Error())
		}
	case "create-task":
	case "exit":
		task.Task{}.CreatTask()
		os.Exit(0)
	default:
		fmt.Println("Invalid command")
	}

}
