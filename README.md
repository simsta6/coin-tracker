# Coin-tracker

Coin-tracker is for tracking CryptoCurrency prices from [coinlore.com](https://www.coinlore.com/)\
All rules which coin-tracker tracks is located in main folder [rules.json](rules.json)
 
## To run program

 	go run ./cmd/coin-tracker

## To run tests

	go test ./cmd/coin-tracker
	go test ./internal/rule
	go test ./pkg/coinlore
