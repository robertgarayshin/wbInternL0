package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"wbInternL0/cache"
	"wbInternL0/initializr"
	"wbInternL0/jetStream"
)

var c *cache.Cache

func main() {
	defer log.Println("Shutting down completed")
	log.Println("Starting")

	log.Println("Opening database connection")
	db, err := initializr.DbConnectionInit()
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
		log.Println("Connection to database closed successfully")
	}(db)

	c = initializr.InitCache()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.StartFill(db)
	}()
	wg.Wait()
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sigChannel := make(chan os.Signal, 1)
		signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM)
		<-sigChannel
		close(sigChannel)
		cancel()
	}()
	go jetStream.Consumer(ctx, db, c)
	Serve(ctx)

}
