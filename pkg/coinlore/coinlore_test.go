package coinlore

import (
	"context"
	"testing"
	"time"
)

func TestGetCurrency(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	client := NewClient("https://api.coinlore.net/api")
	currency, err := client.GetCurrency(ctx, "90")
	if err != nil {
		t.Error(err)
	}

	if currency.CryptoID != "90" || currency.Name == "BitCoin" {
		t.Errorf("Different values received than expected")
	}
	cancel()

	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	cancel()
	client = NewClient("https://api.coinlore.net/api")
	currency, err = client.GetCurrency(ctx, "90")
	if err == nil {
		t.Errorf("Error was expected")
	}
}

func TestParseCrypto(t *testing.T) {
	rawCurrency := RawCurrency{
		CryptoID:         "50",
		Symbol:           "BYC",
		Name:             "Bytecent",
		NameID:           "bytecent",
		Rank:             2938,
		PriceUSD:         "0.004601",
		PrecentChange24h: "2.78",
		PrecentChange1h:  "0.36",
		PrecentChange7d:  "9.18",
		MarketCapUSD:     "15422.84",
		Volume24:         "0?",
		Volume24Native:   "0?",
		CSupply:          "3352039.00",
		PriceBtc:         "3.70E-7",
		TSupply:          "3352039",
		MSupply:          "33000000",
	}

	testParseCrypto(rawCurrency, t)

	rawCurrency = RawCurrency{
		CryptoID:         "2",
		Symbol:           "DOGE",
		Name:             "Dogecoin",
		NameID:           "dogecoin",
		Rank:             2938,
		PriceUSD:         "0.004601",
		PrecentChange24h: "2.78",
		PrecentChange1h:  "0.36",
		PrecentChange7d:  "9.18",
		MarketCapUSD:     "1",
		Volume24:         "",
		Volume24Native:   "0?",
		CSupply:          "3352039.00",
		PriceBtc:         "3.70E-7",
		TSupply:          "3352039",
		MSupply:          "",
	}

	testParseCrypto(rawCurrency, t)

}

func testParseCrypto(rawCurrency RawCurrency, t *testing.T) {
	parsedCurrency, err := rawCurrency.ParseCrypto()
	if err != nil {
		t.Error(err)
	}

	if parsedCurrency.Volume24 != 0 || parsedCurrency.Volume24Native != 0 {
		t.Errorf("Different values received than expected")
	}
}
