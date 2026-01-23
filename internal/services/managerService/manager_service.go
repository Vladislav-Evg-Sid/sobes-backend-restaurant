package managerservice

import "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"

type managerStorage interface {
	GetMaxCapacity() int
	GetCurrentClientCount() int
	SetNewClient(models.Client) error
}

type waiterService interface {
	ValidateOrder(models.Order) error
	TransferOrderToKitchen(models.Order) error
}

type ManagerService struct {
	Storage managerStorage
	Waiter  waiterService
}

func NewManagerService(storage managerStorage, waiter waiterService) *ManagerService {
	return &ManagerService{
		Storage: storage,
		Waiter:  waiter,
	}
}
