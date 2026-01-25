package mydb

import (
	"fmt"

	"github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"
)

func (db *DataBase) SetNewClient(newClient models.Client) error {
	db.Clients[newClient.Id] = clientData{
		name: newClient.Name,
		age:  newClient.Age,
	}
	return nil
}

func (db *DataBase) WriteOffProducts(ingredientsToWriteOff models.Ingredients) error {
	for keyDel, valueDel := range ingredientsToWriteOff {
		for keyCurrent := range db.Ingredients {
			if keyCurrent == keyDel {
				db.Ingredients[keyCurrent] -= valueDel
			}
		}
	}
	return nil
}

func (db *DataBase) DeleteClient(clientId string) error {
	if _, exist := db.Clients[clientId]; exist {
		delete(db.Clients, clientId)
		return nil
	} else {
		return fmt.Errorf("No client id")
	}
}

func (db *DataBase) SetNewProducts() error {
	valueSetter := getAddedIngredients()
	for key, value := range valueSetter {
		db.Ingredients[key] += value
	}
	return nil
}
