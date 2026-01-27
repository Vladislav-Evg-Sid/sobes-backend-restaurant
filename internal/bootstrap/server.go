package bootstrap

import (
	"fmt"

	managerserviceapi "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/api/managerServiceAPI"
	"github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"
)

func AppRun(managerAPI *managerserviceapi.ManagerServiceAPI, pathToJSON string) {
	fmt.Println("AppRun start")
	fmt.Print("Reading data... ")
	clients, orders, err := readDataFromJSON(pathToJSON)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("data loaded")

	allClientsInserved := true
	for allClientsInserved {
		if len(clients) == 0 {
			fmt.Println("Done")
			return
		}
		fmt.Printf("Inserting clients (clients in queue: %d)... ", len(clients))
		err = insertClients(managerAPI, clients)
		if err != nil {
			if err.Error() == "The hall is full" {
				allClientsInserved = false
			} else {
				fmt.Println(err)
				return
			}
		}
		fmt.Println("clients inserted")

		fmt.Println("Processing clients...")
		for len(clients) > 0 {
			fmt.Printf("Clients in queue: %d. ", len(clients))
			err = processingOrders(managerAPI, clients[0], orders)
			fmt.Println("Clients processed")
			if err != nil { // TODO: Добавить проверки ошибок (например, если нет ингридиентов, то не падать)
				fmt.Println(err)
				return
			}
			clients = clients[1:]
		}
	}
}

func insertClients(managerAPI *managerserviceapi.ManagerServiceAPI, clients []models.Client) error {
	return managerAPI.SetClients(clients)
}

func processingOrders(managerAPI *managerserviceapi.ManagerServiceAPI, client models.Client, orders []models.Order) error {
	return managerAPI.ProcessingClients(client, orders)
}
