package error

import "fmt"

type TenantNotFoundError struct {
	error
}

func NewTenantNotFoundError(id string) error {
	return &TenantNotFoundError{
		error: fmt.Errorf("tenant with id %s not found", id),
	}
}
