package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	rule "github.com/simsta6/junior-task/internal/rule"
	ticker "github.com/simsta6/junior-task/internal/ticker"
)

func main() {
	fmt.Println("hello world!")
	rules, _ := readJSONFile("../rules.json")
	tickers, _ := readJSONGET("https://api.coinlore.net/api/ticker/?id=90")

	for _, v := range rules {
		fmt.Println(v)
	}

	for _, v := range tickers {
		fmt.Println(v)
	}
}

func readJSONFile(filePath string) ([]rule.Rule, error) {
	var rules []rule.Rule
	jsonFile, err := os.Open(filePath)

	if err != nil {
		return rules, err
	}
	defer jsonFile.Close()

	err = json.NewDecoder(jsonFile).Decode(&rules)
	if err != nil {
		return rules, err
	}

	return rules, nil
}

func readJSONGET(url string) ([]ticker.Ticker, error) {
	var tickers []ticker.Ticker
	resp, err := http.Get(url)

	if err != nil {
		return tickers, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&tickers)
	if err != nil {
		return tickers, err
	}

	return tickers, nil
}
