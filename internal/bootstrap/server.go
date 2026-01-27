package bootstrap

import (
	"fmt"

	managerserviceapi "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/api/managerServiceAPI"
	"github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"
)

func AppRun(managerAPI *managerserviceapi.ManagerServiceAPI, pathToJSON string) {
	clients, orders, err := readDataFromJSON(pathToJSON)
	if err != nil {
		fmt.Println(err)
		return
	}

	allClientsInserved := true
	lastClientIndex := 0
	for allClientsInserved {
		lastClientIndex, err = insertClients(managerAPI, clients, lastClientIndex)
		if err.Error() == "The hall is full" {
			allClientsInserved = false
		} else if err != nil {
			fmt.Println(err)
			return
		}

		for len(clients) > 0 {
			err = processingOrders(managerAPI, clients[0], orders)
			if err != nil {
				fmt.Println(err)
				return
			}
			clients = clients[1:]
		}
	}
}

func insertClients(managerAPI *managerserviceapi.ManagerServiceAPI, clients []models.Client, currentClientIndex int) (int, error) {
	return managerAPI.SetClients(clients[currentClientIndex:])
}

func processingOrders(managerAPI *managerserviceapi.ManagerServiceAPI, client models.Client, orders []models.Order) error {
	return managerAPI.ProcessingClients(client, orders)
}
