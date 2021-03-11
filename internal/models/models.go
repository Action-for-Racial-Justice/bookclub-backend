package models

import "time"

//HealthCheck struct to describe current system health
type HealthCheck struct {
	Timestamp time.Time `json:"timestamp"`
	Healthy   bool      `json:"healthy"`
}
