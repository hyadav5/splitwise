package expense_manager

import (
	"github.com/golang/mock/gomock"
	"log"
	"os"
	"projects/splitwise/pkg/models"
	"projects/splitwise/pkg/user_management"
	"testing"
)

func TestMain(m *testing.M) {
	log.SetOutput(os.Stdout)
	var um user_management.UserManagementInterface
	um = user_management.NewUserManagement()
	_ = NewExpenseManager(um)
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func tearDown() {}

func Test_AddExpense(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	u1 := models.User{
		Id:      1,
		Name:    "hemant",
		Contact: "7411382773",
		Email:   "hemant.yadav@outlook.com",
	}

	u2 := models.User{
		Id:      1,
		Name:    "kd",
		Contact: "7411382773",
		Email:   "hemant.yadav@outlook.com",
	}

	u3 := models.User{
		Id:      1,
		Name:    "pandey",
		Contact: "7411382773",
		Email:   "hemant.yadav@outlook.com",
	}

	u4 := models.User{
		Id:      1,
		Name:    "shyam",
		Contact: "7411382773",
		Email:   "hemant.yadav@outlook.com",
	}

	expenseManager.AddUser(u1)
	expenseManager.AddUser(u2)
	expenseManager.AddUser(u3)
	expenseManager.AddUser(u4)

	var userList []string
	userList = append(userList, u1.Name)
	userList = append(userList, u2.Name)
	userList = append(userList, u3.Name)
	userList = append(userList, u4.Name)

	expense := models.Expense{
		Type:   "EQUAL",
		Amount: 1000,
		PaidBy: u1.Name,
		Among:  userList,
	}

	expenseManager.AddExpense(expense)
}

func Test_AddGroupExpense(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	u1 := models.User{
		Id:      1,
		Name:    "hemant",
		Contact: "7411382773",
		Email:   "hemant.yadav@outlook.com",
	}

	u2 := models.User{
		Id:      1,
		Name:    "kd",
		Contact: "7411382773",
		Email:   "hemant.yadav@outlook.com",
	}

	u3 := models.User{
		Id:      1,
		Name:    "pandey",
		Contact: "7411382773",
		Email:   "hemant.yadav@outlook.com",
	}

	u4 := models.User{
		Id:      1,
		Name:    "shyam",
		Contact: "7411382773",
		Email:   "hemant.yadav@outlook.com",
	}

	expenseManager.AddUser(u1)
	expenseManager.AddUser(u2)
	expenseManager.AddUser(u3)
	expenseManager.AddUser(u4)

	var userList []string
	userList = append(userList, u1.Name)
	userList = append(userList, u2.Name)
	userList = append(userList, u3.Name)
	userList = append(userList, u4.Name)

	expenseManager.um.AddGroup("flatmates", userList)

	gExpense := models.GroupExpense{
		Amount:    2000,
		PaidBy:    u1.Name,
		GroupName: "flatmates",
	}

	expenseManager.AddGroupExpense(gExpense)
}
