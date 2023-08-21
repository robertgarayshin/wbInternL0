package cache

import (
	"sync"
	"wbInternL0/models"
)

type Cache struct {
	sync.RWMutex
	Orders map[string]models.Order
}
