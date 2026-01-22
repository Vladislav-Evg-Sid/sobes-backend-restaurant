package managerserviceapi

import "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"

type managerService interface {
	SetClient(models.Client) error
	ProcessingClient(models.Order) error
}

type ManagerServiceAPI struct {
	Service managerService
}

func NewManagerService(service managerService) *ManagerServiceAPI {
	return &ManagerServiceAPI{
		Service: service,
	}
}
