package main

import (
	"ToDo/entity"
	"ToDo/storage"
	"bufio"
	"os"
)

var inMemoryStorage = &storage.InMemoryMap{}

func main() {
	command := getCommand()

	switch command {
	case "register":
		entity.User{}.RegisterUserMethod(inMemoryStorage)
	case "register_new":
	}
	return
}

func getCommand() string {
	scanner := bufio.NewScanner(os.Stdin)
	println("Enter command: ")
	scanner.Scan()
	command := scanner.Text()
	return command
}
