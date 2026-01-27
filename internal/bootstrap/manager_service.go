package bootstrap

import (
	managerservice "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/services/managerService"
	waiterservice "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/services/waiterService"
	"github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/storage/mydb"
)

func GetManagerService(storage *mydb.DataBase, waiter *waiterservice.WaiterService) *managerservice.ManagerService {
	return managerservice.NewManagerService(storage, waiter)
}
