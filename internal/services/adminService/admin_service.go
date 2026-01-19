package adminservice

type adminStorage interface {
	GetMaxCapacity() (int, error)
	GetCurrentClientCount() (int, error)
}

type AdminService struct {
	db adminStorage
}

func NewAdminService(storage adminStorage) *AdminService {
	return &AdminService{
		db: storage,
	}
}
