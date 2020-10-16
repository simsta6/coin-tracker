package internal

// RawRule is ..
type RawRule struct {
	CryptoID string  `json:"crypto_id"`
	Price    float64 `json:"price"`
	Rule     string  `json:"rule"`
}

type Rule struct {
	CryptoID string
	Price    float64
	Rule     string
}

//CheckRule is for comparing prices
func (r *RawRule) CheckRule(price float64) (ruleUsed bool, err error) {
	switch {
	case r.Rule == "lt":
		if price < r.Price {
			ruleUsed = true
			return ruleUsed, err
		}
	case r.Rule == "gt":
		if price > r.Price {
			ruleUsed = true
			return ruleUsed, err
		}
	}

	return ruleUsed, err
}
