package handlers

import (
	"net/http"
	"time"
	"strconv"
	
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	m "colgate/dgsi/api/models"
)

type MemberHandler struct {
	db *gorm.DB
}

func NewMemberHandler(db *gorm.DB) *MemberHandler {
	return &MemberHandler{db}
}

// get all members
func (handler MemberHandler) Index(c *gin.Context) {
	members := []m.WPMember{}	
	handler.db.Table("wp_members").Order("member_id desc").Find(&members)
	c.JSON(http.StatusOK, &members)
}

// create new member
func (handler MemberHandler) Create(c *gin.Context) {
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	email := c.PostForm("email")
	contact_no := c.PostForm("contact_no")
	region := c.PostForm("region")
	city := c.PostForm("city")
	is_visited,_ := strconv.ParseBool(c.PostForm("is_visited"))
	consent,_ := strconv.ParseBool(c.PostForm("consent"))
	
	member := m.WPMember{}
	query := handler.db.Table("wp_members").Last(&member)
	var newMemberId string

	if query.RowsAffected > 0 {
		newId, _ := strconv.Atoi(member.MemberId)
		newMemberId = strconv.Itoa(newId + 1)
	} else {
		year := strconv.Itoa(time.Now().UTC().Year())
		newMemberId = year[2:] + "000001"
	}

	existingMember := m.WPMember{}

	if handler.db.Where("member_email_address = ?",email).First(&existingMember).RowsAffected == 1 {
		respond(http.StatusBadRequest,"Email already taken!",c,true)	
	} else {
		newMember := m.WPMember{}
		newMember.MemberId = newMemberId
		newMember.MemberFname = first_name
		newMember.MemberLname = last_name
		newMember.MemberCountryRegion = region
		newMember.MemberCity = city
		newMember.MemberEmailAddress = email
		newMember.MemberMobile = contact_no
		newMember.IsVisited = is_visited
		newMember.ConsentToUserData = consent
		newMember.MemberRegistrationDate = time.Now().UTC().String()

		result := handler.db.Table("wp_members").Create(&newMember)
		if result.RowsAffected > 0 {
			c.JSON(http.StatusCreated,newMember)
		} else {
			respond(http.StatusBadRequest,result.Error.Error(),c,true)
		}
	}
}

// search a member
func (handler MemberHandler) Search(c *gin.Context) {
	first_name := c.Query("first_name")
	last_name := c.Query("last_name")
	members := []m.WPMember{}
	handler.db.Table("wp_members").Where("member_fname = ? and member_lname = ?",first_name,last_name).Find(&members)
	c.JSON(http.StatusOK,members)
}

func (handler MemberHandler) SearchById(c *gin.Context) {
	member_id := c.Param("member_id")
	member := m.WPMember{}
	handler.db.Table("wp_members").Where("member_id = ?",member_id).First(&member)
	c.JSON(http.StatusOK,member)
}

