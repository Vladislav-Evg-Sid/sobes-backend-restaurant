package managerservice

import "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"

type managerStorage interface {
	GetMaxCapacity() int
	GetCurrentClientCount() int
	SetNewClient(newClient models.Client) error
	DeleteClient(clientId string) error
}

type waiterService interface {
	ValidateOrder(currentOrder models.Order) error
	TransferOrderToKitchen(currentOrder models.Order) error
	SetClientAge(clientAge int) error
	DeleteClientAge() error
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
