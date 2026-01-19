package adminserviceapi

type adminService interface {
	SetClient() error    // TODO: Прописать структуры для взаимодействия
	DeleteClient() error // TODO: Прописать структуры для взаимодействия
}

type AdminServiceAPI struct {
	Service adminService
}

func NewAdminService(service adminService) *AdminServiceAPI {
	return &AdminServiceAPI{
		Service: service,
	}
}
