package main

import (
	"github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/config"
	"github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/bootstrap"
)

func main() {
	cfg, err := config.LoadConfig("config.yml")
	if err != nil {
		panic(err)
	}

	mydbStorage := bootstrap.GetMydbStorage(cfg.Hall.MaxCapacity)
	kitchenService := bootstrap.GetKitchenService(mydbStorage, cfg.Kitchen.DishCountToSetProduct)
	waiterWervice := bootstrap.GetWaiterService(mydbStorage, kitchenService)
	managerService := bootstrap.GetManagerService(mydbStorage, waiterWervice)
	managerAPI := bootstrap.GetManagerAPI(managerService)

	bootstrap.AppRun(managerAPI)
}
