package managerserviceapi

import "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"

func (m *ManagerServiceAPI) SetClients(newClients []models.Client) (int, error) {
	for count, client := range newClients {
		err := m.Service.SetClient(client)
		if err != nil {
			return count, err
		}
	}
	return 0, nil
}
