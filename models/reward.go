package models

import (
	"time"
)

type Reward struct {
	Id int `json:"id"`
	VisitorId  string `form:"member_id" json:"member_id" binding:"required"`
	ItemId  string `form:"item_id" json:"item_id" binding:"required"`
	Qty  int `form:"qty" json:"qty" binding:"required"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}