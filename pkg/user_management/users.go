package user_management

import (
	"errors"
	"log"
	"projects/splitwise/pkg/models"
)

type UserManagementInterface interface {
	AddUser(user models.User)
	GetUser(name string) (models.User, error)
	GetUsers() []models.User
	AddGroup(name string, users []string)
	GetGroupMembers(name string) []string
}

type UserManagement struct {
	userMap  map[string]models.User
	groupMap map[string][]string
}

var userManagement UserManagement

func NewUserManagement() *UserManagement {
	userManagement.userMap = make(map[string]models.User)
	userManagement.groupMap = make(map[string][]string)
	return &userManagement
}

func (um *UserManagement) AddUser(user models.User) {
	um.userMap[user.Name] = user
	log.Printf("Added User: %s", user.Name)
}

func (um *UserManagement) GetUser(name string) (user models.User, err error) {
	log.Printf("Getting User: %s", name)
	elem, ok := um.userMap[name]
	if ok {
		return elem, nil
	}
	return models.User{}, errors.New("No User Found")
}

func (um *UserManagement) GetUsers() (user []models.User) {
	log.Printf("Getting All Users")
	var userList []models.User

	for _, value := range um.userMap {
		userList = append(userList, value)
	}

	return userList
}

func (um *UserManagement) AddGroup(name string, users []string) {
	um.groupMap[name] = users
	log.Printf("Added Group: %s", name)
}

func (um *UserManagement) GetGroupMembers(name string) []string {
	log.Printf("Getting Group Members: %s", name)
	return um.groupMap[name]
}
