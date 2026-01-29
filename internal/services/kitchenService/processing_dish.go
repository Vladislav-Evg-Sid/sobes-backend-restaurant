package kitchenservice

import (
	"fmt"

	"github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"
)

func (k *KitchenService) ProcessingDish(dish models.DishData, dishName string, dishCount int) error {
	k.autoAddingIngredients()
	availableIngredients := k.Storage.GetIngredients()
	for ingName, ingCount := range dish.Ingredients {
		if availableIngredients[ingName] < ingCount*dishCount {
			k.statistic.countRefusedOrdersWithoutIngredients++
			return fmt.Errorf("Not enough ingredients")
		}
	}

	k.statistic.countSuccessProcessedOrders++
	k.statistic.topProcessingDishes[dishName] += dishCount
	for _, ingCount := range dish.Ingredients {
		k.statistic.countWriteOffIngredients += ingCount * dishCount
	}

	return k.Storage.WriteOffProducts(dish.Ingredients, dishCount)
}
