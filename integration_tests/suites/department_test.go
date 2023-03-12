package suites

import (
	"github.com/Constantine27K/crnt-data-manager/integration_tests/fixtures"
	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/department"
	"github.com/stretchr/testify/require"
)

func (s *CrntDMSuite) TestDepartment_Create() {
	dep := fixtures.CreateDepartment()

	depCreateResp, err := s.departmentService.CreateDepartment(s.ctx, &desc.DepartmentCreateRequest{Department: dep})
	require.NoError(s.T(), err)
	departmentID := depCreateResp.GetId()

	require.Greater(s.T(), departmentID, int64(0))

	depGetResp, err := s.departmentService.GetDepartmentByID(s.ctx, &desc.DepartmentGetByIDRequest{Id: departmentID})
	require.NoError(s.T(), err)
	gotDep := depGetResp.GetDepartment()

	s.assertDepartment(dep, gotDep)

	// constraint should occur
	_, err = s.departmentService.CreateDepartment(s.ctx, &desc.DepartmentCreateRequest{Department: dep})
	require.Error(s.T(), err)
}

func (s *CrntDMSuite) TestDepartment_Update() {
	dep := fixtures.CreateDepartment()

	depCreateResp, err := s.departmentService.CreateDepartment(s.ctx, &desc.DepartmentCreateRequest{Department: dep})
	require.NoError(s.T(), err)
	departmentID := depCreateResp.GetId()

	require.Greater(s.T(), departmentID, int64(0))

	depGetResp, err := s.departmentService.GetDepartmentByID(s.ctx, &desc.DepartmentGetByIDRequest{Id: departmentID})
	require.NoError(s.T(), err)
	gotDep := depGetResp.GetDepartment()

	s.assertDepartment(dep, gotDep)

	updDep := fixtures.CreateDepartment(
		fixtures.WithProjects(1),
	)

	depUpdateResp, err := s.departmentService.UpdateDepartment(s.ctx, &desc.DepartmentUpdateRequest{Id: departmentID, Department: updDep})
	require.NoError(s.T(), err)
	depUpdatedID := depUpdateResp.GetId()

	require.Equal(s.T(), departmentID, depUpdatedID)

	depGetUpdatedResp, err := s.departmentService.GetDepartmentByID(s.ctx, &desc.DepartmentGetByIDRequest{Id: departmentID})
	require.NoError(s.T(), err)
	gotUpdatedDep := depGetUpdatedResp.GetDepartment()

	s.assertDepartment(updDep, gotUpdatedDep)
}

func (s *CrntDMSuite) TestDepartment_AddProject() {
	dep := fixtures.CreateDepartment()

	depCreateResp, err := s.departmentService.CreateDepartment(s.ctx, &desc.DepartmentCreateRequest{Department: dep})
	require.NoError(s.T(), err)
	departmentID := depCreateResp.GetId()

	require.Greater(s.T(), departmentID, int64(0))

	depGetResp, err := s.departmentService.GetDepartmentByID(s.ctx, &desc.DepartmentGetByIDRequest{Id: departmentID})
	require.NoError(s.T(), err)
	gotDep := depGetResp.GetDepartment()

	s.assertDepartment(dep, gotDep)

	depUpdateResp, err := s.departmentService.DepartmentAddProject(s.ctx, &desc.DepartmentAddProjectRequest{Id: departmentID, ProjectId: 1})
	require.NoError(s.T(), err)
	depUpdatedID := depUpdateResp.GetId()

	require.Equal(s.T(), departmentID, depUpdatedID)

	depGetUpdatedResp, err := s.departmentService.GetDepartmentByID(s.ctx, &desc.DepartmentGetByIDRequest{Id: departmentID})
	require.NoError(s.T(), err)
	gotUpdatedDep := depGetUpdatedResp.GetDepartment()

	require.Equal(s.T(), 1, len(gotUpdatedDep.GetProjects()))
	require.Equal(s.T(), int64(1), gotUpdatedDep.GetProjects()[0])
}

func (s *CrntDMSuite) TestDepartment_RemoveProject() {
	dep := fixtures.CreateDepartment()

	depCreateResp, err := s.departmentService.CreateDepartment(s.ctx, &desc.DepartmentCreateRequest{Department: dep})
	require.NoError(s.T(), err)
	departmentID := depCreateResp.GetId()

	require.Greater(s.T(), departmentID, int64(0))

	depGetResp, err := s.departmentService.GetDepartmentByID(s.ctx, &desc.DepartmentGetByIDRequest{Id: departmentID})
	require.NoError(s.T(), err)
	gotDep := depGetResp.GetDepartment()

	s.assertDepartment(dep, gotDep)

	depUpdateResp, err := s.departmentService.DepartmentAddProject(s.ctx, &desc.DepartmentAddProjectRequest{Id: departmentID, ProjectId: 1})
	require.NoError(s.T(), err)
	depUpdatedID := depUpdateResp.GetId()

	require.Equal(s.T(), departmentID, depUpdatedID)

	depGetUpdatedResp, err := s.departmentService.GetDepartmentByID(s.ctx, &desc.DepartmentGetByIDRequest{Id: departmentID})
	require.NoError(s.T(), err)
	gotUpdatedDep := depGetUpdatedResp.GetDepartment()

	require.Equal(s.T(), 1, len(gotUpdatedDep.GetProjects()))
	require.Equal(s.T(), int64(1), gotUpdatedDep.GetProjects()[0])

	depRemoveResp, err := s.departmentService.DepartmentRemoveProject(s.ctx, &desc.DepartmentRemoveProjectRequest{Id: departmentID, ProjectId: 1})
	require.NoError(s.T(), err)
	depRemovedID := depRemoveResp.GetId()

	require.Equal(s.T(), departmentID, depRemovedID)

	depGetRemovedResp, err := s.departmentService.GetDepartmentByID(s.ctx, &desc.DepartmentGetByIDRequest{Id: departmentID})
	require.NoError(s.T(), err)
	gotRemovedDep := depGetRemovedResp.GetDepartment()

	require.Equal(s.T(), 0, len(gotRemovedDep.GetProjects()))
}

func (s *CrntDMSuite) assertDepartment(expected, actual *desc.Department) {
	require.Equal(s.T(), expected.GetName(), actual.GetName())
	require.Equal(s.T(), expected.GetProjects(), actual.GetProjects())
}
