package managerservice

import "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"

func (m *ManagerService) ProcessingClient(order models.Order) error {
	err := m.Waiter.ValidateOrder(order)
	if err != nil {
		return err
	}

	err = m.Waiter.TransferOrderToKitchen(order)
	if err != nil {
		return err
	}

	return m.DeleteClient(order.Client.Id)
}
