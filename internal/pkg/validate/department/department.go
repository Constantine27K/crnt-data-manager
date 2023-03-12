package department

import (
	"fmt"

	depPack "github.com/Constantine27K/crnt-data-manager/pkg/api/department"
)

type Validator interface {
	Check(department *depPack.Department) error
}

type checker func(department *depPack.Department) error

type validator struct {
	checks []checker
}

func NewValidator() Validator {
	return &validator{
		checks: []checker{
			checkName,
		},
	}
}

func (v *validator) Check(department *depPack.Department) error {
	for _, check := range v.checks {
		err := check(department)
		if err != nil {
			return err
		}
	}

	return nil
}

func checkName(department *depPack.Department) error {
	if len(department.GetName()) == 0 {
		return fmt.Errorf("department name should not be nil")
	}

	return nil
}
