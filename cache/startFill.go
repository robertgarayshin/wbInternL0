package cache

import (
	"database/sql"
	"wbInternL0/repository/read"
)

// Стартовое заполнение кэша
func (c *Cache) StartFill(db *sql.DB) {
	orders := read.ReadAll(db)
	for i := range orders {
		c.Set(orders[i].OrderUid, orders[i])
	}
}
