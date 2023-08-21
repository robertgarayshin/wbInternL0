package cache

import (
	"wbInternL0/models"
)

func (c *Cache) Set(key string, order models.Order) {
	// Вставка элемента в кэш
	c.Lock()

	defer c.Unlock()

	c.Orders[key] = order

}
