package team

import (
	teamPack "github.com/Constantine27K/crnt-data-manager/pkg/api/team"
)

type Validator interface {
	Check(issue *teamPack.Team) error
}

type checker func(issue *teamPack.Team) error

type validator struct {
	checks []checker
}

func NewIssueValidator() Validator {
	return &validator{
		checks: []checker{},
	}
}

func (v *validator) Check(team *teamPack.Team) error {
	for _, check := range v.checks {
		err := check(team)
		if err != nil {
			return err
		}
	}

	return nil
}
