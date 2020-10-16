package currencytracker

import (
	"strconv"
)

//Parser is for reading files
type Parser interface {
	ParseCrypto() (parsedData CurrencyData, err error)
}

//ParseCrypto is used to parse data from RawCurrencyData to useable data
func (raw RawCurrencyData) ParseCrypto() (parsedData CurrencyData, err error) {
	parsedData.CryptoID = raw.CryptoID
	parsedData.Name = raw.Name
	price, err := strconv.ParseFloat(raw.PriceUSD, 32)
	if err != nil {
		return CurrencyData{}, err
	}
	parsedData.PriceUSD = price

	return parsedData, err
}
