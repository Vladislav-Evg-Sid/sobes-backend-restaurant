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

	return m.Storage.SetNewClient(newClient)
}
