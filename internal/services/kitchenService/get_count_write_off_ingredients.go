package kitchenservice

func (k *KitchenService) GetCountWriteOffIngredients() int {
	return k.statistic.countWriteOffIngredients
}
