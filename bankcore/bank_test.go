package bank

import "testing"

func TestAccount(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "Erik",
			Address: "Curitiba, Brazil",
			Phone:   "(41) 1111-1111",
		},
		Number:  1000,
		Balance: 0,
	}

	if account.Name == "" {
		t.Error("can't create an Account object")
	}
}

func TestDeposit(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "Erik",
			Address: "Curitiba, Brazil",
			Phone:   "(41) 1111-1111",
		},
		Number:  1000,
		Balance: 0,
	}

	account.Deposit(10)

	if account.Balance != 10 {
		t.Error("balance is not being update after a deposit")
	}
}

func TestDepositInvalid(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "Erik",
			Address: "Curitiba, Brazil",
			Phone:   "(41) 1111-1111",
		},
		Number:  1000,
		Balance: 0,
	}

	if err := account.Deposit(-10); err == nil {
		t.Error("only positive numbers should be allowed to deposit")
	}
}

func TestWithdraw(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "Erik",
			Address: "Curitiba, Brazil",
			Phone:   "(41) 1111-1111",
		},
		Number:  1000,
		Balance: 0,
	}

	account.Deposit(10)
	account.Withdraw(10)

	if account.Balance != 0 {
		t.Error("balance is not being updated after withdraw")
	}
}

func TestStatement(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "Erik",
			Address: "Curitiba, Brazil",
			Phone:   "(41) 1111-1111",
		},
		Number:  1000,
		Balance: 0,
	}

	account.Deposit(50)
	statement := account.Statement()

	if statement != "1000 - Erik - 50" {
		t.Error("statement doesnt have the proper format")
	}
}

func TestTransfer(t *testing.T) {
	var accounts = map[float64]*Account{}

	accounts[1000] = &Account{
		Customer: Customer{
			Name:    "Erik",
			Address: "Curitiba, Brazil",
			Phone:   "(41) 1111-1111",
		},
		Number:  1000,
		Balance: 50,
	}

	accounts[1001] = &Account{
		Customer: Customer{
			Name:    "Camila",
			Address: "Fortaleza, Brazil",
			Phone:   "(81) 2222-2222",
		},
		Number:  1001,
		Balance: 0,
	}

	accounts[1000].Transfer(50, accounts[1001])

	if accounts[1000].Balance != 0 {
		t.Error("balance is not being updated after transfer.")
	}

	if accounts[1001].Balance != 50 {
		t.Error("the transfer is not being completed.")
	}

}

func TestTransferWithBalanceZero(t *testing.T) {
	var accounts = map[float64]*Account{}

	accounts[1000] = &Account{
		Customer: Customer{
			Name:    "Erik",
			Address: "Curitiba, Brazil",
			Phone:   "(41) 1111-1111",
		},
		Number:  1000,
		Balance: 0,
	}

	accounts[1001] = &Account{
		Customer: Customer{
			Name:    "Camila",
			Address: "Fortaleza, Brazil",
			Phone:   "(81) 2222-2222",
		},
		Number:  1001,
		Balance: 0,
	}

	if err := accounts[1000].Transfer(50, accounts[1001]); err == nil {
		t.Error("balance should be greater than zero to transfer.")
	}

}

func TestTransferInvalid(t *testing.T) {
	var accounts = map[float64]*Account{}

	accounts[1000] = &Account{
		Customer: Customer{
			Name:    "Erik",
			Address: "Curitiba, Brazil",
			Phone:   "(41) 1111-1111",
		},
		Number:  1000,
		Balance: 100,
	}

	accounts[1001] = &Account{
		Customer: Customer{
			Name:    "Camila",
			Address: "Fortaleza, Brazil",
			Phone:   "(81) 2222-2222",
		},
		Number:  1001,
		Balance: 0,
	}

	if err := accounts[1000].Transfer(-50, accounts[1001]); err == nil {
		t.Error("only positive numbers should be allowed to transfer.")
	}

}
