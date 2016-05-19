package models

import "time"

type Member struct {
	MemberId  string `json:"member_id"`
	MemberFname  string `json:"member_fname"`
	MemberLname  string `json:"member_lname"`
	MemberCountryRegion  string `json:"member_country_region"`
	MemberCity  string `json:"member_city"`
	MemberEmailAddress  string `json:"member_email_address"`
	MemberMobile  string `json:"member_mobile"`
	IsVisited bool `json:"is_visited"`
	ConsetToUserData bool `json:consent_to_use_data`
	MemberRegistrationDate time.Time `json:"member_registration_date"`
}


