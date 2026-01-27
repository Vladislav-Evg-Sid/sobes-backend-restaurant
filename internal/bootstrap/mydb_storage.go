package bootstrap

import "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/storage/mydb"

func GetMydbStorage(hallMaxCapacity int) *mydb.DataBase {
	return mydb.GetDataBase(hallMaxCapacity)
}
