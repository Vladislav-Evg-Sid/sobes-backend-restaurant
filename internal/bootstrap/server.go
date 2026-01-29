package bootstrap

import (
	"fmt"

	managerserviceapi "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/api/managerServiceAPI"
	"github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"
)

func AppRun(managerAPI *managerserviceapi.ManagerServiceAPI, pathToJSON string) {
	restaurantProcessing(managerAPI, pathToJSON)
}

func restaurantProcessing(managerAPI *managerserviceapi.ManagerServiceAPI, pathToJSON string) {
	fmt.Println("logs: AppRun start")
	fmt.Print("logs: Reading data... ")
	clients, orders, err := readDataFromJSON(pathToJSON)
	if err != nil {
		fmt.Printf("\nError: %s\n", err)
		return
	}
	fmt.Println("data loaded")

	allClientsInserved := false
	for !allClientsInserved {
		allClientsInserved = true
		if len(clients) == 0 {
			fmt.Println("logs: Done")
			return
		}
		fmt.Printf("logs: Inserting clients (clients in queue: %d)... ", len(clients))
		insertedCount, err := insertClients(managerAPI, clients)
		clientCountToStopProcessing := 0
		if err != nil {
			if err.Error() == "The hall is full" {
				allClientsInserved = false
				fmt.Printf("\nlogs: The hall is full. Inserted clients: %d. Total clients: %d\n", insertedCount, len(clients))
				clientCountToStopProcessing = len(clients) - insertedCount
			} else {
				fmt.Printf("\nError: %s\n", err)
				return
			}
		} else {
			fmt.Println("clients inserted")
		}

		fmt.Println("logs: Processing clients...")
		for len(clients) > clientCountToStopProcessing {
			fmt.Printf("logs: Clients in queue: %d... ", len(clients))
			err = processingOrders(managerAPI, clients[0], orders)
			fmt.Println("Client processed")
			if err != nil {
				if err.Error() == "Not enough ingredients" {
					fmt.Println("logs: Not enough ingredients")
				} else if err.Error() == "Invalid age category" {
					fmt.Println("logs: Invalid age category")
				} else {
					fmt.Printf("Error: %s\n", err)
					return
				}
			}
			clients = clients[1:]
		}
	}
}

func insertClients(managerAPI *managerserviceapi.ManagerServiceAPI, clients []models.Client) (int, error) {
	return managerAPI.SetClients(clients)
}

func processingOrders(managerAPI *managerserviceapi.ManagerServiceAPI, client models.Client, orders []models.Order) error {
	return managerAPI.ProcessingClients(client, orders)
}
