package bootstrap

import (
	managerserviceapi "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/api/managerServiceAPI"
	managerservice "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/services/managerService"
)

func GetManagerAPI(service *managerservice.ManagerService) *managerserviceapi.ManagerServiceAPI {
	return managerserviceapi.NewManagerService(service)
}
