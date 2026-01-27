package bootstrap

import (
	kitchenservice "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/services/kitchenService"
	"github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/storage/mydb"
)

func GetKitchenService(storage *mydb.DataBase, dishCountToSetProduct int) *kitchenservice.KitchenService {
	return kitchenservice.NewKitchenService(storage, dishCountToSetProduct)
}
