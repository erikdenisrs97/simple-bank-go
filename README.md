# Simple Bank

This application is a simple bank to explore net/http package and testing package of Golang.

The bank API supports these operations:
- deposit: deposit value on the specified account;
- withdraw: withdraw value from the specified account ;
- transfer: transfer value between accounts;

# How to Run

This application has build using Go v1.20

Go to bankapi folder and run: `go run main.go`

### Api Calls:

Transfer: `curl "http://localhost:8000/transfer?origin=1000&destiny=1001&amount=100"`

Deposit: `curl "http://localhost:8000/deposit?number=1001&amount=100"`

Withdraw: `curl "http://localhost:8000/withdraw?number=1001&amount=100"`