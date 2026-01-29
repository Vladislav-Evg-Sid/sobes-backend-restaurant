package kitchenservice

func (k *KitchenService) GetCountRefusedOrdersWithoutIngredients() int {
	return k.statistic.countRefusedOrdersWithoutIngredients
}
