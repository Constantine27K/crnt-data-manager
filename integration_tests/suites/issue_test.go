package suites

import (
	"fmt"
	"time"

	"github.com/Constantine27K/crnt-data-manager/integration_tests/fixtures"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/comments"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/project"
	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/tasks/issue"
	"github.com/stretchr/testify/require"
)

func (s *CrntDMSuite) TestIssue_Create() {
	proj := fixtures.CreateProject(
		fixtures.WithResponsibleTeams(1),
	)

	respCreateProject, err := s.projectService.CreateProject(s.ctx, &project.ProjectCreateRequest{Project: proj})
	require.NoError(s.T(), err)
	projectID := respCreateProject.GetId()

	require.Greater(s.T(), projectID, int64(0))

	respGetProject, err := s.projectService.GetProjectByID(s.ctx, &project.ProjectGetByIDRequest{Id: projectID})
	require.NoError(s.T(), err)
	gotProject := respGetProject.GetProject()

	s.assertProject(proj, gotProject)

	issue := fixtures.CreateIssue(
		fixtures.WithProject(respCreateProject.GetId()),
	)

	respCreateIssue, err := s.issueService.CreateIssue(s.ctx, &desc.IssueCreateRequest{Issue: issue})
	require.NoError(s.T(), err)
	issueID := respCreateIssue.GetId()

	require.Greater(s.T(), issueID, int64(0))

	respGetIssue, err := s.issueService.GetIssueByID(s.ctx, &desc.IssueGetByIDRequest{Id: issueID})
	require.NoError(s.T(), err)
	gotIssue := respGetIssue.GetIssue()

	s.assertIssues(issue, gotIssue)
	s.assertCompositeName(proj, gotIssue, 1)
}

func (s *CrntDMSuite) TestIssue_GetInfo() {
	proj := fixtures.CreateProject(
		fixtures.WithResponsibleTeams(1),
	)

	respCreateProject, err := s.projectService.CreateProject(s.ctx, &project.ProjectCreateRequest{Project: proj})
	require.NoError(s.T(), err)
	projectID := respCreateProject.GetId()

	require.Greater(s.T(), projectID, int64(0))

	respGetProject, err := s.projectService.GetProjectByID(s.ctx, &project.ProjectGetByIDRequest{Id: projectID})
	require.NoError(s.T(), err)
	gotProject := respGetProject.GetProject()

	s.assertProject(proj, gotProject)

	issue := fixtures.CreateIssue(
		fixtures.WithProject(respCreateProject.GetId()),
	)

	respCreateIssue, err := s.issueService.CreateIssue(s.ctx, &desc.IssueCreateRequest{Issue: issue})
	require.NoError(s.T(), err)
	issueID := respCreateIssue.GetId()

	require.Greater(s.T(), issueID, int64(0))

	respGetIssue, err := s.issueService.GetIssueByID(s.ctx, &desc.IssueGetByIDRequest{Id: issueID})
	require.NoError(s.T(), err)
	gotIssue := respGetIssue.GetIssue()

	s.assertIssues(issue, gotIssue)
	s.assertCompositeName(proj, gotIssue, 1)

	respGetIssueInfo, err := s.issueService.GetIssueInfoByID(s.ctx, &desc.IssueInfoGetByIDRequest{Id: issueID})
	require.NoError(s.T(), err)
	gotIssueInfo := respGetIssueInfo.GetIssueInfo()

	s.assertIssueInfo(gotIssue, gotIssueInfo)
}

func (s *CrntDMSuite) TestIssue_Update() {
	proj := fixtures.CreateProject(
		fixtures.WithResponsibleTeams(1),
	)

	respCreateProject, err := s.projectService.CreateProject(s.ctx, &project.ProjectCreateRequest{Project: proj})
	require.NoError(s.T(), err)
	projectID := respCreateProject.GetId()

	require.Greater(s.T(), projectID, int64(0))

	respGetProject, err := s.projectService.GetProjectByID(s.ctx, &project.ProjectGetByIDRequest{Id: projectID})
	require.NoError(s.T(), err)
	gotProject := respGetProject.GetProject()

	s.assertProject(proj, gotProject)

	issue := fixtures.CreateIssue(
		fixtures.WithProject(respCreateProject.GetId()),
	)

	respCreateIssue, err := s.issueService.CreateIssue(s.ctx, &desc.IssueCreateRequest{Issue: issue})
	require.NoError(s.T(), err)
	issueID := respCreateIssue.GetId()

	require.Greater(s.T(), issueID, int64(0))

	respGetIssue, err := s.issueService.GetIssueByID(s.ctx, &desc.IssueGetByIDRequest{Id: issueID})
	require.NoError(s.T(), err)
	gotIssue := respGetIssue.GetIssue()

	s.assertIssues(issue, gotIssue)
	s.assertCompositeName(proj, gotIssue, 1)

	updIssue := &desc.Issue{Name: "New Name"}
	respUpdateIssue, err := s.issueService.UpdateIssue(s.ctx, &desc.IssueUpdateRequest{Id: issueID, Issue: updIssue})
	require.NoError(s.T(), err)
	issueUpdateID := respUpdateIssue.GetId()

	require.Equal(s.T(), issueID, issueUpdateID)

	respGetUpdatedIssue, err := s.issueService.GetIssueByID(s.ctx, &desc.IssueGetByIDRequest{Id: issueID})
	require.NoError(s.T(), err)
	gotUpdatedIssue := respGetUpdatedIssue.GetIssue()

	require.Equal(s.T(), gotUpdatedIssue.GetName(), "New Name")
}

func (s *CrntDMSuite) TestIssue_Comments() {
	proj := fixtures.CreateProject(
		fixtures.WithResponsibleTeams(1),
	)

	respCreateProject, err := s.projectService.CreateProject(s.ctx, &project.ProjectCreateRequest{Project: proj})
	require.NoError(s.T(), err)
	projectID := respCreateProject.GetId()

	require.Greater(s.T(), projectID, int64(0))

	respGetProject, err := s.projectService.GetProjectByID(s.ctx, &project.ProjectGetByIDRequest{Id: projectID})
	require.NoError(s.T(), err)
	gotProject := respGetProject.GetProject()

	s.assertProject(proj, gotProject)

	issue := fixtures.CreateIssue(
		fixtures.WithProject(respCreateProject.GetId()),
		fixtures.WithComment("Bob", "Any dummy comment"),
	)

	respCreateIssue, err := s.issueService.CreateIssue(s.ctx, &desc.IssueCreateRequest{Issue: issue})
	require.NoError(s.T(), err)
	issueID := respCreateIssue.GetId()

	require.Greater(s.T(), issueID, int64(0))

	respGetIssue, err := s.issueService.GetIssueByID(s.ctx, &desc.IssueGetByIDRequest{Id: issueID})
	require.NoError(s.T(), err)
	gotIssue := respGetIssue.GetIssue()

	s.assertIssues(issue, gotIssue)
	s.assertCompositeName(proj, gotIssue, 1)
	s.assertComments(issue.GetComments(), gotIssue.GetComments())
}

func (s *CrntDMSuite) TestIssue_Children() {
	proj := fixtures.CreateProject(
		fixtures.WithResponsibleTeams(1),
	)

	respCreateProject, err := s.projectService.CreateProject(s.ctx, &project.ProjectCreateRequest{Project: proj})
	require.NoError(s.T(), err)
	projectID := respCreateProject.GetId()

	require.Greater(s.T(), projectID, int64(0))

	respGetProject, err := s.projectService.GetProjectByID(s.ctx, &project.ProjectGetByIDRequest{Id: projectID})
	require.NoError(s.T(), err)
	gotProject := respGetProject.GetProject()

	s.assertProject(proj, gotProject)

	issue := fixtures.CreateIssue(
		fixtures.WithProject(respCreateProject.GetId()),
		fixtures.WithComment("Bob", "Any dummy comment"),
	)

	respCreateIssue, err := s.issueService.CreateIssue(s.ctx, &desc.IssueCreateRequest{Issue: issue})
	require.NoError(s.T(), err)
	issueID := respCreateIssue.GetId()

	require.Greater(s.T(), issueID, int64(0))

	child := fixtures.CreateChildIssue(
		fixtures.WithProject(respCreateProject.GetId()),
	)

	respCreateChild, err := s.issueService.CreateSubtask(s.ctx, &desc.IssueCreateSubtaskRequest{Id: issueID, Child: child})
	require.NoError(s.T(), err)
	childID := respCreateChild.GetId()

	require.Greater(s.T(), childID, int64(0))

	respGetIssue, err := s.issueService.GetIssueByID(s.ctx, &desc.IssueGetByIDRequest{Id: issueID})
	require.NoError(s.T(), err)
	gotIssue := respGetIssue.GetIssue()

	s.assertIssues(issue, gotIssue)
	s.assertCompositeName(proj, gotIssue, 1)
	s.assertComments(issue.GetComments(), gotIssue.GetComments())

	require.Equal(s.T(), 1, len(gotIssue.GetChildren()))
	gotChild := gotIssue.GetChildren()[0]
	s.assertIssues(child, gotChild)
	s.assertCompositeName(proj, gotChild, 2)
	s.assertComments(child.GetComments(), gotChild.GetComments())
}

func (s *CrntDMSuite) assertIssues(expected, actual *desc.Issue) {
	require.Equal(s.T(), expected.GetName(), actual.GetName())
	require.Equal(s.T(), expected.GetParentId(), actual.GetParentId())
	require.Equal(s.T(), expected.GetType(), actual.GetType())
	require.Equal(s.T(), expected.GetDescription(), actual.GetDescription())
	require.Equal(s.T(), expected.GetAuthor(), actual.GetAuthor())
	require.Equal(s.T(), expected.GetAssigned(), actual.GetAssigned())
	require.Equal(s.T(), expected.GetQa(), actual.GetQa())
	require.Equal(s.T(), expected.GetReviewer(), actual.GetReviewer())
	require.Equal(s.T(), expected.GetTemplate(), actual.GetTemplate())
	require.WithinDuration(s.T(), expected.GetCreatedAt().AsTime(), actual.GetCreatedAt().AsTime(), 1*time.Second)
	require.WithinDuration(s.T(), expected.GetDeadline().AsTime(), actual.GetDeadline().AsTime(), 1*time.Second)
	require.Equal(s.T(), expected.GetStatus(), actual.GetStatus())
	require.Equal(s.T(), expected.GetPriority(), actual.GetPriority())
	require.Equal(s.T(), expected.GetStoryPoints(), actual.GetStoryPoints())
	require.Equal(s.T(), expected.GetSprintId(), actual.GetSprintId())
	require.Equal(s.T(), expected.GetProjectId(), actual.GetProjectId())
	require.Equal(s.T(), expected.GetPayment(), actual.GetPayment())
}

func (s *CrntDMSuite) assertIssueInfo(issue *desc.Issue, info *desc.IssueInfo) {
	require.Equal(s.T(), issue.GetId(), info.GetId())
	require.Equal(s.T(), issue.GetName(), info.GetName())
	require.Equal(s.T(), issue.GetCompositeName(), info.GetCompositeName())
	require.Equal(s.T(), issue.GetType(), info.GetType())
	require.Equal(s.T(), issue.GetAssigned(), info.GetAssigned())
	require.Equal(s.T(), issue.GetStatus(), info.GetStatus())
	require.Equal(s.T(), issue.GetPriority(), info.GetPriority())
	require.Equal(s.T(), issue.GetStoryPoints(), info.GetStoryPoints())
}

func (s *CrntDMSuite) assertCompositeName(proj *project.Project, issue *desc.Issue, number int) {
	require.Equal(s.T(), fmt.Sprintf("%s-%v", proj.GetShortName(), number), issue.GetCompositeName())
}

func (s *CrntDMSuite) assertComments(expected, actual *comments.Comments) {
	expComms, actComms := expected.GetItems(), actual.GetItems()
	require.Equal(s.T(), len(expComms), len(actComms))

	for i := 0; i < len(expComms); i++ {
		require.Equal(s.T(), expComms[i].GetAuthor(), actComms[i].GetAuthor())
		require.Equal(s.T(), expComms[i].GetText(), actComms[i].GetText())
		require.WithinDuration(s.T(), expComms[i].GetWrittenAt().AsTime(), actComms[i].GetWrittenAt().AsTime(), 1*time.Second)
	}
}
