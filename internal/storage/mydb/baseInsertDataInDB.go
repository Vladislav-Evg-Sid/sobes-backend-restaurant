package mydb

import agecategories "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/enums/ageCategories"

func baseInsertDataInTableIngridient(table []*ingridient) []*ingridient {
	return append(
		table,
		&ingridient{
			id:    1,
			name:  "Яйцо куриное",
			count: 0,
		},
		&ingridient{
			id:    2,
			name:  "Молоко",
			count: 0,
		},
		&ingridient{
			id:    3,
			name:  "Сыр твердый",
			count: 0,
		},
		&ingridient{
			id:    4,
			name:  "Соль",
			count: 0,
		},
		&ingridient{
			id:    5,
			name:  "Масло сливочное",
			count: 0,
		},
		&ingridient{
			id:    6,
			name:  "Помидор",
			count: 0,
		},
		&ingridient{
			id:    7,
			name:  "Моцарелла",
			count: 0,
		},
		&ingridient{
			id:    8,
			name:  "Базилик свежий",
			count: 0,
		},
		&ingridient{
			id:    9,
			name:  "Масло оливковое",
			count: 0,
		},
		&ingridient{
			id:    10,
			name:  "Хлеб для тостов",
			count: 0,
		},
		&ingridient{
			id:    11,
			name:  "Авокадо",
			count: 0,
		},
		&ingridient{
			id:    12,
			name:  "Лимонный сок",
			count: 0,
		},
		&ingridient{
			id:    13,
			name:  "Перец черный молотый",
			count: 0,
		},
		&ingridient{
			id:    14,
			name:  "Спагетти",
			count: 0,
		},
		&ingridient{
			id:    15,
			name:  "Чеснок",
			count: 0,
		},
		&ingridient{
			id:    16,
			name:  "Перец чили хлопьями",
			count: 0,
		},
		&ingridient{
			id:    17,
			name:  "Петрушка свежая",
			count: 0,
		},
		&ingridient{
			id:    18,
			name:  "Йогурт греческий",
			count: 0,
		},
		&ingridient{
			id:    19,
			name:  "Мёд",
			count: 0,
		},
		&ingridient{
			id:    20,
			name:  "Грецкий орех",
			count: 0,
		},
		&ingridient{
			id:    21,
			name:  "Багет",
			count: 0,
		},
		&ingridient{
			id:    22,
			name:  "Красное сухое вино",
			count: 0,
		},
		&ingridient{
			id:    23,
			name:  "Сахар",
			count: 0,
		},
		&ingridient{
			id:    24,
			name:  "Гвоздика",
			count: 0,
		},
		&ingridient{
			id:    25,
			name:  "Корица",
			count: 0,
		},
		&ingridient{
			id:    26,
			name:  "Апельсин",
			count: 0,
		},
		&ingridient{
			id:    27,
			name:  "Мука пшеничная",
			count: 0,
		},
		&ingridient{
			id:    28,
			name:  "Томатный соус",
			count: 0,
		},
		&ingridient{
			id:    29,
			name:  "Бекон или панчетта",
			count: 0,
		},
		&ingridient{
			id:    30,
			name:  "Сливки 20%",
			count: 0,
		},
		&ingridient{
			id:    31,
			name:  "Сыр пармезан",
			count: 0,
		},
		&ingridient{
			id:    32,
			name:  "Креветки очищенные",
			count: 0,
		},
		&ingridient{
			id:    33,
			name:  "Белый ром",
			count: 0,
		},
		&ingridient{
			id:    34,
			name:  "Содовая вода",
			count: 0,
		},
		&ingridient{
			id:    35,
			name:  "Сахарный сироп",
			count: 0,
		},
		&ingridient{
			id:    36,
			name:  "Лайм",
			count: 0,
		},
		&ingridient{
			id:    37,
			name:  "Мята свежая",
			count: 0,
		},
		&ingridient{
			id:    38,
			name:  "Лёд",
			count: 0,
		},
	)
}

func baseInsertDataInTableDish(table []*dish) []*dish {
	return append(
		table,
		&dish{
			id:          1,
			name:        "Омлет с сыром",
			ageCategory: agecategories.Lower18,
		},
		&dish{
			id:          2,
			name:        "Салат \"Капрезе\"",
			ageCategory: agecategories.Lower18,
		},
		&dish{
			id:          3,
			name:        "Тост с авокадо",
			ageCategory: agecategories.Lower18,
		},
		&dish{
			id:          4,
			name:        "Паста \"Аглио э олио\"",
			ageCategory: agecategories.Lower18,
		},
		&dish{
			id:          5,
			name:        "Греческий йогурт с мёдом и орехами",
			ageCategory: agecategories.Lower18,
		},
		&dish{
			id:          6,
			name:        "Брускетта с томатами",
			ageCategory: agecategories.Lower18,
		},
		&dish{
			id:          7,
			name:        "Пицца \"Маргарита\"",
			ageCategory: agecategories.Lower18,
		},
		&dish{
			id:          8,
			name:        "Глинтвейн",
			ageCategory: agecategories.More18,
		},
		&dish{
			id:          9,
			name:        "Паста \"Карбонара\"",
			ageCategory: agecategories.More18,
		},
		&dish{
			id:          10,
			name:        "Креветки в чесночном соусе",
			ageCategory: agecategories.More18,
		},
		&dish{
			id:          11,
			name:        "Коктейль \"Мохито\"",
			ageCategory: agecategories.More18,
		},
	)
}

func baseInsertDataInTableDishIngridien(table []*dishIngridients) []*dishIngridients {
	return append(
		table,
		&dishIngridients{
			dishID:       1,
			ingridientId: 1,
			count:        2,
		},
		&dishIngridients{
			dishID:       1,
			ingridientId: 2,
			count:        2,
		},
		&dishIngridients{
			dishID:       1,
			ingridientId: 3,
			count:        30,
		},
		&dishIngridients{
			dishID:       1,
			ingridientId: 4,
			count:        5,
		},
		&dishIngridients{
			dishID:       1,
			ingridientId: 5,
			count:        10,
		},

		&dishIngridients{
			dishID:       2,
			ingridientId: 6,
			count:        2,
		},
		&dishIngridients{
			dishID:       2,
			ingridientId: 7,
			count:        125,
		},
		&dishIngridients{
			dishID:       2,
			ingridientId: 8,
			count:        10,
		},
		&dishIngridients{
			dishID:       2,
			ingridientId: 9,
			count:        2,
		},
		&dishIngridients{
			dishID:       2,
			ingridientId: 4,
			count:        5,
		},

		&dishIngridients{
			dishID:       3,
			ingridientId: 10,
			count:        2,
		},
		&dishIngridients{
			dishID:       3,
			ingridientId: 11,
			count:        1,
		},
		&dishIngridients{
			dishID:       3,
			ingridientId: 12,
			count:        1,
		},
		&dishIngridients{
			dishID:       3,
			ingridientId: 4,
			count:        5,
		},
		&dishIngridients{
			dishID:       3,
			ingridientId: 13,
			count:        5,
		},

		&dishIngridients{
			dishID:       4,
			ingridientId: 14,
			count:        200,
		},
		&dishIngridients{
			dishID:       4,
			ingridientId: 15,
			count:        4,
		},
		&dishIngridients{
			dishID:       4,
			ingridientId: 16,
			count:        1,
		},
		&dishIngridients{
			dishID:       4,
			ingridientId: 17,
			count:        3,
		},
		&dishIngridients{
			dishID:       4,
			ingridientId: 9,
			count:        4,
		},

		&dishIngridients{
			dishID:       5,
			ingridientId: 18,
			count:        200,
		},
		&dishIngridients{
			dishID:       5,
			ingridientId: 19,
			count:        2,
		},
		&dishIngridients{
			dishID:       5,
			ingridientId: 20,
			count:        20,
		},

		&dishIngridients{
			dishID:       6,
			ingridientId: 21,
			count:        1,
		},
		&dishIngridients{
			dishID:       6,
			ingridientId: 6,
			count:        2,
		},
		&dishIngridients{
			dishID:       6,
			ingridientId: 15,
			count:        1,
		},
		&dishIngridients{
			dishID:       6,
			ingridientId: 8,
			count:        5,
		},
		&dishIngridients{
			dishID:       6,
			ingridientId: 9,
			count:        1,
		},

		&dishIngridients{
			dishID:       7,
			ingridientId: 27,
			count:        200,
		},
		&dishIngridients{
			dishID:       7,
			ingridientId: 28,
			count:        5,
		},
		&dishIngridients{
			dishID:       7,
			ingridientId: 7,
			count:        150,
		},
		&dishIngridients{
			dishID:       7,
			ingridientId: 8,
			count:        5,
		},
		&dishIngridients{
			dishID:       7,
			ingridientId: 9,
			count:        1,
		},

		&dishIngridients{
			dishID:       8,
			ingridientId: 22,
			count:        750,
		},
		&dishIngridients{
			dishID:       8,
			ingridientId: 23,
			count:        2,
		},
		&dishIngridients{
			dishID:       8,
			ingridientId: 24,
			count:        3,
		},
		&dishIngridients{
			dishID:       8,
			ingridientId: 25,
			count:        1,
		},
		&dishIngridients{
			dishID:       8,
			ingridientId: 26,
			count:        1,
		},

		&dishIngridients{
			dishID:       9,
			ingridientId: 14,
			count:        200,
		},
		&dishIngridients{
			dishID:       9,
			ingridientId: 29,
			count:        100,
		},
		&dishIngridients{
			dishID:       9,
			ingridientId: 1,
			count:        2,
		},
		&dishIngridients{
			dishID:       9,
			ingridientId: 31,
			count:        50,
		},
		&dishIngridients{
			dishID:       9,
			ingridientId: 30,
			count:        100,
		},

		&dishIngridients{
			dishID:       10,
			ingridientId: 32,
			count:        300,
		},
		&dishIngridients{
			dishID:       10,
			ingridientId: 15,
			count:        3,
		},
		&dishIngridients{
			dishID:       10,
			ingridientId: 5,
			count:        50,
		},
		&dishIngridients{
			dishID:       10,
			ingridientId: 17,
			count:        2,
		},
		&dishIngridients{
			dishID:       10,
			ingridientId: 22,
			count:        50,
		},

		&dishIngridients{
			dishID:       11,
			ingridientId: 33,
			count:        50,
		},
		&dishIngridients{
			dishID:       11,
			ingridientId: 34,
			count:        100,
		},
		&dishIngridients{
			dishID:       11,
			ingridientId: 35,
			count:        20,
		},
		&dishIngridients{
			dishID:       11,
			ingridientId: 36,
			count:        1,
		},
		&dishIngridients{
			dishID:       11,
			ingridientId: 37,
			count:        10,
		},
		&dishIngridients{
			dishID:       11,
			ingridientId: 38,
			count:        100,
		},
	)
}
