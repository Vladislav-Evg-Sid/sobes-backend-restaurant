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

func (db *DataBase) WriteOffProducts(ingridientsToWriteOff models.Ingridients) error {
	for keyDel, valueDel := range ingridientsToWriteOff {
		for keyCurrent := range db.Ingridients {
			if keyCurrent == keyDel {
				db.Ingridients[keyCurrent] -= valueDel
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
