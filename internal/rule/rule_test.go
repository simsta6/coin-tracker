package rule

import (
	"testing"
)

func TestCompare(t *testing.T) {
	rule := Rule{
		Price:    1000,
		Operator: "lt",
	}

	testCompare(rule, 100, false, t)

	rule = Rule{
		Price:    10,
		Operator: "lt",
	}

	testCompare(rule, 100, true, t)

	rule = Rule{
		Price:    1000,
		Operator: "gt",
	}

	testCompare(rule, 100, true, t)

	rule = Rule{
		Price:    10,
		Operator: "gt",
	}

	testCompare(rule, 100, false, t)
}

func testCompare(rule Rule, price float64, setUsedProperty bool, t *testing.T) {
	err := rule.Compare(price)
	if err != nil {
		t.Error(err)
	}

	if setUsedProperty == rule.Used {
		t.Errorf("Rule.Compare test has failed. Did not set Used field. Expected: %v, got: %v", setUsedProperty, rule.Used)
	}
}
