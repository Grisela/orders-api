package handler

import (
	"fmt"
	"net/http"
)

type Order struct{}

func (o *Order) Create(w http.ResponseWriter, r *http.Request){
	fmt.Println(("Create an Order"))
}

func (o *Order) List(w http.ResponseWriter, r *http.Request){
	fmt.Println(("List Orders"))
}

func (o *Order) GetById(w http.ResponseWriter, r *http.Request){
	fmt.Println(("Get an order by id"))
}

func (o *Order) UpdateById(w http.ResponseWriter, r *http.Request){
	fmt.Println(("Update an order by id"))
}

func (o *Order) DeleteById(w http.ResponseWriter, r *http.Request){
	fmt.Println(("Delete an order by id"))
}