package currencytracker

import (
	"encoding/json"
	"os"
)

//WriteRulesJSON is
func WriteRulesJSON(rules RawRules, filePath string) (err error) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(rules)

	if err != nil {
		return err
	}
	return nil
}
