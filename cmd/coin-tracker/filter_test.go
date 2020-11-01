package main

import (
	"testing"

	"github.com/simsta6/coin-tracker/internal/rule"
)

func TestFilter(t *testing.T) {
	rules := []rule.Rule{
		{CryptoID: "1", Used: true},
		{CryptoID: "2", Used: true},
		{CryptoID: "3", Used: true},
	}
	usedRules := []rule.Rule{}
	testFilter(rules, usedRules, t, "first case")

	rules = []rule.Rule{
		{CryptoID: "1", Used: false},
		{CryptoID: "2", Used: true},
		{CryptoID: "3", Used: true},
	}
	usedRules = []rule.Rule{}
	testFilter(rules, usedRules, t, "second case")

	rules = []rule.Rule{
		{CryptoID: "1", Used: false},
		{CryptoID: "2", Used: true},
		{CryptoID: "3", Used: false},
	}
	usedRules = []rule.Rule{}
	testFilter(rules, usedRules, t, "third case")

	rules = []rule.Rule{
		{CryptoID: "1", Used: false},
		{CryptoID: "2", Used: false},
		{CryptoID: "3", Used: false},
	}
	usedRules = []rule.Rule{}
	testFilter(rules, usedRules, t, "fourth case")
}

func testFilter(rules []rule.Rule, usedRules []rule.Rule, t *testing.T, testCase string) {
	rules, usedRules = filter(rules, usedRules)

	for _, v := range rules {
		if v.Used {
			t.Errorf("filter function does not work properly. %s \nExpected: %v\nGot: %v", testCase, false, v.Used)
		}
	}

	for _, v := range usedRules {
		if !v.Used {
			t.Errorf("filter function does not work properly. %s \nExpected: %v\nGot: %v", testCase, true, v.Used)
		}
	}
}
