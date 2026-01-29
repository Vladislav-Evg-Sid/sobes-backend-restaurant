package kitchenservice

import "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"

type kitchenStorage interface {
	GetDishByName(dishName string) (models.DishData, error)
	WriteOffProducts(ingredientsToWriteOff models.Ingredients, dishCount int) error
	SetNewProducts() error
	GetIngredients() models.Ingredients
}

type statistic struct {
	countSuccessProcessedOrders          int
	countRefusedOrdersWithoutIngredients int
	topProcessingDishes                  map[string]int
	countWriteOffIngredients             int
}

type KitchenService struct {
	Storage                kitchenStorage
	DishCountToSetProducts int
	CurrentDishCount       int
	statistic              statistic
}

func NewKitchenService(storage kitchenStorage, dishCountToSetProduct int) *KitchenService {
	return &KitchenService{
		Storage:                storage,
		DishCountToSetProducts: dishCountToSetProduct,
		CurrentDishCount:       0,
		statistic:              statistic{topProcessingDishes: make(map[string]int)},
	}
}
