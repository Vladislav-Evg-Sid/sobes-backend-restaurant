package models

type ClientParced struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type OrderParced struct {
	Id     string `json:"id"`
	Client string `json:"id_client"`
	Dich   string `json:"dish"`
	Count  int    `json:"quantity"`
}
