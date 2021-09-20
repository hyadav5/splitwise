package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"projects/splitwise/pkg/api"
)

type Application struct {
	api api.ApiInterface
}

var App Application

func NewApplicationHandler() {
	App.api = api.NewApiInterface()
}

func main() {
	log.Printf("Running Splitwise Backend")

	NewApplicationHandler()

	App.api.PopulateTestData()

	r := mux.NewRouter()
	r.HandleFunc("/adduser", App.api.AddUser)
	r.HandleFunc("/addgroup", App.api.AddGroup)
	r.HandleFunc("/users", App.api.GetUsers)

	r.HandleFunc("/addexpense", App.api.AddExpense)
	r.HandleFunc("/addgroupexpense", App.api.AddGroupExpense)
	r.HandleFunc("/runsettlements", App.api.RunSettlements)

	log.Printf("Serving Splitwise Backend at 8080 port")
	log.Fatal(http.ListenAndServe(":8080", r))
}
