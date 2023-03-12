package fixtures

import (
	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/department"
	"github.com/brianvoe/gofakeit"
)

type optDep func(department *desc.Department)

func CreateDepartment(opts ...optDep) *desc.Department {
	result := &desc.Department{
		Name:     gofakeit.Word(),
		Projects: make([]int64, 0),
	}

	for _, opt := range opts {
		opt(result)
	}

	return result
}

func WithProjects(projs ...int64) optDep {
	return func(department *desc.Department) {
		department.Projects = projs
	}
}
