package kitchenservice

import "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"

type kitchenStorage interface {
	GetDishByName(dishName string) (models.DishData, error)
	WriteOffProducts(models.Ingredients) error
	SetNewProducts() error
	GetIngredients() models.Ingredients
}

type KitchenService struct {
	Storage                kitchenStorage
	DishCountToSetProducts int
	CurrentDishCount       int
}

func NewKitchenService(storage kitchenStorage, dishCountToSetProduct int) *KitchenService {
	return &KitchenService{
		Storage:                storage,
		DishCountToSetProducts: dishCountToSetProduct,
		CurrentDishCount:       0,
	}
}
