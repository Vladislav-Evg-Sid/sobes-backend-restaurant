package main

import (
	"fmt"

	"github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/config"
	"github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/bootstrap"
)

func main() {
	cfg, err := config.LoadConfig("config.yml")
	if err != nil {
		panic(fmt.Sprintf("Ошибка загрузки конфига: %s", err))
	}

	mydbStorage := bootstrap.GetMydbStorage(cfg.Hall.MaxCapacity)
	kitchenService := bootstrap.GetKitchenService(mydbStorage, cfg.Kitchen.DishCountToSetProduct)
	waiterService := bootstrap.GetWaiterService(mydbStorage, kitchenService)
	managerService := bootstrap.GetManagerService(mydbStorage, waiterService)
	managerAPI := bootstrap.GetManagerAPI(managerService)
	adminAPI := bootstrap.GetAdminAPI(managerService, waiterService, kitchenService, mydbStorage)

	bootstrap.AppRun(managerAPI, adminAPI, cfg.App.PathToJSON)
}
