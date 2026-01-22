package kitchenservice

type kitchenStorage interface {
	GetDishList() error      // TODO: Прописать структуры
	WriteOffProducts() error // TODO: Прописать структуры
	SetNewProducts() error
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
