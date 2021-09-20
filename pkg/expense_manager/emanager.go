package expense_manager

import (
	"fmt"
	"log"
	"projects/splitwise/pkg/models"
	"projects/splitwise/pkg/user_management"
)

type ExpenseManagerInterface interface {
	AddUser(user models.User)
	AddExpense(expense models.Expense)
	AddGroupExpense(expense models.GroupExpense)
	RunSettlements() []string
}

type ExpenseManager struct {
	um       user_management.UserManagementInterface
	Expenses []models.Expense
	Khata    map[string]map[string]float64
}

var expenseManager ExpenseManager

func NewExpenseManager(um user_management.UserManagementInterface) *ExpenseManager {
	expenseManager.um = um
	expenseManager.Khata = make(map[string]map[string]float64)
	return &expenseManager
}

func (em *ExpenseManager) AddUser(user models.User) {
	em.Khata[user.Name] = make(map[string]float64)
	em.um.AddUser(user)
}

func (em *ExpenseManager) AddExpense(expense models.Expense) {
	em.Expenses = append(em.Expenses, expense)

	paidBy, _ := em.um.GetUser(expense.PaidBy)

	if expense.Type == models.EXPENSE_TYPE_EQUAL {
		for _, name := range expense.Among {
			// Create Split
			amount := expense.Amount / float64(len(expense.Among))
			user, _ := em.um.GetUser(name)
			split := models.Split{
				User:   user,
				Amount: amount,
			}

			paidTo := split.User.Name

			// Adjust Khata of Payer
			khata := em.Khata[paidBy.Name]
			currentHisab := khata[paidTo]
			currentHisab = currentHisab + split.Amount
			khata[paidTo] = currentHisab

			// Adjust Khata of Receiver
			khata = em.Khata[paidTo]
			currentHisab = khata[paidBy.Name]
			khata[paidBy.Name] = currentHisab - split.Amount
		}
	}
	log.Printf("Current Khata: %v", em.Khata)
}

func (em *ExpenseManager) AddGroupExpense(expense models.GroupExpense) {
	users := em.um.GetGroupMembers(expense.GroupName)
	iexpense := models.Expense{
		Type:   "EQUAL",
		Amount: expense.Amount,
		PaidBy: expense.PaidBy,
		Among:  users,
	}
	em.AddExpense(iexpense)
}

func (em *ExpenseManager) RunSettlements() []string {
	var settlements []string

	for user1, kh := range em.Khata {
		for user2, amount := range kh {
			if user1 != user2 {
				if amount > 0 {
					balance := user2 + " owes " + user1 + ":" + fmt.Sprintf("%f", amount)
					settlements = append(settlements, balance)
					log.Printf("%s owes %s: %v", user2, user1, amount)
				} else if amount < 0 {
					//log.Printf("%s owes %s: %v", user1, user2, amount)
				} else {
					balance := user1 + " and " + user2 + " are all settled up"
					settlements = append(settlements, balance)
					log.Printf("%s owes %s are all settle up", user2, user1)
				}
			}
		}
	}

	return settlements
}
