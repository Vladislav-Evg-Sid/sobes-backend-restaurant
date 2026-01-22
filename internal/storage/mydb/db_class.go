package mydb

type DataBase struct {
	Hall        hall
	Ingridients []*ingridient
	Dishes      []*dish
}

func GetDataBase() *DataBase {
	db := DataBase{}
	// TODO: Реализовать функции для стартового заполнения БД
	return &db
}

// TODO: Переделать класс под синглтон
