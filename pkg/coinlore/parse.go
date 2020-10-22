package coinlore

import (
	"strconv"
)

//ParseCrypter is interface for ParseCrypter
type ParseCrypter interface {
	ParseCrypto() (parsedData Currency, err error)
}

//ParseCrypto is used to parse data from RawCurrency to useable data
func (raw RawCurrency) ParseCrypto() (parsedData Currency, err error) {
	parsedData.CryptoID = raw.CryptoID
	parsedData.Symbol = raw.Symbol
	parsedData.Name = raw.Name
	parsedData.NameID = raw.NameID
	parsedData.Rank = raw.Rank

	parsedData.PriceUSD, err = parseFloat(raw.PriceUSD)
	if err != nil {
		return Currency{}, err
	}

	parsedData.PrecentChange24h, err = parseFloat(raw.PrecentChange24h)
	if err != nil {
		return Currency{}, err
	}

	parsedData.PrecentChange1h, err = parseFloat(raw.PrecentChange1h)
	if err != nil {
		return Currency{}, err
	}

	parsedData.PrecentChange7d, err = parseFloat(raw.PrecentChange7d)
	if err != nil {
		return Currency{}, err
	}

	parsedData.MarketCapUSD, err = parseFloat(raw.MarketCapUSD)
	if err != nil {
		return Currency{}, err
	}

	parsedData.Volume24, err = parseFloat(raw.Volume24)
	if err != nil {
		return Currency{}, err
	}

	parsedData.Volume24Native, err = parseFloat(raw.Volume24Native)
	if err != nil {
		return Currency{}, err
	}

	parsedData.CSupply, err = parseFloat(raw.CSupply)
	if err != nil {
		return Currency{}, err
	}

	parsedData.PriceBtc, err = parseFloat(raw.PriceBtc)
	if err != nil {
		return Currency{}, err
	}

	parsedData.TSupply = raw.PriceBtc
	parsedData.MSupply = raw.MSupply

	return parsedData, err
}

func parseFloat(stringToParse string) (value float64, err error) {
	if stringToParse == "0?" || stringToParse == "" {
		return 0, err
	}
	value, err = strconv.ParseFloat(stringToParse, 64)
	if err != nil {
		return 0, err
	}

	return value, err
}
