package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/jinzhu/gorm"
)

type MockDBConnection struct {
	mock.Mock
}

func (m *MockDBConnection) Connection() *gorm.DB {
	args := m.Called()
	return args.Get(0).(*gorm.DB)
}