package managerserviceapi

import "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"

func (m *ManagerServiceAPI) SetClients(newClients []models.Client) error {
	for _, client := range newClients {
		err := m.Service.SetClient(client)
		if err != nil {
			return err
		}
	}
	return nil
}
