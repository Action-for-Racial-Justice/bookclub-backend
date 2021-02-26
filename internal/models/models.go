package models

import "time"

type HealthCheck struct {
	Timestamp time.Time
	Healthy   bool
}
