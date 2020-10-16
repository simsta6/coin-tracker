package currencytracker

//RawCurrencyData structure
type RawCurrencyData struct {
	CryptoID         string `json:"id"`
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	NameID           string `json:"nameid"`
	Rank             int    `json:"rank"`
	PriceUSD         string `json:"price_usd"`
	PrecentChange24h string `json:"percent_change_24h"`
	PrecentChange1h  string `json:"percent_change_1h"`
	PrecentChange7d  string `json:"percent_change_7d"`
	MarketCapUSD     string `json:"market_cap_usd"`
	Volume24         string `json:"volume24"`
	Volume24Native   string `json:"volume24_native"`
	CSupply          string `json:"csupply"`
	PriceBtc         string `json:"price_btc"`
	TSupply          string `json:"tsupply"`
	MSupply          string `json:"msupply"`
}

//CurrencyData struct is only for values that will be used in program
type CurrencyData struct {
	CryptoID string
	Name     string
	PriceUSD float64
}

//RawCurrencyDataSlice is for keeping data from GET call from API
type RawCurrencyDataSlice []RawCurrencyData

//CurrencyDataSlice is only for values that will be used in program
type CurrencyDataSlice []CurrencyData
