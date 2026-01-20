package managerservice

type managerStorage interface {
	GetMaxCapacity() (int, error)
	GetCurrentClientCount() (int, error)
	SetNewClient() error // TODO: Прописать структуру клиента
}

type waiterService interface {
	ValidOrder() error             // TODO: Прописать структуру для заказа
	TransferOrderToKitchen() error // TODO: Прописать структуру
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
