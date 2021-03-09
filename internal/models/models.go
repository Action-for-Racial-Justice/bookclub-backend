package models

import "time"

//HealthCheck struct to describe current system health
type HealthCheck struct {
	Timestamp time.Time
	Healthy   bool
}
