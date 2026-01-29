package adminserviceapi

import "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"

type managerService interface {
	GetCountServedClients() int
	GetMaxHallLoad() float32
	GetCountRefusedDueOverflowHall() int
}

type waiterService interface {
	GetCountServedOrders() int
	GetCountRefusedOrders() int
	GetCountMore18Dishes() int
	GetMeanCountDishesInOrder() float32
}

type kitchenService interface {
	GetCountSuccessProcessedOrders() int
	GetCountRefusedOrdersWithoutIngredients() int
	GetTop5PopularDishes() [5]models.ProductCount
	GetCountWriteOffIngredients() int
}

type storageService interface {
	GetCurrentIngredientsCount() models.Ingredients
	GetTop5IngredientsMinCount() [5]models.ProductCount
	GetTop5IngredientsMaxCount() [5]models.ProductCount
	GetCountWriteOffIngredients() models.Ingredients
	GetCountAllIngredients() int
}

type AdminServiceAPI struct {
	ManagerService managerService
	WaiterService  waiterService
	KitchenService kitchenService
	StorageService storageService
}

func NewAdminServiceAPI(manager managerService, waiter waiterService, kitchen kitchenService, storage storageService) *AdminServiceAPI {
	return &AdminServiceAPI{
		ManagerService: manager,
		WaiterService:  waiter,
		KitchenService: kitchen,
		StorageService: storage,
	}
}
