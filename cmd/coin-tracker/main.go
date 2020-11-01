package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/simsta6/coin-tracker/internal/rule"
	"github.com/simsta6/coin-tracker/pkg/coinlore"
)

const (
	rulesFilePath = "./rules.json"
	baseURL       = "https://api.coinlore.net/api"
)

var (
	wg    = sync.WaitGroup{}
	mutex = sync.Mutex{}
)

func main() {
	var usedRules []rule.Rule

	doneCh := waitForInterupt()

	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

continueProgram:
	for {
		select {
		case <-doneCh:
			cancelFunc()
			break continueProgram
		case <-ticker.C:
			execute(cancelCtx, usedRules)
		}
	}

	fmt.Println("exiting")
}

func waitForInterupt() <-chan struct{} {
	signalCh := make(chan os.Signal, 1)
	doneCh := make(chan struct{})

	signal.Notify(signalCh, os.Interrupt)

	go func() {
		signalMsg := <-signalCh
		fmt.Println(signalMsg)
		doneCh <- struct{}{}
	}()

	return doneCh
}

//Executes main program
func execute(ctx context.Context, usedRules []rule.Rule) (err error) {
	currencyMap := make(map[string]coinlore.Currency)
	client := coinlore.NewClient(baseURL)

	rules, err := rule.ReadRulesJSON(rulesFilePath)
	if err != nil {
		return err
	}

	n := len(rules)
	wg.Add(n)
	for i := range rules {
		go operateRule(ctx, &rules[i], &currencyMap, client)
	}
	wg.Wait()

	//If no rule was used in this iteration it's no need to rewrite json file
	if rulesWereUsed(rules) {
		rules, usedRules = filter(rules, usedRules)
		err = rule.WriteJSON(rules, rulesFilePath)
		if err != nil {
			return err
		}
	}
	return err
}

//Every calculation for rule
func operateRule(ctx context.Context, r *rule.Rule, data *map[string]coinlore.Currency, client *coinlore.Client) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			if err := getCryptoDataMap(data, r.CryptoID, client); err != nil {
				fmt.Println(err)
				return
			}

			currencyData := (*data)[r.CryptoID]

			if err := r.Compare(currencyData.PriceUSD); err != nil {
				fmt.Println(err)
			}

			if r.Used {
				fmt.Println(r.ToString())
			}
			return
		}
	}
}

//Check if rules were used
func rulesWereUsed(rules []rule.Rule) bool {
	for _, v := range rules {
		if v.Used {
			return true
		}
	}
	return false
}

//Sends a call to server, gets data about crypto currency and puts it in a map by cryptoID
func getCryptoDataMap(data *map[string]coinlore.Currency, cryptoID string, client *coinlore.Client) (err error) {
	rootCtx := context.Background()
	childCtx, cancelReq := context.WithTimeout(rootCtx, time.Second*3)

	rawCurr, err := client.GetCurrency(childCtx, cryptoID)
	if err != nil {
		cancelReq()
		return err
	}
	cancelReq()

	curr, err := rawCurr.ParseCrypto()
	if err != nil {
		return err
	}

	mutex.Lock()
	(*data)[cryptoID] = curr
	mutex.Unlock()

	return err
}
