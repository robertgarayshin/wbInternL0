package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"wbInternL0/models"
)

func allOrdersHandler(w http.ResponseWriter, r *http.Request) {
	var keys []string
	for key := range c.Orders {
		keys = append(keys, key)
	}
	type ViewData struct {
		Title    string
		OrderIds []string
	}
	data := ViewData{
		Title:    "Orders List",
		OrderIds: keys,
	}

	tmpl, _ := template.ParseFiles("templates/ordersList.html")
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}

func orderByIdHandler(w http.ResponseWriter, r *http.Request) {
	type ViewData struct {
		Title string
		Order models.Order
		Items []models.Item
	}
	vars := mux.Vars(r)
	id := vars["id"]
	data := ViewData{
		Title: fmt.Sprintf("Order %s", id),
		Order: c.Orders[id],
		Items: c.Orders[id].Items,
	}
	tmpl, _ := template.ParseFiles("templates/order.html")
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/index.html")
	type viewData struct {
		Title string
	}
	err := tmpl.Execute(w, viewData{Title: "WB Internship L0"})
	if err != nil {
		log.Println(err)
	}
}
