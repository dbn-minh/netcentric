package mockdb

import "mockdb/model"

var Products = []model.Product{
	{ID: 1, Name: "Laptop", Price: 1200.0, CategoryID: 1},
	{ID: 2, Name: "T-Shirt", Price: 20.0, CategoryID: 2},
}
