package internal

// Rule is ..
type Rule struct {
	CryptoID   string  `json:"crypto_id"`
	Price      float64 `json:"price"`
	Rule       string  `json:"rule"`
	CheckedOut bool
}

func (r *Rule) checkRules() bool {
	//bla bla bla
	return true
}

// func (r *rule) readFromJSON(filePath string) ([]rule, error) {
// 	var rules []rule
// 	jsonFile, err := os.Open(filePath)
// 	defer jsonFile.Close()

// 	if err != nil {
// 		return rules, err
// 	}

// 	err = json.NewDecoder(jsonFile).Decode(&rules)

// 	if err != nil {
// 		return rules, err
// 	}

// 	// err = json.Unmarshal([]byte(byteValue), &rules)

// 	if err != nil {
// 		return rules, err
// 	}

// 	return rules, nil
// }
