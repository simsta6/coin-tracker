package internal

// Ticker is ..
type Ticker struct {
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
