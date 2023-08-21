package initializr

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"wbInternL0/config"
)

func DbConnectionInit() (*sql.DB, error) {
	// Инициализируем подключение к БД
	db, err := sql.Open("postgres", config.ConnStr)
	if err != nil {
		log.Println(err)
		return nil, err
	} else {
		log.Printf("Connection to database opened successfully")
	}
	return db, nil
}
