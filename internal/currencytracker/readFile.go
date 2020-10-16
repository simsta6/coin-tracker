package currencytracker

import (
	"encoding/json"
	"os"
)

//Reader is for reading files
type Reader interface {
	ReadRulesJSON(filePath string) (rules RawRules, err error)
	ReadCryptoJSON(filePath string) (rules RawRules, err error)
}

//ReadRulesJSON is for reading rules from specific file
func ReadRulesJSON(filePath string) (rules RawRules, err error) {
	jsonFile, err := os.Open(filePath)

	if err != nil {
		return rules, err
	}
	defer jsonFile.Close()

	err = json.NewDecoder(jsonFile).Decode(&rules)
	if err != nil {
		return rules, err
	}
	return rules, err
}

//ReadCryptocurrencyJSON is for reading cryptocurrency data from specific file
func ReadCryptocurrencyJSON(filePath string) (cryptoCurrency RawCurrencyDataSlice, err error) {
	jsonFile, err := os.Open(filePath)

	if err != nil {
		return cryptoCurrency, err
	}
	defer jsonFile.Close()

	err = json.NewDecoder(jsonFile).Decode(&cryptoCurrency)
	if err != nil {
		return cryptoCurrency, err
	}
	return cryptoCurrency, err
}
