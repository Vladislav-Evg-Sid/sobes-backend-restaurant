package managerserviceapi

import "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"

type managerService interface {
	SetClient(newClient models.Client) error
	DeleteClient(clientId string) error
	ProcessingClient(order models.Order) error
}

type ManagerServiceAPI struct {
	Service managerService
}

func NewManagerService(service managerService) *ManagerServiceAPI {
	return &ManagerServiceAPI{
		Service: service,
	}
}
