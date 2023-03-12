package sprint

import (
	"fmt"

	sprintPack "github.com/Constantine27K/crnt-data-manager/pkg/api/sprint"
)

type Validator interface {
	Check(issue *sprintPack.Sprint) error
}

type checker func(issue *sprintPack.Sprint) error

type validator struct {
	checks []checker
}

func NewValidator() Validator {
	return &validator{
		checks: []checker{
			checkName, checkProject, checkDates,
		},
	}
}

func (v *validator) Check(sprint *sprintPack.Sprint) error {
	for _, check := range v.checks {
		err := check(sprint)
		if err != nil {
			return err
		}
	}

	return nil
}

func checkName(sprint *sprintPack.Sprint) error {
	if len(sprint.GetName()) == 0 {
		return fmt.Errorf("sprint should have a name")
	}

	return nil
}

func checkProject(sprint *sprintPack.Sprint) error {
	if sprint.GetProject() == 0 {
		return fmt.Errorf("sprint should be linked to a project")
	}

	return nil
}

func checkDates(sprint *sprintPack.Sprint) error {
	if !sprint.GetStartedAt().IsValid() || sprint.GetStartedAt().GetSeconds() == 0 {
		return fmt.Errorf("sprint should have a start date")
	}

	if !sprint.GetFinishedAt().IsValid() || sprint.GetFinishedAt().GetSeconds() == 0 {
		return fmt.Errorf("sprint should have a finish date")
	}

	if sprint.GetStartedAt().AsTime().After(sprint.GetFinishedAt().AsTime()) {
		return fmt.Errorf("finish date should after start date")
	}

	return nil
}
