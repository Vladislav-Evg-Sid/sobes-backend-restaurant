package waiterservice

type waiterStorage interface {
	GetClientInfo() error   // TODO: Прописать структуры
	GetDishAgeLimit() error // TODO: Прописать структуры
}

type kitchenService interface {
	GetDishByOrder() error // TODO: Прописать структуры
}

type WaiterService struct {
	Storage waiterStorage
	Kitchen kitchenService
	// TODO: Прописать данные, которые требуется временно хранить
}

func NewWaiterService(storage waiterStorage) *WaiterService {
	return &WaiterService{
		Storage: storage,
	}
}
