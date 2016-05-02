package models

import "time"

type WPMember struct {
	MemberId  string `json:"member_id"`
	MemberFname  string `json:"fname"`
	MemberLname  string `json:"lname"`
	MemberCountryRegion  string `json:"region"`
	MemberCity  string `json:"city"`
	MemberEmailAddress  string `json:"email"`
	MemberMobile  string `json:"mobile"`
	MemberRegisrationDate string `json:"registration_date"`
	MemberUpdateDate time.Time `json:"update_date"`
	MemberUpdateCount int `json:"update_count"`
	isSync string `json:"is_sync"`
	isPrint string `json:"is_print"`
}


