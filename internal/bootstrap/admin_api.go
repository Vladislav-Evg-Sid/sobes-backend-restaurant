package bootstrap

import (
	adminserviceapi "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/api/AdminServiceAPI"
	kitchenservice "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/services/kitchenService"
	managerservice "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/services/managerService"
	waiterservice "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/services/waiterService"
	"github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/storage/mydb"
)

func GetAdminAPI(manager *managerservice.ManagerService, waiter *waiterservice.WaiterService, kitchen *kitchenservice.KitchenService, storage *mydb.DataBase) *adminserviceapi.AdminServiceAPI {
	return adminserviceapi.NewAdminServiceAPI(manager, waiter, kitchen, storage)
}
