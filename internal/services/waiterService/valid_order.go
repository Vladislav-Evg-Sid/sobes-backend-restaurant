package waiterservice

import (
	"fmt"

	agecategories "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/enums/ageCategories"
	"github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"
)

func (w *WaiterService) ValidateOrder(currentOrder models.Order) error {
	if w.clientAge < 18 {
		dish, err := w.Kitchen.GetDishByName(currentOrder.Dish)
		if err != nil {
			return err
		}
		if dish.AgeCategory == agecategories.More18 {
			return fmt.Errorf("Not enough ingredients")
		}
	}
	return nil
}
