package initializr

import (
	"wbInternL0/cache"
	"wbInternL0/models"
)

func InitCache() *cache.Cache {
	// инициализируем карту(map) в паре ключ(string)/значение(Order)
	orders := make(map[string]models.Order)

	c := cache.Cache{
		Orders: orders,
	}
	return &c

}
