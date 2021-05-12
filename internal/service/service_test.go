package service_test

import (
	"os"
	"testing"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/mocks"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/service"
	"github.com/golang/mock/gomock"
	"go.uber.org/zap"
)

type testSuite struct {
	mockMysql      *mocks.MockMysql
	mockRequests   *mocks.MockIRequests
	svc            *service.BookClubService
	mockController *gomock.Controller
}

func createTestSuite(t *testing.T) *testSuite {
	mockController := gomock.NewController(t)

	mockSql := mocks.NewMockMysql(mockController)
	mockValidator := mocks.NewMockValidator(mockController)
	mockRequests := mocks.NewMockIRequests(mockController)
	mockLogger := zap.NewNop()
	service := service.New(mockSql, mockRequests, mockValidator, mockLogger)

	return &testSuite{mockMysql: mockSql, svc: service, mockController: mockController}
}

func TestMain(m *testing.M) {
	retCode := m.Run()
	os.Exit(retCode)
}
