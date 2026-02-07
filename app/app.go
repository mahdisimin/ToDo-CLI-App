package app

import (
	"ToDo/contracts"
	"encoding/json"
	"reflect"
)

type App struct {
	Storage contracts.Storage
}

func (app *App) PersistEntity(entity any) {
	entityJson, err := json.Marshal(entity)
	if err != nil {
		panic(err)
	}
	entityJson = append(entityJson, '\n')
	entityType := reflect.TypeOf(entity)
	entityTypeName := entityType.Name()
	app.Storage.Save(string(entityJson), entityTypeName)
}
