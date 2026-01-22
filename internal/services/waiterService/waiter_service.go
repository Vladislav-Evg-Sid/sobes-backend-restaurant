package waiterservice

import "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"

type waiterStorage interface {
	GetClientInfo(clientId int) (models.Client, error)
	GetDishAgeLimit(dishName string) (int, error)
}

type kitchenService interface {
	GetDishByName(dishName string) (models.Dish, error)
}

type WaiterService struct {
	Storage waiterStorage
	Kitchen kitchenService
	// TODO: Прописать данные, которые требуется временно хранить
}

func NewWaiterService(storage waiterStorage) *WaiterService {
	return &WaiterService{
		Storage: storage,
	}
}
