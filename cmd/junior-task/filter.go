package main

import (
	"github.com/simsta6/junior-task/internal/rule"
)

func filter(rules []rule.Rule, usedRules []rule.Rule) ([]rule.Rule, []rule.Rule) {
	tmp := rules[:0]
	for _, rule := range rules {
		if !rule.Used {
			tmp = append(tmp, rule)
		} else {
			usedRules = append(usedRules, rule)
		}
	}

	return tmp, usedRules
}
