package kitchenservice

func (k *KitchenService) autoAddingIngredients() {
	k.CurrentDishCount++
	if k.CurrentDishCount == k.DishCountToSetProducts {
		k.CurrentDishCount = 0
		k.Storage.SetNewProducts()
	}
}
