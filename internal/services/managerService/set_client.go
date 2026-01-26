package managerservice

import "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"

func (m *ManagerService) SetClient(newClient models.Client) error {
	err := m.Waiter.SetClientAge(newClient.Age)
	if err != nil {
		return err
	}
	return m.Storage.SetNewClient(newClient)
}
