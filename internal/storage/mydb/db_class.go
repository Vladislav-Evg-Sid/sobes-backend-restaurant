package mydb

type DataBase struct {
	hallMaxCapacity int
	clients         int // TODO: Прописать структуру для хранения клиентов
}

func GetDataBase() *DataBase {
	db := DataBase{}
	// TODO: Реализовать функции для стартового заполнения БД
	return &db
}

// TODO: Переделать класс под синглтон
