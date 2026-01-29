package kitchenservice

func (k *KitchenService) GetCountSuccessProcessedOrders() int {
	return k.statistic.countSuccessProcessedOrders
}
