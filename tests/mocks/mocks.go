package mocks

import (
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/mock"
)

type MockDBConnection struct {
	mock.Mock
	ReturnNil bool
}

func (m *MockDBConnection) Connection() *gorm.DB {
	args := m.Called()
	if args.Get(0) == true {
		return nil
	}
	return args.Get(0).(*gorm.DB)
}
