package models

import (
	"time"
)

type Station struct {
	StationId  string `form:"station_id" json:"station_id" binding:"required"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}


