package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	ts := createTestSuite(t)
	health := ts.svc.CheckHealth()

	assert.True(t, health.Healthy, "Ensuring health is healthy")
}
