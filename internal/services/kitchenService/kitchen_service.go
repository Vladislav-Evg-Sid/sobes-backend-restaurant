package kitchenservice

import "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"

type kitchenStorage interface {
	GetDishByName(dishName string) (models.Dishes, error)
	WriteOffProducts(models.Ingredients) error
	SetNewProducts() error
	GetIngredients() models.Ingredients
}

type KitchenService struct {
	Storage                kitchenStorage
	DishCountToSetProducts int
	CurrentDishCount       int
}

func NewKitchenService(storage kitchenStorage, maxDish int) *KitchenService {
	return &KitchenService{
		Storage:                storage,
		DishCountToSetProducts: maxDish,
		CurrentDishCount:       0,
	}
}
