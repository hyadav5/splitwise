package models

type User struct {
	Id      int
	Name    string
	Contact string
	Email   string
}

type Split struct {
	User   User
	Amount float64
}

type ExpenseType int

const (
	EXPENSE_TYPE_EQUAL       = "EQUAL"
	EXPENSE_TYPE_WITH_AMOUNT = "WITH_AMOUNT"
)

type Expense struct {
	ID     int64
	Type   string
	Amount float64
	PaidBy string
	Among  []string
}

// Group expense are by default equally distributed among group members.
type GroupExpense struct {
	ID        int64
	Amount    float64
	PaidBy    string
	GroupName string
}

// ==============================================================================
// External models to be shared to shared to external world or Frontend

type AddGroupRequest struct {
	Name  string
	Users []string
}

// ==============================================================================
