package user_management

import (
	"github.com/golang/mock/gomock"
	"log"
	"os"
	"projects/splitwise/pkg/models"
	"testing"
)

func TestMain(m *testing.M) {
	log.SetOutput(os.Stdout)
	_ = NewUserManagement()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func tearDown() {}

func Test_AddUser_GetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := models.User{
		Id:      1,
		Name:    "hemant",
		Contact: "7411382773",
		Email:   "hemant.yadav@outlook.com",
	}

	userManagement.AddUser(user)
	user_expected, _ := userManagement.GetUser(user.Name)
	if user_expected != user {
		t.Errorf("Failed to get user. Test Failed.")
	}
}

func Test_GetUser_Failed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := models.User{
		Id:      1,
		Name:    "hemant",
		Contact: "7411382773",
		Email:   "hemant.yadav@outlook.com",
	}

	delete(userManagement.userMap, user.Name)
	user_expected, _ := userManagement.GetUser(user.Name)
	if user_expected == user { // User should not be their.
		t.Errorf("Failed to get user. Test Failed.")
	}
}

func Test_GetUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := models.User{
		Id:      1,
		Name:    "hemant",
		Contact: "7411382773",
		Email:   "hemant.yadav@outlook.com",
	}

	userManagement.AddUser(user)
	users := userManagement.GetUsers()
	if users[0] != user {
		t.Errorf("Failed to get users. Test Failed.")
	}
}

func Test_AddGroup_GetGroupMembers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := models.User{
		Id:      1,
		Name:    "hemant",
		Contact: "7411382773",
		Email:   "hemant.yadav@outlook.com",
	}
	var groupName = "GroupName"

	var expectedMember []string
	expectedMember = append(expectedMember, user.Name)

	userManagement.userMap = make(map[string]models.User)
	userManagement.AddUser(user)
	var userList []string
	userList = append(userList, user.Name)
	userManagement.AddGroup(groupName, userList)
	expectedMember = userManagement.GetGroupMembers(groupName)
	if expectedMember[0] != user.Name {
		t.Errorf("Failed to add group and get group members. Test Failed.")
	}
}
