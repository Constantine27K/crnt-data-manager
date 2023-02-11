package validate

import (
	issueVal "github.com/Constantine27K/crnt-data-manager/internal/pkg/validate/issue"
	teamVal "github.com/Constantine27K/crnt-data-manager/internal/pkg/validate/team"
	issuePack "github.com/Constantine27K/crnt-data-manager/pkg/api/tasks/issue"
	teamPack "github.com/Constantine27K/crnt-data-manager/pkg/api/team"
)

type Validator interface {
	CheckIssue(issue *issuePack.Issue) error
	CheckTeam(team *teamPack.Team) error
}

type validator struct {
	issue issueVal.Validator
	team  teamVal.Validator
}

func NewValidator() Validator {
	return &validator{
		issue: issueVal.NewIssueValidator(),
		team:  teamVal.NewIssueValidator(),
	}
}

func (v *validator) CheckIssue(issue *issuePack.Issue) error {
	return v.issue.Check(issue)
}

func (v *validator) CheckTeam(team *teamPack.Team) error {
	return v.team.Check(team)
}
