package managerserviceapi

import (
	"fmt"

	"github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"
)

func (m *ManagerServiceAPI) ProcessingClients(client models.Client, orders []models.Order) error {
	for _, ord := range orders {
		if ord.Client.Id == client.Id {
			return m.Service.ProcessingClient(ord)
		}
	}
	return fmt.Errorf("No client's order")
}
