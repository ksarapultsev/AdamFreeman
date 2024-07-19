package main

import "fmt"

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

// type Account struct {
// 	accountNumber int
// 	expenses      []Expense
// }

// func calcTotal(expenses []Expense) (total float64) {
// 	for _, item := range expenses {
// 		total += item.getCost(true)
// 	}
// 	return
// }

func main() {

	// account := Account{
	// 	accountNumber: 12345,
	// 	expenses:      []Expense{Product{"Kayak", "Watersports", 275}, Service{"Boat Cover", 12, 89.50}},
	// }

	// for _, expense := range account.expenses {
	// 	fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
	// }
	// fmt.Println("Total:", calcTotal(account.expenses))

	// product := Product{"Kayak", "Watersports", 275}

	// var expense Expense = product

	// product.price = 100

	// fmt.Println("Product field value:", product.price)
	// fmt.Println("Expense method result:", expense.getCost(false))

	// expenses := []Expense{
	// 	Product{"Kayak", "Watersports", 276},
	// 	Service{"Boat Cover", 12, 89.50},
	// }

	// for _, expense := range expenses {
	// 	fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
	// }
	// fmt.Println("Total:", calcTotal(expenses))
	fmt.Println("<--------------------------------------------->")
	expenses := []Expense{
		Service{"Boat Cover", 12, 89.50, []string{}},
		Service{"Paddle Protect", 12, 8, []string{}},
		&Product{"Kayak", "Watersports", 275},
	}

	for _, expense := range expenses {
		if s, ok := expense.(Service); ok {
			fmt.Println("Service1: ", s.description, "Price:", s.getCost(true))
		} else {
			fmt.Println("Expense1:", expense.getName(), "Cost:", expense.getCost(true))
		}

		// monthlyFee*float64(s.durationMonths))
	}

}
