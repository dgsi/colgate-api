package handlers

import (
	"net/http"
	"time"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	m "colgate/dgsi/api/models"
)

type RewardHandler struct {
	db *gorm.DB
}

func NewRewardHandler(db *gorm.DB) *RewardHandler {
	return &RewardHandler{db}
}

// get all rewards
func (handler RewardHandler) Index(c *gin.Context) {
	rewards := []m.Reward{}	
	handler.db.Raw("SELECT * from qry_rewards ORDER BY date_created DESC").Scan(&rewards)
	c.JSON(http.StatusOK, &rewards)
}

// create new reward
func (handler RewardHandler) Create(c *gin.Context) {
	now := time.Now().UTC()
	member_id := c.PostForm("member_id")

	member := m.Member{}
	handler.db.Table("wp_members").Where("member_id = ?",member_id).First(&member)

	if (member.MemberId != "") {
		transactions := []m.Transaction{}
		handler.db.Table("transaction").Where("member_id = ?",member_id).Find(&transactions)	

		if len(transactions) < 4 {
			respond(http.StatusBadRequest,"Visitor is required to visit all zones",c,true)	
		} else {
			reward := m.Reward{}
			handler.db.Table("reward").Where("visitor_id = ?",member_id).First(&reward)
			if (reward.VisitorId == "") {
				handler.db.Exec("INSERT INTO reward VALUES (null,?,'n/a',?,?,?)",member_id,0,now,now)
				respond(http.StatusCreated,"Rewards successfully claimed by user",c,false)	
			} else {
				respond(http.StatusBadRequest,"Visitor already claimed his/her reward!",c,true)
			}
		}
	} else {
		respond(http.StatusBadRequest,"Member not found!",c,true)
	}
}

func (handler RewardHandler) GetRewardsByUser(c *gin.Context) {
	member_id := c.Param("member_id")
	rewards := []m.Reward{}
	handler.db.Table("qry_rewards").Where("visitor_id = ?",member_id).Find(&rewards)
	c.JSON(http.StatusOK,rewards);
}