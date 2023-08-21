package cache

import (
	"wbInternL0/models"
)

func (c *Cache) Set(key string, order models.Order) {

	c.Lock()

	defer c.Unlock()

	c.Orders[key] = order

}
