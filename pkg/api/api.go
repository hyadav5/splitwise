package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"projects/splitwise/pkg/expense_manager"
	"projects/splitwise/pkg/models"
	"projects/splitwise/pkg/user_management"
)

type ApiInterface interface {
	PopulateTestData()
	AddUser(w http.ResponseWriter, req *http.Request)
	AddGroup(w http.ResponseWriter, req *http.Request)
	GetGroupMembers(name string) []string
	GetUsers(w http.ResponseWriter, req *http.Request)
	AddExpense(w http.ResponseWriter, req *http.Request)
	AddGroupExpense(w http.ResponseWriter, req *http.Request)
	RunSettlements(w http.ResponseWriter, req *http.Request)
}

type Api struct {
	em expense_manager.ExpenseManagerInterface
	um user_management.UserManagementInterface
}

var api Api

func NewApiInterface() *Api {
	api.um = user_management.NewUserManagement()
	api.em = expense_manager.NewExpenseManager(api.um)
	return &api
}

func (api *Api) AddUser(w http.ResponseWriter, req *http.Request) {
	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	api.em.AddUser(user)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Added user: %s\n", user.Name)
}

func (api *Api) AddGroup(w http.ResponseWriter, req *http.Request) {
	var addGroupRequest models.AddGroupRequest
	err := json.NewDecoder(req.Body).Decode(&addGroupRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	api.um.AddGroup(addGroupRequest.Name, addGroupRequest.Users)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Added Group: %s\n", addGroupRequest.Name)
}

func (api *Api) GetUsers(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Users: %v\n", api.um.GetUsers())
}

func (api *Api) AddExpense(w http.ResponseWriter, req *http.Request) {
	var expense models.Expense
	err := json.NewDecoder(req.Body).Decode(&expense)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	api.em.AddExpense(expense)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Added Expense")
	api.em.AddExpense(expense)
}

func (api *Api) AddGroupExpense(w http.ResponseWriter, req *http.Request) {
	var groupExpense models.GroupExpense
	err := json.NewDecoder(req.Body).Decode(&groupExpense)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	api.em.AddGroupExpense(groupExpense)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Added Group Expense")
}

func (api *Api) GetGroupMembers(name string) []string {
	return api.um.GetGroupMembers(name)
}

func (api *Api) RunSettlements(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Users: %v\n", api.em.RunSettlements())
}

func (api *Api) PopulateTestData() {
	u1 := models.User{
		Id:      1,
		Name:    "hemant",
		Contact: "700000004",
		Email:   "hemant.xxxx@outlook.com",
	}

	u2 := models.User{
		Id:      1,
		Name:    "kd",
		Contact: "700000004",
		Email:   "kd.xxxx@outlook.com",
	}

	u3 := models.User{
		Id:      1,
		Name:    "pandey",
		Contact: "700000004",
		Email:   "pandey.xxxx@outlook.com",
	}

	u4 := models.User{
		Id:      1,
		Name:    "shyam",
		Contact: "700000004",
		Email:   "shyam.xxxx@outlook.com",
	}

	var userList []string
	userList = append(userList, u1.Name)
	userList = append(userList, u2.Name)
	userList = append(userList, u3.Name)
	userList = append(userList, u4.Name)

	api.em.AddUser(u1)
	api.em.AddUser(u2)
	api.em.AddUser(u3)
	api.em.AddUser(u4)
	api.um.AddGroup("flatmates", userList)

	expense := models.Expense{
		Type:   "EQUAL",
		Amount: 1000,
		PaidBy: u1.Name,
		Among:  userList,
	}
	api.em.AddExpense(expense)

	expense = models.Expense{
		Type:   "EQUAL",
		Amount: 1000,
		PaidBy: u2.Name,
		Among:  userList,
	}
	api.em.AddExpense(expense)

	gExpense := models.GroupExpense{
		Amount:    2000,
		PaidBy:    u1.Name,
		GroupName: "flatmates",
	}
	api.em.AddGroupExpense(gExpense)

	api.em.RunSettlements()
}
