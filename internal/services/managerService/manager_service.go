package managerservice

type managerStorage interface {
	GetMaxCapacity() (int, error)
	GetCurrentClientCount() (int, error)
	SetNewClient() error // TODO: Прописать структуру клиента
}

type ManagerService struct {
	db managerStorage
}

func NewManagerService(storage managerStorage) *ManagerService {
	return &ManagerService{
		db: storage,
	}
}
