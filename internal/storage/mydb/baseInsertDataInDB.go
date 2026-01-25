package mydb

import agecategories "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/enums/ageCategories"

func insertDataInIngredients(ingredients ingridient) ingridient {
	ingredients["Яйцо куриное"] = 0
	ingredients["Молоко"] = 0
	ingredients["Сыр твердый"] = 0
	ingredients["Соль"] = 0
	ingredients["Масло сливочное"] = 0
	ingredients["Помидор"] = 0
	ingredients["Моцарелла"] = 0
	ingredients["Базилик свежий"] = 0
	ingredients["Масло оливковое"] = 0
	ingredients["Хлеб для тостов"] = 0
	ingredients["Авокадо"] = 0
	ingredients["Лимонный сок"] = 0
	ingredients["Перец черный молотый"] = 0
	ingredients["Спагетти"] = 0
	ingredients["Чеснок"] = 0
	ingredients["Перец чили хлопьями"] = 0
	ingredients["Петрушка свежая"] = 0
	ingredients["Йогурт греческий"] = 0
	ingredients["Мёд"] = 0
	ingredients["Грецкий орех"] = 0
	ingredients["Багет"] = 0
	ingredients["Красное сухое вино"] = 0
	ingredients["Сахар"] = 0
	ingredients["Гвоздика"] = 0
	ingredients["Корица"] = 0
	ingredients["Апельсин"] = 0
	ingredients["Мука пшеничная"] = 0
	ingredients["Томатный соус"] = 0
	ingredients["Бекон или панчетта"] = 0
	ingredients["Сливки 20%"] = 0
	ingredients["Сыр пармезан"] = 0
	ingredients["Креветки очищенные"] = 0
	ingredients["Белый ром"] = 0
	ingredients["Содовая вода"] = 0
	ingredients["Сахарный сироп"] = 0
	ingredients["Лайм"] = 0
	ingredients["Мята свежая"] = 0
	ingredients["Лёд"] = 0

	return ingredients
}

func insertDataInDishes(dishes dish) dish {
	// Омлет с сыром
	dishes["Омлет с сыром"] = dishData{
		ageCategory: agecategories.Lower18,
		ingredients: ingridient{
			"Яйцо куриное":    2,
			"Молоко":          2,
			"Сыр твердый":     30,
			"Соль":            5,
			"Масло сливочное": 10,
		},
	}

	// Салат "Капрезе"
	dishes["Салат \"Капрезе\""] = dishData{
		ageCategory: agecategories.Lower18,
		ingredients: ingridient{
			"Помидор":         2,
			"Моцарелла":       125,
			"Базилик свежий":  10,
			"Масло оливковое": 2,
			"Соль":            5,
		},
	}

	// Тост с авокадо
	dishes["Тост с авокадо"] = dishData{
		ageCategory: agecategories.Lower18,
		ingredients: ingridient{
			"Хлеб для тостов":      2,
			"Авокадо":              1,
			"Лимонный сок":         1,
			"Соль":                 5,
			"Перец черный молотый": 5,
		},
	}

	// Паста "Аглио э олио"
	dishes["Паста \"Аглио э олио\""] = dishData{
		ageCategory: agecategories.Lower18,
		ingredients: ingridient{
			"Спагетти":            200,
			"Чеснок":              4,
			"Перец чили хлопьями": 1,
			"Петрушка свежая":     3,
			"Масло оливковое":     4,
		},
	}

	// Греческий йогурт с мёдом и орехами
	dishes["Греческий йогурт с мёдом и орехами"] = dishData{
		ageCategory: agecategories.Lower18,
		ingredients: ingridient{
			"Йогурт греческий": 200,
			"Мёд":              2,
			"Грецкий орех":     20,
		},
	}

	// Брускетта с томатами
	dishes["Брускетта с томатами"] = dishData{
		ageCategory: agecategories.Lower18,
		ingredients: ingridient{
			"Багет":           1,
			"Помидор":         2,
			"Чеснок":          1,
			"Базилик свежий":  5,
			"Масло оливковое": 1,
		},
	}

	// Пицца "Маргарита"
	dishes["Пицца \"Маргарита\""] = dishData{
		ageCategory: agecategories.Lower18,
		ingredients: ingridient{
			"Мука пшеничная":  200,
			"Томатный соус":   5,
			"Моцарелла":       150,
			"Базилик свежий":  5,
			"Масло оливковое": 1,
		},
	}

	// Глинтвейн
	dishes["Глинтвейн"] = dishData{
		ageCategory: agecategories.More18,
		ingredients: ingridient{
			"Красное сухое вино": 750,
			"Сахар":              2,
			"Гвоздика":           3,
			"Корица":             1,
			"Апельсин":           1,
		},
	}

	// Паста "Карбонара"
	dishes["Паста \"Карбонара\""] = dishData{
		ageCategory: agecategories.More18,
		ingredients: ingridient{
			"Спагетти":           200,
			"Бекон или панчетта": 100,
			"Яйцо куриное":       2,
			"Сыр пармезан":       50,
			"Сливки 20%":         100,
		},
	}

	// Креветки в чесночном соусе
	dishes["Креветки в чесночном соусе"] = dishData{
		ageCategory: agecategories.More18,
		ingredients: ingridient{
			"Креветки очищенные": 300,
			"Чеснок":             3,
			"Масло сливочное":    50,
			"Петрушка свежая":    2,
			"Красное сухое вино": 50,
		},
	}

	// Коктейль "Мохито"
	dishes["Коктейль \"Мохито\""] = dishData{
		ageCategory: agecategories.More18,
		ingredients: ingridient{
			"Белый ром":      50,
			"Содовая вода":   100,
			"Сахарный сироп": 20,
			"Лайм":           1,
			"Мята свежая":    10,
			"Лёд":            100,
		},
	}

	return dishes
}

func getAddedIngredients() ingridient {
	return ingridient{
		"Сыр твердый":          90,
		"Соль":                 45,
		"Базилик свежий":       60,
		"Перец черный молотый": 15,
		"Чеснок":               24,
		"Мята свежая":          30,
		"Корица":               3,
		"Молоко":               6,
		"Спагетти":             1200,
		"Багет":                3,
		"Гвоздика":             9,
		"Белый ром":            150,
		"Масло сливочное":      180,
		"Авокадо":              3,
		"Йогурт греческий":     600,
		"Грецкий орех":         60,
		"Сливки 20%":           300,
		"Хлеб для тостов":      6,
		"Лимонный сок":         3,
		"Бекон или панчетта":   300,
		"Сыр пармезан":         150,
		"Креветки очищенные":   900,
		"Лёд":                  300,
		"Красное сухое вино":   2400,
		"Сахарный сироп":       60,
		"Апельсин":             3,
		"Моцарелла":            825,
		"Перец чили хлопьями":  3,
		"Петрушка свежая":      15,
		"Мёд":                  6,
		"Содовая вода":         300,
		"Яйцо куриное":         12,
		"Помидор":              12,
		"Масло оливковое":      24,
		"Сахар":                6,
		"Мука пшеничная":       600,
		"Томатный соус":        15,
		"Лайм":                 3,
	}
}
