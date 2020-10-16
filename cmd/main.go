package main

import (
	"fmt"
	"time"

	currencyTracker "github.com/simsta6/junior-task/internal/currencytracker"
	api "github.com/simsta6/junior-task/pkg/apicalls"
)

const (
	baseURL          = "https://api.coinlore.net/api/ticker/?id=90"
	rulesFilePath    = "C:\\Users\\simas\\Desktop\\junior-task\\rules.json"
	currencyFilePath = "C:\\Users\\simas\\Desktop\\junior-task\\currency.json"
)

func main() {
	var usedRules []currencyTracker.RawRule
	for {
		rules, err := currencyTracker.ReadRulesJSON(rulesFilePath)

		if err != nil {
			fmt.Println(err)
		}

		err = api.DownloadFile(baseURL, currencyFilePath)
		if err != nil {
			fmt.Println(err)
		}
		rawCurr, err := currencyTracker.ReadCryptocurrencyJSON(currencyFilePath)
		if err != nil {
			fmt.Println(err)
		}

		curr, err := rawCurr[0].ParseCrypto()
		if err != nil {
			fmt.Println(err)
		}

		currencyTracker.InitMessages(curr)

		for i := range rules {
			answer, _ := rules[i].Compare(curr.PriceUSD)
			fmt.Println(answer)
		}

		for i, rule := range rules {
			if rule.Used {
				n := len(rules) - 1
				usedRules = append(usedRules, rule)
				rules[i] = rules[n]
				rules[n] = currencyTracker.RawRule{}
				rules = rules[:n]
			}
		}

		fmt.Println(rules)

		fmt.Println(curr)

		currencyTracker.WriteRulesJSON(rules, rulesFilePath)
		time.Sleep(30 * time.Second)

	}

}
