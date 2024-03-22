package main

import (
	"fmt"
	"os"

	"github.com/xackery/yakuku/rule"
	"gopkg.in/yaml.v3"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println("Failed:", err)
		os.Exit(1)
	}
}

func run() error {
	baseRules, err := loadRule("data/rule_base.yaml")
	if err != nil {
		return fmt.Errorf("load base rule: %w", err)
	}

	rules, err := loadRule("data/rule.yaml")
	if err != nil {
		return fmt.Errorf("load rule: %w", err)
	}

	// remove baseRules that are in rules
	for _, rule := range rules.Rules {
		for i, baseRule := range baseRules.Rules {
			if rule.Name == baseRule.Name {
				baseRules.Rules = append(baseRules.Rules[:i], baseRules.Rules[i+1:]...)
				break
			}
		}
	}

	out := `
---
title: Rule Values
weight: 0
bookToC: false
---

## Changed Values

Here is a list of rules that have been changed from base PEQ:

name|value|Notes
---|---|---
`

	for _, rule := range rules.Rules {
		out += fmt.Sprintf("%s|%s|%s\n", rule.Name, rule.Value, rule.Notes)
	}

	out += `
## Default Values

Here is a list of rules that are default in PEQ:

name|value|Notes
---|---|---
`

	for _, rule := range baseRules.Rules {
		out += fmt.Sprintf("%s|%s|%s\n", rule.Name, rule.Value, rule.Notes)
	}

	err = os.WriteFile("content/rule/_index.en.md", []byte(out), 0644)
	if err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}

func loadRule(path string) (*rule.RuleYaml, error) {
	r, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	rule := &rule.RuleYaml{}
	dec := yaml.NewDecoder(r)
	dec.KnownFields(true)
	err = dec.Decode(rule)
	if err != nil {
		return nil, fmt.Errorf("decode: %w", err)
	}

	return rule, nil
}
