# Coin-Tracker

Coin-Tracker is for tracking CryptoCurrency prices from [coinlore.com](https://www.coinlore.com/)\
All rules which Coin-Tracker tracks is located in main folder [rules.json](rules.json)
 
## To run program

 	go run ./cmd/coin-tracker

## To run tests

	go test ./cmd/coin-tracker
	go test ./internal/rule
	go test ./pkg/coinlore
