package rule

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

//Rule is for keeping rule data from .json file
type Rule struct {
	CryptoID string  `json:"crypto_id"`
	Price    float64 `json:"price"`
	Operator string  `json:"rule"`
	Used     bool    `json:"-"`
}

//ReadRulesJSON is for reading rules from specific file
func ReadRulesJSON(filePath string) (rules []Rule, err error) {
	jsonFile, err := os.Open(filePath)

	if err != nil {
		return rules, err
	}
	defer jsonFile.Close()

	if err = json.NewDecoder(jsonFile).Decode(&rules); err != nil {
		return rules, err
	}
	return rules, err
}

//WriteJSON is
func WriteJSON(data []Rule, filePath string) (err error) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	if err = json.NewEncoder(file).Encode(data); err != nil {
		return err
	}
	return nil
}

//Compare compaeres prices
func (r *Rule) Compare(price float64) (err error) {
	switch {
	case r.Operator == "lt":
		if price < r.Price {
			r.Used = true
			return err
		}
	case r.Operator == "gt":
		if price > r.Price {
			r.Used = true
			return err
		}
	default:
		return errors.New("The rule was not in the right format in the file. Use gt or lt only")
	}
	return err
}

//ToString is
func (r *Rule) ToString() (answer string) {
	if !r.Used {
		return ""
	}

	var operator string
	switch r.Operator {
	case "lt":
		operator = "less than"
	case "gt":
		operator = "greater than"
	}

	answer = fmt.Sprintf("Cryptocurrency id:%s price was %s %.2f", r.CryptoID, operator, r.Price)

	return answer
}
