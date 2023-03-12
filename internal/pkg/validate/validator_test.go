package validate

import (
	"github.com/Constantine27K/crnt-data-manager/pkg/api/project"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/sprint"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/team"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"

	"github.com/Constantine27K/crnt-data-manager/pkg/api/state/status"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/state/template"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/tasks/issue"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
)

func TestIssueValidator(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name    string
		input   *issue.Issue
		wantErr bool
	}{
		{
			name: "not matching template and status",
			input: &issue.Issue{
				Template: template.Template_TEMPLATE_DEVELOPMENT,
				Status:   &status.IssueStatus{Status: &status.IssueStatus_Common{}},
			},
			wantErr: true,
		},
		{
			name: "no name",
			input: &issue.Issue{
				Template: template.Template_TEMPLATE_DEVELOPMENT,
				Status:   &status.IssueStatus{Status: &status.IssueStatus_Development{Development: &status.IssueStatusDevelopment{Status: status.Development_STATUS_DEVELOPMENT_BACKLOG}}},
			},
			wantErr: true,
		},
		{
			name: "too long name",
			input: &issue.Issue{
				Template: template.Template_TEMPLATE_DEVELOPMENT,
				Status:   &status.IssueStatus{Status: &status.IssueStatus_Development{Development: &status.IssueStatusDevelopment{Status: status.Development_STATUS_DEVELOPMENT_BACKLOG}}},
				Name:     gofakeit.Sentence(60),
			},
			wantErr: true,
		},
		{
			name: "no type",
			input: &issue.Issue{
				Template: template.Template_TEMPLATE_DEVELOPMENT,
				Status:   &status.IssueStatus{Status: &status.IssueStatus_Development{Development: &status.IssueStatusDevelopment{Status: status.Development_STATUS_DEVELOPMENT_BACKLOG}}},
				Name:     gofakeit.Sentence(2),
			},
			wantErr: true,
		},
		{
			name: "no description",
			input: &issue.Issue{
				Template:    template.Template_TEMPLATE_DEVELOPMENT,
				Status:      &status.IssueStatus{Status: &status.IssueStatus_Development{Development: &status.IssueStatusDevelopment{Status: status.Development_STATUS_DEVELOPMENT_BACKLOG}}},
				Name:        gofakeit.Sentence(2),
				Type:        issue.IssueType_ISSUE_TYPE_TASK,
				Description: "",
			},
			wantErr: true,
		},
		{
			name: "too long description",
			input: &issue.Issue{
				Template:    template.Template_TEMPLATE_DEVELOPMENT,
				Status:      &status.IssueStatus{Status: &status.IssueStatus_Development{Development: &status.IssueStatusDevelopment{Status: status.Development_STATUS_DEVELOPMENT_BACKLOG}}},
				Name:        gofakeit.Sentence(2),
				Type:        issue.IssueType_ISSUE_TYPE_TASK,
				Description: gofakeit.Sentence(100),
			},
			wantErr: true,
		},
		{
			name: "no author",
			input: &issue.Issue{
				Template:    template.Template_TEMPLATE_DEVELOPMENT,
				Status:      &status.IssueStatus{Status: &status.IssueStatus_Development{Development: &status.IssueStatusDevelopment{Status: status.Development_STATUS_DEVELOPMENT_BACKLOG}}},
				Name:        gofakeit.Sentence(2),
				Type:        issue.IssueType_ISSUE_TYPE_TASK,
				Description: gofakeit.Sentence(5),
			},
			wantErr: true,
		},
		{
			name: "no project",
			input: &issue.Issue{
				Template:    template.Template_TEMPLATE_DEVELOPMENT,
				Status:      &status.IssueStatus{Status: &status.IssueStatus_Development{Development: &status.IssueStatusDevelopment{Status: status.Development_STATUS_DEVELOPMENT_BACKLOG}}},
				Name:        gofakeit.Sentence(2),
				Type:        issue.IssueType_ISSUE_TYPE_TASK,
				Description: gofakeit.Sentence(5),
				Author:      gofakeit.Name(),
			},
			wantErr: true,
		},
		{
			name: "positive",
			input: &issue.Issue{
				Template:    template.Template_TEMPLATE_DEVELOPMENT,
				Status:      &status.IssueStatus{Status: &status.IssueStatus_Development{Development: &status.IssueStatusDevelopment{Status: status.Development_STATUS_DEVELOPMENT_BACKLOG}}},
				Name:        gofakeit.Sentence(2),
				Type:        issue.IssueType_ISSUE_TYPE_TASK,
				Description: gofakeit.Sentence(5),
				Author:      gofakeit.Name(),
				ProjectId:   5,
			},
			wantErr: false,
		},
	}

	val := NewValidator()

	for _, test := range cases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			err := val.CheckIssue(test.input)
			if test.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestProjectValidator(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name    string
		input   *project.Project
		wantErr bool
	}{
		{
			name:    "no name",
			input:   &project.Project{},
			wantErr: true,
		},
		{
			name: "too long name",
			input: &project.Project{
				Name: gofakeit.Sentence(10),
			},
			wantErr: true,
		},
		{
			name: "no short name",
			input: &project.Project{
				Name: gofakeit.Word(),
			},
			wantErr: true,
		},
		{
			name: "too long short name",
			input: &project.Project{
				Name:      gofakeit.Word(),
				ShortName: gofakeit.Sentence(10),
			},
			wantErr: true,
		},
		{
			name: "no responsible team",
			input: &project.Project{
				Name:      gofakeit.Word(),
				ShortName: "ORG",
			},
			wantErr: true,
		},
		{
			name: "positive",
			input: &project.Project{
				Name:               gofakeit.Word(),
				ShortName:          "ORG",
				ResponsibleTeamIds: []int64{1},
			},
			wantErr: false,
		},
	}

	val := NewValidator()

	for _, test := range cases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			err := val.CheckProject(test.input)
			if test.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestSprintValidator(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name    string
		input   *sprint.Sprint
		wantErr bool
	}{
		{
			name:    "no name",
			input:   &sprint.Sprint{},
			wantErr: true,
		},
		{
			name: "no project",
			input: &sprint.Sprint{
				Name: gofakeit.Word(),
			},
			wantErr: true,
		},
		{
			name: "no start date",
			input: &sprint.Sprint{
				Name:    gofakeit.Word(),
				Project: 1,
			},
			wantErr: true,
		},
		{
			name: "no finish date",
			input: &sprint.Sprint{
				Name:      gofakeit.Word(),
				Project:   1,
				StartedAt: timestamppb.Now(),
			},
			wantErr: true,
		},
		{
			name: "start after finish",
			input: &sprint.Sprint{
				Name:       gofakeit.Word(),
				Project:    1,
				StartedAt:  timestamppb.Now(),
				FinishedAt: timestamppb.New(time.Now().Add(-24 * time.Hour)),
			},
			wantErr: true,
		},
		{
			name: "positive",
			input: &sprint.Sprint{
				Name:       gofakeit.Word(),
				Project:    1,
				StartedAt:  timestamppb.Now(),
				FinishedAt: timestamppb.New(time.Now().Add(24 * time.Hour)),
			},
			wantErr: false,
		},
	}

	val := NewValidator()

	for _, test := range cases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			err := val.CheckSprint(test.input)
			if test.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestTeamValidator(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name    string
		input   *team.Team
		wantErr bool
	}{
		{
			name:    "no name",
			input:   &team.Team{},
			wantErr: true,
		},
		{
			name: "no tech owner",
			input: &team.Team{
				Name: gofakeit.Word(),
			},
			wantErr: true,
		},
		{
			name: "no business owner",
			input: &team.Team{
				Name:      gofakeit.Word(),
				TechOwner: gofakeit.Name(),
			},
			wantErr: true,
		},
		{
			name: "no department",
			input: &team.Team{
				Name:          gofakeit.Word(),
				TechOwner:     gofakeit.Name(),
				BusinessOwner: gofakeit.Name(),
			},
			wantErr: true,
		},
		{
			name: "positive",
			input: &team.Team{
				Name:          gofakeit.Word(),
				TechOwner:     gofakeit.Name(),
				BusinessOwner: gofakeit.Name(),
				Department:    gofakeit.Word(),
			},
			wantErr: false,
		},
	}

	val := NewValidator()

	for _, test := range cases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			err := val.CheckTeam(test.input)
			if test.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
