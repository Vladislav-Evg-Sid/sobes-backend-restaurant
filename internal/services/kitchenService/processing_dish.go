package kitchenservice

import (
	"fmt"

	"github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"
)

func (k *KitchenService) ProcessingDish(dish models.DishData) error { // TODO: Добавить автопоплнение ингридиентов
	availableIngredients := k.Storage.GetIngredients()
	for ingName, ingCount := range dish.Ingredients {
		if availableIngredients[ingName] < ingCount {
			return fmt.Errorf("Not enough ingredients")
		}
	}

	return k.Storage.WriteOffProducts(dish.Ingredients)
}
