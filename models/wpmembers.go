package models

type WPMember struct {
	MemberId  string `json:"member_id"`
	MemberFname  string `json:"fname"`
	MemberLname  string `json:"lname"`
	MemberCountryRegion  string `json:"region"`
	MemberCity  string `json:"city"`
	MemberEmailAddress  string `json:"email"`
	MemberMobile  string `json:"mobile"`
	MemberRegistrationDate string `json:"registration_date"`
	IsVisited bool `json:"is_visited"`
	ConsentToUserData bool `json:consent_to_use_data`
}


