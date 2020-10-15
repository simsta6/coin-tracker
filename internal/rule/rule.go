package internal

// Rule is ..
type Rule struct {
	CryptoID   string  `json:"crypto_id"`
	Price      float64 `json:"price"`
	Rule       string  `json:"rule"`
	CheckedOut bool
}

func (r *Rule) checkRules() bool {
	return true
}
