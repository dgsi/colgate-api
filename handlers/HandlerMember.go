package handlers

import (
	"net/http"
	"time"
	"fmt"
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
	now := time.Now().UTC()
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	email := c.PostForm("email")
	contact_no := c.PostForm("contact_no")
	region := c.PostForm("region")
	municipality := c.PostForm("city")
	is_visited,_ := strconv.ParseBool(c.PostForm("is_visited"))
	
	member := m.WPMember{}
	handler.db.Table("wp_members").Where("member_email_address = ?",email).First(&member)

	if member.MemberEmailAddress != "" {
		respond(http.StatusBadRequest,"Email already taken!",c,true)	
	} else {
		member := m.WPMember{}
		handler.db.Raw("select * from wp_members order by member_id desc limit 1").Scan(&member)
		lastMemberId := member.MemberId;
		fmt.Println("LAST MEMBER ID ----> " + member.MemberId)
		newId, err := strconv.Atoi(lastMemberId)
		if err == nil {
			newMemberId := strconv.Itoa(newId + 1)
			handler.db.Exec("INSERT INTO wp_members(member_id,member_fname,member_lname,member_email_address,member_mobile,member_registration_date,member_country_region,member_city,is_visited) VALUES (?,?,?,?,?,?,?,?,?)",newMemberId,first_name,last_name,email,contact_no,now,region,municipality,is_visited)
	
			newMember := m.WPMember{};
			handler.db.Raw("select * from wp_members order by member_id desc limit 1").Scan(&newMember)
			c.JSON(http.StatusCreated,newMember)
		} else {
			respond(http.StatusInternalServerError,"failed to convert to int",c,true)
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

