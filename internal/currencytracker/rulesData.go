package currencytracker

import (
	"fmt"
)

var (
	greaterThan string
	lessThan    string
)

//InitMessages inits messages
func InitMessages(currency CurrencyData) {
	greaterThan = fmt.Sprintf("Cryptocurrency id:%s %s price was higher than ", currency.CryptoID, currency.Name)
	lessThan = fmt.Sprintf("Cryptocurrency id:%s %s price was less than ", currency.CryptoID, currency.Name)
}

//RawRule is for keeping rule data from .json file
type RawRule struct {
	CryptoID string  `json:"crypto_id"`
	Price    float64 `json:"price"`
	Operator string  `json:"rule"`
	Used     bool    `json:"-"`
}

//RawRules is for keeping all rules from json file
type RawRules []RawRule

//Comparer compares prices while having regard to rule
type Comparer interface {
	Compare(price float64) (answer string, err error)
}

//Compare compaeres prices
func (r *RawRule) Compare(price float64) (answer string, err error) {
	switch {
	case r.Operator == "lt":
		if price < r.Price {
			lessThan += fmt.Sprintf("%.2f", price)
			r.Used = true
			return lessThan, err
		}
	case r.Operator == "gt":
		if price > r.Price {
			greaterThan += fmt.Sprintf("%.2f", price)
			r.Used = true
			return greaterThan, err
		}
	}

	return "", err
}
