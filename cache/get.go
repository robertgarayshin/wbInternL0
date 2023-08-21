package cache

import (
	"wbInternL0/models"
)

func (c *Cache) Get(key string) (models.Order, bool) {
	// Получение элемента из кэша по ключу
	c.RLock()

	defer c.RUnlock()

	order, found := c.Orders[key]

	// ключ не найден
	if !found {
		return order, false
	}

	return order, true
}
