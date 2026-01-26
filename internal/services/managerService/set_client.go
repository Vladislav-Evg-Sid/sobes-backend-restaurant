package managerservice

import (
	"fmt"

	"github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"
)

func (m *ManagerService) SetClient(newClient models.Client) error {
	maxCap := m.Storage.GetMaxCapacity()
	currentClientCount := m.Storage.GetCurrentClientCount()
	if currentClientCount == maxCap {
		return fmt.Errorf("The hall is full")
	}

	err := m.Waiter.SetClientAge(newClient.Age)
	if err != nil {
		return err
	}
	return m.Storage.SetNewClient(newClient)
}
