package models

import (
	"time"
)

type Item struct {
	Id int `json:"id"`
	ItemName  string `form:"item_name" json:"item_name" binding:"required"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}