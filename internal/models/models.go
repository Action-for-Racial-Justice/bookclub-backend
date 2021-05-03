package models

import "time"

// HealthCheck contains a timestamp and a boolean indicator
// swagger:model HealthCheck
//HealthCheck struct to describe current system health
type HealthCheck struct {
	Timestamp time.Time `json:"timestamp"`
	Healthy   bool      `json:"healthy"`
}
