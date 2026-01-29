package bootstrap

import (
	"fmt"
	"math"

	adminserviceapi "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/api/AdminServiceAPI"
	managerserviceapi "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/api/managerServiceAPI"
	"github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"
)

func AppRun(managerAPI *managerserviceapi.ManagerServiceAPI, adminAPI *adminserviceapi.AdminServiceAPI, pathToJSON string) {
	restaurantProcessing(managerAPI, pathToJSON)
	adminProcessing(adminAPI)
}

func adminProcessing(adminAPI *adminserviceapi.AdminServiceAPI) {
	fmt.Println("\nРабочий день закончился, админ смотрит результаты:")
	fmt.Println("\tАнализ по манагеру:")
	fmt.Printf("\t\tКоличество обслуженных клиентов: %d\n", adminAPI.ManagerService.GetCountServedClients())
	fmt.Printf("\t\tМаксимальная нагрузка на зал: %f%%\n", math.Round(float64(100*adminAPI.ManagerService.GetMaxHallLoad())))
	fmt.Printf("\t\tКоличество отказов по причине переполненности: %d\n", adminAPI.ManagerService.GetCountRefusedDueOverflowHall())

	fmt.Println("\tАнализ по официанту:")
	fmt.Printf("\t\tОбщее количество обслуженных заказов: %d\n", adminAPI.WaiterService.GetCountServedOrders())
	fmt.Printf("\t\tКоличество отклонённых заказов: %d\n", adminAPI.WaiterService.GetCountRefusedOrders())
	fmt.Printf("\t\tКоличество заказов с 18+ блюдами: %d\n", adminAPI.WaiterService.GetCountMore18Dishes())
	fmt.Printf("\t\tСреднее количество блюд в заказе: %f\n", math.Round(float64(100*adminAPI.WaiterService.GetMeanCountDishesInOrder()))/100)

	fmt.Println("\tАнализ по кухне:")
	fmt.Printf("\t\tКоличество успешно приготовленных заказов: %d\n", adminAPI.KitchenService.GetCountSuccessProcessedOrders())
	fmt.Printf("\t\tКоличество отказов из-за отсутствия ингредиентов: %d\n", adminAPI.KitchenService.GetCountRefusedOrdersWithoutIngredients())
	fmt.Printf("\t\tСамые часто готовимые блюда: %s\n", arr2str(adminAPI.KitchenService.GetTop5PopularDishes()))
	fmt.Printf("\t\tОбщее количество списанных ингредиентов: %d\n", adminAPI.KitchenService.GetCountWriteOffIngredients())

	fmt.Println("\tАнализ по складу:")
	fmt.Printf("\t\tТекущие остатки ингредиентов: %s\n", modelsIngredients2strList(adminAPI.StorageService.GetCurrentIngredientsCount()))
	fmt.Printf("\t\tИнгредиенты с минимальным запасом: %s\n", arr2str(adminAPI.StorageService.GetTop5IngredientsMinCount()))
	fmt.Printf("\t\tИнгредиенты с максимальным запасом: %s\n", arr2str(adminAPI.StorageService.GetTop5IngredientsMaxCount()))
	fmt.Printf("\t\tКоличество списаний: %s\n", modelsIngredients2strList(adminAPI.StorageService.GetCountWriteOffIngredients()))
	fmt.Printf("\t\tОбщее количество всех ингредиентов: %d\n", adminAPI.StorageService.GetCountAllIngredients())
}

func modelsIngredients2strList(ing models.Ingredients) string {
	ingStr := ""
	for name, count := range ing {
		ingStr += fmt.Sprintf("\n\t\t\t%s: %d", name, count)
	}
	return ingStr
}

func arr2str(productList [5]models.ProductCount) string {
	productsStr := ""
	for _, dish := range productList {
		productsStr += fmt.Sprintf("%s (%d); ", dish.Name, dish.Count)
	}
	return productsStr
}

func restaurantProcessing(managerAPI *managerserviceapi.ManagerServiceAPI, pathToJSON string) {
	fmt.Println("logs: AppRun start")
	fmt.Print("logs: Reading data... ")
	clients, orders, err := readDataFromJSON(pathToJSON)
	if err != nil {
		fmt.Printf("\nError: %s\n", err)
		return
	}
	fmt.Println("data loaded")

	allClientsInserved := false
	for !allClientsInserved {
		allClientsInserved = true
		if len(clients) == 0 {
			fmt.Println("logs: Done")
			return
		}
		fmt.Printf("logs: Inserting clients (clients in queue: %d)... ", len(clients))
		insertedCount, err := insertClients(managerAPI, clients)
		clientCountToStopProcessing := 0
		if err != nil {
			if err.Error() == "The hall is full" {
				allClientsInserved = false
				fmt.Printf("\nlogs: The hall is full. Inserted clients: %d. Total clients: %d\n", insertedCount, len(clients))
				clientCountToStopProcessing = len(clients) - insertedCount
			} else {
				fmt.Printf("\nError: %s\n", err)
				return
			}
		} else {
			fmt.Println("clients inserted")
		}

		fmt.Println("logs: Processing clients...")
		for len(clients) > clientCountToStopProcessing {
			fmt.Printf("logs: Clients in queue: %d... ", len(clients))
			err = processingOrders(managerAPI, clients[0], orders)
			fmt.Println("Client processed")
			if err != nil {
				if err.Error() == "Not enough ingredients" {
					fmt.Println("logs: Not enough ingredients")
				} else if err.Error() == "Invalid age category" {
					fmt.Println("logs: Invalid age category")
				} else {
					fmt.Printf("Error: %s\n", err)
					return
				}
			}
			clients = clients[1:]
		}
	}
}

func insertClients(managerAPI *managerserviceapi.ManagerServiceAPI, clients []models.Client) (int, error) {
	return managerAPI.SetClients(clients)
}

func processingOrders(managerAPI *managerserviceapi.ManagerServiceAPI, client models.Client, orders []models.Order) error {
	return managerAPI.ProcessingClients(client, orders)
}
