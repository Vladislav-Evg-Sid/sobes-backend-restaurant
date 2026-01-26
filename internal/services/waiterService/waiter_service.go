package waiterservice

import (
	agecategories "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/enums/ageCategories"
	"github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"
)

type waiterStorage interface {
	GetClientInfo(clientId string) (models.Client, error)
	GetDishAgeLimit(dishName string) (agecategories.AgeCat, error)
}

type kitchenService interface {
	GetDishByName(dishName string) (models.DishData, error)
	ProcessingDish(dish models.DishData) error
}

type WaiterService struct {
	Storage   waiterStorage
	Kitchen   kitchenService
	clientAge int
	// TODO: Прописать данные, которые требуется временно хранить
}

func NewWaiterService(storage waiterStorage, kitchen kitchenService) *WaiterService {
	return &WaiterService{
		Storage:   storage,
		Kitchen:   kitchen,
		clientAge: -1,
	}
}
