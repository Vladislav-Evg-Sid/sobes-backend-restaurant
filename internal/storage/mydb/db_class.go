package mydb

type DataBase struct {
	MaxHallCapacity int
	Ingridients     ingridient
	Dishes          dish
	Clients         client
}

func getIngridientsStruct() ingridient {
	ingr := make(ingridient)
	ingr = insertDataInIngredients(ingr)
	return ingr
}

func getDishesStruct() dish {
	dish := make(dish)
	dish = insertDataInDishes(dish)
	return dish
}

func getClientsStruct() client {
	return make(client)
}

var instance *DataBase = nil

func GetDataBase(maxCapacity int) *DataBase {
	if instance == nil {
		instance = &DataBase{
			MaxHallCapacity: maxCapacity,
			Ingridients:     getIngridientsStruct(),
			Dishes:          getDishesStruct(),
			Clients:         getClientsStruct(),
		}
	}
	return instance
}
