package mydb

import (
	"fmt"

	"github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"
)

func (db *DataBase) SetNewClient(newClient models.Client) error {
	db.Hall.clients = append(
		db.Hall.clients,
		&client{
			id:   newClient.Id,
			name: newClient.Name,
			age:  newClient.Age,
		},
	)
	return nil
}

func (db *DataBase) WriteOffProducts(ingridientsToDel []models.Ingridient) error {
	for _, delIngridient := range ingridientsToDel {
		for i, ingridient := range db.Ingridients {
			if ingridient.id == delIngridient.Id {
				db.Ingridients = append(db.Ingridients[:i], db.Ingridients[i+1:]...)
				break
			}
		}
	}
	return nil
}

func (db *DataBase) DeleteClient(clientId string) error {
	for i, client := range db.Hall.clients {
		if client.id == clientId {
			db.Hall.clients = append(db.Hall.clients[:i], db.Hall.clients[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("No client id")
}
