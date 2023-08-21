package cache

import (
	"sync"
	"wbInternL0/models"
)

type Cache struct {
	// Структура кэша
	sync.RWMutex
	Orders map[string]models.Order
}
