	package models

import (
	"time"
)

type Reward struct {
	Id int `json:"id"`
	VisitorId  string `json:"member_id"`
	Member_Name  string `json:"member_name"`	
	Member_Country_Region string `json:"country"`
	Member_City  string `json:"city"`
	Member_Email_Address  string `json:"email"`
	Member_Mobile  string `json:"mobile"`
	DateCreated time.Time `json:"date_created"`
}