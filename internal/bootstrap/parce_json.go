package bootstrap

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"
)

func readDataFromJSON(pathToJSON string) ([]models.Client, []models.Order, error) {
	jsonFile, err := os.Open(pathToJSON)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	defer jsonFile.Close()
	byteData, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, nil, err
	}

	var data models.DataParced
	err = json.Unmarshal(byteData, &data)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	clients := parceClients(data.Clients)
	orders := parceOrders(data.Orders, clients)

	return clients, orders, nil
}

func parceClients(data_clients []models.ClientParced) []models.Client {
	clients := make([]models.Client, 0)
	for _, cli := range data_clients {
		clients = append(
			clients,
			models.Client{
				Id:   cli.Id,
				Name: cli.Name,
				Age:  cli.Age,
			},
		)
	}
	return clients
}

func parceOrders(data_orders []models.OrderParced, clients []models.Client) []models.Order {
	orders := make([]models.Order, 0)
	for _, ord := range data_orders {
		orders = append(
			orders,
			models.Order{
				Id:     ord.Id,
				Dish:   ord.Dich,
				Count:  ord.Count, // TODO: Переделать обработку блюда под несколько блюд
				Client: mapClientFromID(clients, ord.Client),
			},
		)
	}
	return orders
}

func mapClientFromID(clients []models.Client, clientId string) *models.Client {
	for _, cli := range clients {
		if clientId == cli.Id {
			return &cli
		}
	}
	panic("Ошибка парсинга. Нет клиента, указанного в заказах")
}
