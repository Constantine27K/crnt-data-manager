package project

import (
	"fmt"

	projectPack "github.com/Constantine27K/crnt-data-manager/pkg/api/project"
)

type Validator interface {
	Check(issue *projectPack.Project) error
}

type checker func(issue *projectPack.Project) error

type validator struct {
	checks []checker
}

func NewValidator() Validator {
	return &validator{
		checks: []checker{
			checkNames, checkTeams,
		},
	}
}

func (v *validator) Check(project *projectPack.Project) error {
	for _, check := range v.checks {
		err := check(project)
		if err != nil {
			return err
		}
	}

	return nil
}

func checkNames(project *projectPack.Project) error {
	if len(project.GetName()) == 0 || len(project.GetName()) > 15 {
		return fmt.Errorf("project name should be less or equal to 15 symbols")
	}

	if len(project.GetShortName()) < 3 || len(project.GetShortName()) > 5 {
		return fmt.Errorf("project short name should be greater or equal to 3 symbols, or less or equal to 5 symbols")
	}

	return nil
}

func checkTeams(project *projectPack.Project) error {
	if len(project.GetResponsibleTeamIds()) == 0 {
		return fmt.Errorf("project should have at least one responsible team")
	}

	return nil
}
