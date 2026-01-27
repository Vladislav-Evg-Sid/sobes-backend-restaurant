package bootstrap

import (
	kitchenservice "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/services/kitchenService"
	waiterservice "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/services/waiterService"
	"github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/storage/mydb"
)

func GetWaiterService(storage *mydb.DataBase, kitchen *kitchenservice.KitchenService) *waiterservice.WaiterService {
	return waiterservice.NewWaiterService(storage, kitchen)
}
