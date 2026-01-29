package managerservice

import (
	"fmt"

	"github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"
)

func (m *ManagerService) SetClient(newClient models.Client) error {
	maxCap := m.Storage.GetMaxCapacity()
	currentClientCount := m.Storage.GetCurrentClientCount()

	if currentClientCount == maxCap {
		m.statistic.maxHallLoad = 1.0
		m.statistic.countRefusedDueOverflowHall++
		return fmt.Errorf("The hall is full")
	}

	hallLoad := float32(currentClientCount+1) / float32(maxCap)
	if hallLoad > m.statistic.maxHallLoad {
		m.statistic.maxHallLoad = hallLoad
	}

	return m.Storage.SetNewClient(newClient)
}
