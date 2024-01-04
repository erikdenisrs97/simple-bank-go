package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/erikdenisrs97/bank"
)

var accounts = map[float64]*CustomAccount{}

func main() {
	accounts[1000] = &CustomAccount{
		Account: &bank.Account{
			Customer: bank.Customer{
				Name:    "Erik",
				Address: "Curitiba, Brazil",
				Phone:   "(41) 1111-1111",
			},
			Number: 1000,
		},
	}

	accounts[1001] = &CustomAccount{
		Account: &bank.Account{
			Customer: bank.Customer{
				Name:    "Camila",
				Address: "Fortaleza, Brazil",
				Phone:   "(81) 2222-2222",
			},
			Number: 1001,
		},
	}

	http.HandleFunc("/statement", statement)
	http.HandleFunc("/deposit", deposit)
	http.HandleFunc("/withdraw", withdraw)
	http.HandleFunc("/transfer", transfer)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func statement(w http.ResponseWriter, req *http.Request) {
	accountNumber := req.URL.Query().Get("number")

	if accountNumber == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(accountNumber, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			json.NewEncoder(w).Encode(bank.Statement(account))
		}
	}
}

func deposit(w http.ResponseWriter, req *http.Request) {
	accountNumber := req.URL.Query().Get("number")
	amount := req.URL.Query().Get("amount")

	if accountNumber == "" {
		fmt.Fprintf(w, "Account number is missing")
		return
	}

	if number, err := strconv.ParseFloat(accountNumber, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amount, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			err := account.Deposit(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprintf(w, account.Statement())
			}
		}
	}
}

func withdraw(w http.ResponseWriter, req *http.Request) {
	accountNumber := req.URL.Query().Get("number")
	amount := req.URL.Query().Get("amount")

	if accountNumber == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(accountNumber, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amount, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			err := account.Withdraw(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprintf(w, account.Statement())
			}
		}
	}
}

func transfer(w http.ResponseWriter, req *http.Request) {
	sourceAccountNumber := req.URL.Query().Get("origin")
	targetAccountNumber := req.URL.Query().Get("destiny")
	amount := req.URL.Query().Get("amount")

	if sourceAccountNumber == "" || targetAccountNumber == "" {
		fmt.Fprintf(w, "Some account number is missing!")
	}

	if sourceNumber, err := strconv.ParseFloat(sourceAccountNumber, 64); err != nil {
		fmt.Fprintf(w, "Invalid source account number!")
	} else if targetNumber, err := strconv.ParseFloat(targetAccountNumber, 64); err != nil {
		fmt.Fprintf(w, "Invalid target account number!")
	} else if amount, err := strconv.ParseFloat(amount, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		sourceAccount, ok := accounts[sourceNumber]
		if !ok {
			fmt.Fprintf(w, "Source account with number %v can't be found!", sourceNumber)
		}
		targetAccount, ok := accounts[targetNumber]
		if !ok {
			fmt.Fprintf(w, "Target account with number %v can't be found!", targetNumber)
		} else {
			err := sourceAccount.Transfer(amount, targetAccount.Account)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprintf(w, sourceAccount.Statement())
			}
		}

	}

}

type CustomAccount struct {
	*bank.Account
}

func (c *CustomAccount) Statement() string {
	json, err := json.Marshal(c)
	if err != nil {
		return err.Error()
	}

	return string(json)
}
