package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import source "github.com/MasterOfBinary/gobatch/source"

// Source is an autogenerated mock type for the Source type
type Source struct {
	mock.Mock
}

// Read provides a mock function with given fields: ctx, items, errs
func (_m *Source) Read(ctx context.Context, items chan<- interface{}, errs chan<- error) {
	_m.Called(ctx, items, errs)
}

var _ source.Source = (*Source)(nil)
