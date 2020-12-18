package api

import "github.com/elemes1/gobridge_2/domain"

type StorageMock struct {
	ListFn func() ([]domain.Todo, error)
}

func (m StorageMock) List() ([]domain.Todo, error) {
	if m.ListFn == nil {
		return nil, nil
	}
	return m.ListFn()
}
