package models

import (
	"time"
)

type TransactionMember struct {
	Id int `json:"id"`
	MemberId  string `json:"member_id"`
	MemberFname  string `json:"member_fname"`
	MemberLname  string `json:"member_lname"`
	MemberCountryRegion  string `json:"member_country_region"`
	MemberCity  string `json:"member_city"`
	MemberEmailAddress  string `json:"member_email_address"`
	MemberMobile  string `json:"member_mobile"`
	StationId  string `json:"station_id"`
	TransactionType string `json:"transaction_type"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}