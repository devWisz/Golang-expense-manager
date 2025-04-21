package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Transaction struct {
	Type        string  `json:"type"`
	Amount      float64 `json:"amount"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	Date        string  `json:"date"`
}

var transactions []Transaction
var fileName = "budget.json"

func main() {
	loadData()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n📌 Choose an option:")
		fmt.Println("1. Add Income")
		fmt.Println("2. Add Expense")
		fmt.Println("3. View Balance")
		fmt.Println("4. List All Transactions")
		fmt.Println("5. Monthly Summary")
		fmt.Println("6. Exit")
		fmt.Print("👉 ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			addTransaction("income", reader)
		case "2":
			addTransaction("expense", reader)
		case "3":
			showBalance()
		case "4":
			listTransactions()
		case "5":
			showMonthlySummary()
		case "6":
			saveData()
			fmt.Println(" Exiting. Your data is saved.")
			return
		default:
			fmt.Println("Invalid option.")
		}
	}
}

func addTransaction(tType string, reader *bufio.Reader) {
	fmt.Print("Enter amount: ")
	amountStr, _ := reader.ReadString('\n')
	amountStr = strings.TrimSpace(amountStr)
	amount, _ := strconv.ParseFloat(amountStr, 64)

	fmt.Print("Enter category (e.g., food, reward, travel): ")
	category, _ := reader.ReadString('\n')
	category = strings.TrimSpace(category)

	fmt.Print("Enter description: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	t := Transaction{
		Type:        tType,
		Amount:      amount,
		Category:    category,
		Description: description,
		Date:        time.Now().Format("2006-01-02"),
	}

	transactions = append(transactions, t)
	saveData()

	fmt.Println("Transaction recorded.")
}

func showBalance() {
	income := 0.0
	expense := 0.0
	for _, t := range transactions {
		if t.Type == "income" {
			income += t.Amount
		} else if t.Type == "expense" {
			expense += t.Amount
		}
	}
	fmt.Printf("\n Income: %.2f\n Expense: %.2f\n Balance: %.2f\n", income, expense, income-expense)
}

func listTransactions() {
	fmt.Println("\n All Transactions:")
	for _, t := range transactions {
		fmt.Printf("[%s] %s: ₹%.2f (%s) - %s\n", t.Date, strings.Title(t.Type), t.Amount, t.Category, t.Description)
	}
}

func showMonthlySummary() {
	currentMonth := time.Now().Format("2006-01")
	income := 0.0
	expense := 0.0

	fmt.Println("\n Monthly Summary:")
	for _, t := range transactions {
		if strings.HasPrefix(t.Date, currentMonth) {
			if t.Type == "income" {
				income += t.Amount
			} else {
				expense += t.Amount
			}
		}
	}
	fmt.Printf("This month:  Income: ₹%.2f,  Expense: ₹%.2f, Balance: ₹%.2f\n", income, expense, income-expense)
}

func saveData() {
	data, _ := json.MarshalIndent(transactions, "", "  ")
	_ = os.WriteFile(fileName, data, 0644)
}

func loadData() {
	data, err := os.ReadFile(fileName)
	if err == nil {
		_ = json.Unmarshal(data, &transactions)
	}
}
