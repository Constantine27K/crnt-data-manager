package validate

import (
	issueVal "github.com/Constantine27K/crnt-data-manager/internal/pkg/validate/issue"
	projectVal "github.com/Constantine27K/crnt-data-manager/internal/pkg/validate/project"
	sprintVal "github.com/Constantine27K/crnt-data-manager/internal/pkg/validate/sprint"
	teamVal "github.com/Constantine27K/crnt-data-manager/internal/pkg/validate/team"
	projectPack "github.com/Constantine27K/crnt-data-manager/pkg/api/project"
	sprintPack "github.com/Constantine27K/crnt-data-manager/pkg/api/sprint"
	issuePack "github.com/Constantine27K/crnt-data-manager/pkg/api/tasks/issue"
	teamPack "github.com/Constantine27K/crnt-data-manager/pkg/api/team"
)

type Validator interface {
	CheckIssue(issue *issuePack.Issue) error
	CheckTeam(team *teamPack.Team) error
	CheckProject(project *projectPack.Project) error
	CheckSprint(sprint *sprintPack.Sprint) error
}

type validator struct {
	issue   issueVal.Validator
	team    teamVal.Validator
	project projectVal.Validator
	sprint  sprintVal.Validator
}

func NewValidator() Validator {
	return &validator{
		issue:   issueVal.NewValidator(),
		team:    teamVal.NewValidator(),
		project: projectVal.NewValidator(),
		sprint:  sprintVal.NewValidator(),
	}
}

func (v *validator) CheckIssue(issue *issuePack.Issue) error {
	return v.issue.Check(issue)
}

func (v *validator) CheckTeam(team *teamPack.Team) error {
	return v.team.Check(team)
}

func (v *validator) CheckProject(project *projectPack.Project) error {
	return v.project.Check(project)
}

func (v *validator) CheckSprint(sprint *sprintPack.Sprint) error {
	return v.sprint.Check(sprint)
}
