package administratorservice

type administratorStorage interface {
	GetMaxCapacity() (int, error)
	GetCurrentClientCount() (int, error)
	SetNewClient() error // TODO: Прописать структуру клиента
}

type AdministratorService struct {
	db administratorStorage
}

func NewAdministratorService(storage administratorStorage) *AdministratorService {
	return &AdministratorService{
		db: storage,
	}
}
