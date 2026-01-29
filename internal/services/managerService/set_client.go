package managerservice

import (
	"fmt"

	"github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"
)

func (m *ManagerService) SetClient(newClient models.Client) error {
	maxCap := m.Storage.GetMaxCapacity()
	currentClientCount := m.Storage.GetCurrentClientCount()
	hallLoad := float32(currentClientCount) / float32(maxCap)

	if hallLoad > m.statistic.maxHallLoad {
		m.statistic.maxHallLoad = hallLoad
	}

	if currentClientCount == maxCap {
		m.statistic.countRefusedDueOverflowHall++
		return fmt.Errorf("The hall is full")
	}

	return m.Storage.SetNewClient(newClient)
}
