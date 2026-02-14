package entity

import (
	"ToDo/storage"
	"bufio"
	"fmt"
	"os"
)

type Category struct {
	Id    int
	Name  string
	Color string
}

func CreateCategory() (category Category, err error) {
	fileHandler := storage.FileHandler{}
	_, categoryCount, _ := fileHandler.Load("Category")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter category name : ")
	scanner.Scan()
	category.Name = scanner.Text()
	fmt.Println("Please enter category color : ")
	scanner.Scan()
	category.Color = scanner.Text()
	category.Id = categoryCount + 1

	return category, nil
}
