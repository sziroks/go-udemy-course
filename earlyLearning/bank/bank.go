package main

import (
	"fmt"
	"strconv"
)

func main() {
	var user_balance float64

	fmt.Println("Welcome to the bank!")
	fmt.Println("Please select an option:")

out:
	for {
		showPossibleOptions()

		var user_input string
		fmt.Scan(&user_input)
		choice, err := strconv.Atoi(user_input)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}
		switch choice {
		case 1:
			checkBalance(user_balance)
		case 2:
			user_balance = deposit(user_balance)
		case 3:
			user_balance = withdraw(user_balance)
		case 4:
			break out
		default:
			fmt.Println("Invalid input")
		}
	}

}

func showPossibleOptions() {
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Println("4. Exit")
	fmt.Print("Your choice: ")
}

func checkBalance(balance float64) {
	fmt.Println("Your balance is: ", balance)
}

func deposit(balance float64) float64 {
	var amount float64
	fmt.Print("Enter the amount you want to deposit: ")
	fmt.Scan(&amount)

	if amount < 0 {
		fmt.Println("Invalid amount")
		return balance
	}
	fmt.Printf("Deposited: %v, your new balance is: %v\n", amount, balance+amount)
	return balance + amount
}

func withdraw(balance float64) float64 {
	var amount float64
	fmt.Print("Enter the amount you want to withdraw: ")
	fmt.Scan(&amount)
	if amount < 0 {
		fmt.Println("Invalid amount")
		return balance
	}
	if amount > balance {
		fmt.Println("Insufficient funds")
		return balance
	}
	fmt.Printf("Withdrew: %v, your new balance is: %v\n", amount, balance-amount)
	return balance - amount
}
