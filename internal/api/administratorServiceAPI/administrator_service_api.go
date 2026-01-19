package administratorserviceapi

type administratorService interface {
	SetClient() error    // TODO: Прописать структуры для взаимодействия
	DeleteClient() error // TODO: Прописать структуры для взаимодействия
}

type AdministratorServiceAPI struct {
	Service administratorService
}

func NewAdministratorService(service administratorService) *AdministratorServiceAPI {
	return &AdministratorServiceAPI{
		Service: service,
	}
}
