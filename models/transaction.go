package models

import (
	"time"
)

type Transaction struct {
	Id int `json:"id"`
	MemberId  string `form:"member_id" json:"member_id" binding:"required"`
	StationId  string `form:"station_id" json:"station_id" binding:"required"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}