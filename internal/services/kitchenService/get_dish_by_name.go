package kitchenservice

import "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"

func (k *KitchenService) GetDishByName(dishName string) (models.Dishes, error) {
	return k.Storage.GetDishByName(dishName)
}
