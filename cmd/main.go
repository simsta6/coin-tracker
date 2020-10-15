package main

import (
	"encoding/json"
	"fmt"
	"os"

	r "github.com/simsta6/junior-task/internal"
)

func main() {
	fmt.Println("hello world!")
	rules, _ := readFromJSON("../rules.json")

	for _, v := range rules {
		fmt.Println(v)
	}
}

func readFromJSON(filePath string) ([]r.Rule, error) {
	var rules []r.Rule
	jsonFile, err := os.Open(filePath)
	defer jsonFile.Close()

	if err != nil {
		return rules, err
	}

	err = json.NewDecoder(jsonFile).Decode(&rules)

	if err != nil {
		return rules, err
	}

	return rules, nil
}
