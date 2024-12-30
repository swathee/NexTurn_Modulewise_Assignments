//Exercise 2: Bank Transaction System

package main

import (
	"errors"
	"fmt"
)

// Account struct to store account details
type Account struct {
	ID              int
	Name            string
	Balance         float64
	TransactionHist []string
}

// Global slice to store all accounts
var accounts []Account

// Deposit function to deposit money into an account
func (acc *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}
	acc.Balance += amount
	acc.TransactionHist = append(acc.TransactionHist, fmt.Sprintf("Deposited: %.2f", amount))
	return nil
}

// Withdraw function to withdraw money from an account
func (acc *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be greater than zero")
	}
	if acc.Balance < amount {
		return errors.New("insufficient balance")
	}
	acc.Balance -= amount
	acc.TransactionHist = append(acc.TransactionHist, fmt.Sprintf("Withdrew: %.2f", amount))
	return nil
}

// ViewTransactionHistory function to display transaction history
func (acc *Account) ViewTransactionHistory() {
	fmt.Println("Transaction History:")
	if len(acc.TransactionHist) == 0 {
		fmt.Println("No transactions found.")
		return
	}
	for i, transaction := range acc.TransactionHist {
		fmt.Printf("%d. %s\n", i+1, transaction)
	}
}

// FindAccountByID finds an account by its ID
func FindAccountByID(id int) (*Account, error) {
	for i := range accounts {
		if accounts[i].ID == id {
			return &accounts[i], nil
		}
	}
	return nil, errors.New("account not found")
}

// Menu function for account management
func menu() {
	fmt.Println("\n--- Bank Transaction System ---")
	fmt.Println("1. Add Account")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Println("4. View Balance")
	fmt.Println("5. Transaction History")
	fmt.Println("6. Exit")
	fmt.Print("Enter your choice: ")
}

func main() {
	// Pre-populated accounts
	accounts = append(accounts, Account{ID: 1, Name: "Ken", Balance: 4000})
	accounts = append(accounts, Account{ID: 2, Name: "Kevin", Balance: 8000})
	accounts = append(accounts, Account{ID: 2, Name: "Yogesh", Balance: 7000})
	accounts = append(accounts, Account{ID: 2, Name: "Nevin", Balance: 2000})

	const (
		AddAccountOption         = 1
		DepositOption            = 2
		WithdrawOption           = 3
		ViewBalanceOption        = 4
		TransactionHistoryOption = 5
		ExitOption               = 6
	)

	for {
		menu()
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case AddAccountOption:
			fmt.Print("Enter Account ID: ")
			var id int
			fmt.Scan(&id)

			fmt.Print("Enter Account Name: ")
			var name string
			fmt.Scan(&name)

			fmt.Print("Enter Initial Balance: ")
			var balance float64
			fmt.Scan(&balance)

			accounts = append(accounts, Account{ID: id, Name: name, Balance: balance})
			fmt.Println("Account added successfully.")

		case DepositOption:
			fmt.Print("Enter Account ID: ")
			var id int
			fmt.Scan(&id)

			account, err := FindAccountByID(id)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			fmt.Print("Enter amount to deposit: ")
			var amount float64
			fmt.Scan(&amount)

			if err := account.Deposit(amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Deposit successful.")
			}

		case WithdrawOption:
			fmt.Print("Enter Account ID: ")
			var id int
			fmt.Scan(&id)

			account, err := FindAccountByID(id)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			fmt.Print("Enter amount to withdraw: ")
			var amount float64
			fmt.Scan(&amount)

			if err := account.Withdraw(amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Withdrawal successful.")
			}

		case ViewBalanceOption:
			fmt.Print("Enter Account ID: ")
			var id int
			fmt.Scan(&id)

			account, err := FindAccountByID(id)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			fmt.Printf("Current balance for %s: %.2f\n", account.Name, account.Balance)

		case TransactionHistoryOption:
			fmt.Print("Enter Account ID: ")
			var id int
			fmt.Scan(&id)

			account, err := FindAccountByID(id)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			account.ViewTransactionHistory()

		case ExitOption:
			fmt.Println("Exiting the program. Thank you!")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
