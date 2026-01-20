package managerserviceapi

type managerService interface {
	SetClient() error        // TODO: Прописать структуры для взаимодействия
	ProcessingClient() error // TODO: Сделать структуру для принятия клиента
}

type ManagerServiceAPI struct {
	Service managerService
}

func NewManagerService(service managerService) *ManagerServiceAPI {
	return &ManagerServiceAPI{
		Service: service,
	}
}
