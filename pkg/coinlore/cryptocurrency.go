package coinlore

//RawCurrency structure is for keeping data from api
type RawCurrency struct {
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

//Currency struct is for parsed data from RawCurrency
type Currency struct {
	CryptoID         string
	Symbol           string
	Name             string
	NameID           string
	Rank             int
	PriceUSD         float64
	PrecentChange24h float64
	PrecentChange1h  float64
	PrecentChange7d  float64
	MarketCapUSD     float64
	Volume24         float64
	Volume24Native   float64
	CSupply          float64
	PriceBtc         float64
	TSupply          string
	MSupply          string
}
