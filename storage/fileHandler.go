package storage

import (
	"ToDo/contracts"
	"errors"
	"os"
	"strings"
)

type FileHandler struct {
	path string
}

func (file FileHandler) Save(data string, entity string) {
	switch entity {
	case "User":
		file.path = contracts.UserFilePath
	case "Task":
		file.path = contracts.TaskFilePath
	case "Category":
		file.path = contracts.CategoryFilePath

	}
	f, oErr := os.OpenFile(file.path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if oErr != nil {
		errors.New(oErr.Error())
	}
	defer f.Close()
	f.Write([]byte(data))
}

func (file FileHandler) Load(entity string) (string, int, error) {
	var dataByte = make([]byte, 1024)
	switch entity {
	case "User":
		file.path = contracts.UserFilePath
	case "task":
		file.path = contracts.TaskFilePath
	case "category":
		file.path = contracts.CategoryFilePath

	}
	f, oErr := os.OpenFile(file.path, os.O_RDONLY, 0777)
	if oErr != nil {
		return "", 0, oErr
	}
	defer f.Close()
	dataLength, _ := f.Read(dataByte)
	dataByte = dataByte[:dataLength-1]
	dataSting := string(dataByte)
	dataSting = strings.Trim(dataSting, "\n")
	dataSlice := strings.Split(dataSting, "\n")
	dataCount := len(dataSlice)
	return dataSting, dataCount, nil
}
