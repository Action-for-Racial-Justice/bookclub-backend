package handlers_test

import (
	"os"
	"testing"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/handlers"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/mocks"
	"github.com/golang/mock/gomock"
)

type testSuite struct {
	mockService    *mocks.MockService
	handlers       *handlers.BookClubHandler
	mockController *gomock.Controller
}

func createTestSuite(t *testing.T) *testSuite {
	mockController := gomock.NewController(t)

	mockSvc := mocks.NewMockService(mockController)
	handlers, err := handlers.New(mockSvc)

	if err != nil {
		panic(err)
	}

	return &testSuite{
		mockService:    mockSvc,
		handlers:       handlers,
		mockController: mockController}
}

func TestMain(m *testing.M) {
	retCode := m.Run()
	os.Exit(retCode)
}
