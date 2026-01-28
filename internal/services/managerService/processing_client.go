package managerservice

import (
	"fmt"

	"github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"
)

func (m *ManagerService) ProcessingClient(order models.Order) error {
	err := m.Waiter.SetClientAge(order.Client.Age)
	if err != nil {
		return err
	}

	err = m.Waiter.ValidateOrder(order)
	if err != nil {
		return err
	}

	err = m.Waiter.TransferOrderToKitchen(order)
	if err != nil {
		if err.Error() == "Not enough ingredients" {
			errDelC := m.DeleteClient(order.Client.Id)
			if errDelC != nil {
				return fmt.Errorf("%s, %s", err, errDelC)
			}
		}
		return err
	}

	return m.DeleteClient(order.Client.Id)
}
