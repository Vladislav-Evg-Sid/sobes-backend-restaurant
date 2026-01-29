package mydb

type statistic struct {
	writedOffIngredients ingridient
}

type DataBase struct {
	MaxHallCapacity int
	Ingredients     ingridient
	Dishes          dish
	Clients         client
	statistic       statistic
}

func getIngredientsStruct() ingridient {
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
			Ingredients:     getIngredientsStruct(),
			Dishes:          getDishesStruct(),
			Clients:         getClientsStruct(),
		}
		instance.SetNewProducts()
	}
	return instance
}
