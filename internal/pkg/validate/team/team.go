package team

import (
	"fmt"
	teamPack "github.com/Constantine27K/crnt-data-manager/pkg/api/team"
)

type Validator interface {
	Check(issue *teamPack.Team) error
}

type checker func(issue *teamPack.Team) error

type validator struct {
	checks []checker
}

func NewValidator() Validator {
	return &validator{
		checks: []checker{
			checkName, checkOwners, checkDepartment,
		},
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

func checkName(team *teamPack.Team) error {
	if len(team.GetName()) == 0 {
		return fmt.Errorf("team's name should not be empty")
	}

	return nil
}

func checkOwners(team *teamPack.Team) error {
	if len(team.GetTechOwner()) == 0 {
		return fmt.Errorf("team should have a tech owner")
	}

	if len(team.GetBusinessOwner()) == 0 {
		return fmt.Errorf("team should have a business owner")
	}

	return nil
}

func checkDepartment(team *teamPack.Team) error {
	if len(team.GetDepartment()) == 0 {
		return fmt.Errorf("team should be linked to a department")
	}

	return nil
}
