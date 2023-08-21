package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type msg string

func (m msg) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprint(resp, m)
	if err != nil {
		return
	}
}

func Serve(ctx context.Context) {
	server := http.Server{Addr: ":8080"}
	router := mux.NewRouter()
	router.HandleFunc("/orders", allOrdersHandler)
	router.HandleFunc("/orders/{id}", orderByIdHandler)
	router.HandleFunc("/", mainHandler)
	http.Handle("/", router)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			return
		}
	}()
	for {
		select {
		case <-ctx.Done():
			log.Println("Shutting down server")
			err := server.Shutdown(ctx)
			if err != nil {
				panic(err)
			}
			return
		}
	}
}
