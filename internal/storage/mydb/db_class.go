package mydb

type DataBase struct {
	Hall            hall
	Ingridients     []*ingridient
	Dishes          []*dish
	DishIngridients []*dishIngridients
}

func getHallStruct(maxCapacity int) hall {
	return hall{
		maxCapacity: maxCapacity,
		clients:     make([]*client, 0),
	}
}

func getIngridientsStruct() []*ingridient {
	ingr := make([]*ingridient, 0)
	ingr = append(ingr, baseInsertDataInTableIngridient(ingr)...)
	return ingr
}

func getDishesStruct() []*dish {
	dish := make([]*dish, 0)
	dish = append(dish, baseInsertDataInTableDish(dish)...)
	return dish
}

func getDishIngridientsStruct() []*dishIngridients {
	di := make([]*dishIngridients, 0)
	di = append(di, baseInsertDataInTableDishIngridien(di)...)
	return di
}

func GetDataBase(maxCapacity int) *DataBase {
	return &DataBase{
		Hall:            getHallStruct(maxCapacity),
		Ingridients:     getIngridientsStruct(),
		Dishes:          getDishesStruct(),
		DishIngridients: getDishIngridientsStruct(),
	}
}

// TODO: Переделать класс под синглтон
