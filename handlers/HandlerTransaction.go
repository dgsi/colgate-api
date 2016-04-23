package handlers

import (
	"net/http"
	"time"
	
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	m "colgate/dgsi/api/models"
)

type TransactionHandler struct {
	db *gorm.DB
}

func NewTransactionHandler(db *gorm.DB) *TransactionHandler {
	return &TransactionHandler{db}
}

// get all transaction
func (handler TransactionHandler) Index(c *gin.Context) {
	transactions := []m.Transaction{}	
	handler.db.Table("transaction").Find(&transactions)
	c.JSON(http.StatusOK, &transactions)
}

// create new transaction
func (handler TransactionHandler) Create(c *gin.Context) {
	now := time.Now().UTC()
	member_id := c.PostForm("member_id")
	station_id := c.PostForm("station_id")

	if member_id == "" {
		respond(http.StatusBadRequest,"Please supply the member's id!",c,true)
	} else if station_id == "" {
		respond(http.StatusBadRequest,"Please supply the station id!",c,true)
	} else {
		station := m.Station{}
		handler.db.Table("station").Where("station_id = ?",station_id).First(&station)

		if station.StationId != "" {
			member := m.Member{}
			handler.db.Table("wp_members").Where("member_id = ?",member_id).First(&member)

			if (member.MemberId != "") {
				transaction := m.Transaction{}
				handler.db.Table("transaction").Where("member_id = ? AND station_id = ?",member_id,station_id).First(&transaction)	
			
				if transaction.MemberId == "" {
					handler.db.Exec("INSERT INTO transaction VALUES (null,?,?,?,?)",member_id,station_id,now,now)
					respond(http.StatusCreated,"New transaction successfully created!",c,false)			
				} else {
					respond(http.StatusBadRequest,"Member already visited this station!",c,true)	
				}
			} else {
				respond(http.StatusBadRequest,"Member not found!",c,true)	
			}
		} else {
			respond(http.StatusBadRequest,"Station not found!",c,true)	
		}
	}
}

// show all transaction of a member
func (handler TransactionHandler) ShowMemberTransactions(c *gin.Context) {
	member_id := c.Param("member_id")
	if member_id != "" {
		member := m.Member{}
		handler.db.Table("wp_members").Where("member_id = ?",member_id).First(&member)

		if (member.MemberId != "") {
			transactions := []m.Transaction{}	
			handler.db.Table("transaction").Where("member_id = ?",member_id).Find(&transactions)
			c.JSON(http.StatusOK, &transactions)		
		} else {
			respond(http.StatusBadRequest,"Member not found!",c,true)
		}
	} else {
		respond(http.StatusBadRequest,"Please supply the member's id!",c,true)
	}
}

//show all transaction by a station
func (handler TransactionHandler) ShowStationTransactions(c *gin.Context) {
	station_id := c.Param("station_id")
	if station_id != "" {
		station := m.Station{}
		handler.db.Table("station").Where("station_id = ?",station_id).First(&station)

		if station.StationId != "" {	
			query := []m.TransactionMember{}	
			handler.db.Raw("SELECT t.id,m.member_id,m.member_fname,m.member_lname,m.member_country_region,m.member_city,m.member_email_address,m.member_mobile,t.station_id,t.date_created,t.date_updated FROM wp_members AS m INNER JOIN transaction as T ON t.member_id=m.member_id").Scan(&query)
			c.JSON(http.StatusOK, &query)
		} else {
			respond(http.StatusBadRequest,"Station not found!",c,true)
		}
	} else {
		respond(http.StatusBadRequest,"Please supply the station id!",c,true)
	}
}


