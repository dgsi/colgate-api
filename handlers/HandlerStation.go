package handlers

import (
	"net/http"
	"time"
	
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	m "colgate/dgsi/api/models"
)

type StationHandler struct {
	db *gorm.DB
}

func NewStationHandler(db *gorm.DB) *StationHandler {
	return &StationHandler{db}
}

// get all stations
func (handler StationHandler) Index(c *gin.Context) {
	stations := []m.Station{}	
	handler.db.Table("station").Find(&stations)
	c.JSON(http.StatusOK, &stations)
}

// create new station
func (handler StationHandler) Create(c *gin.Context) {
	now := time.Now().UTC()
	station_id := c.PostForm("station_id")
	
	station := m.Station{}
	handler.db.Table("station").Where("station_id = ?",station_id).First(&station)

	if station.StationId != "" {
		respond(http.StatusBadRequest,"Station id already in used",c,true)	
	} else {
		handler.db.Exec("INSERT INTO station VALUES (?,?,?)",station_id,now,now)
		respond(http.StatusCreated,"New station successfully created!",c,false)	
	}
}

// update station
func (handler StationHandler) Update(c *gin.Context) {
	station_id := c.Param("station_id")
	new_station_id := c.PostForm("new_station_id")
	if new_station_id != "" {
		station := m.Station{}
		handler.db.Table("station").Where("station_id = ?",station_id).First(&station)

		if station.StationId != "" {
			existingStation := m.Station{}
			handler.db.Table("station").Where("station_id = ?",new_station_id).First(&existingStation)
			if existingStation.StationId != "" {
				respond(http.StatusBadRequest,"Your desired station id was already used!",c,true)
			} else {
				now := time.Now().UTC()
				handler.db.Exec("UPDATE station SET station_id = ?, date_updated = ? WHERE station_id = ?",new_station_id,now,station_id)
				respond(http.StatusOK,"Station successfully updated",c,false)		
			}
		} else {
			respond(http.StatusBadRequest,"Station not found!",c,true)	
		}		
	} else {
		respond(http.StatusBadRequest,"Invalid new station id",c,true)	
	}
}

func (handler StationHandler) Login(c *gin.Context) {
	station_id := c.PostForm("station_id")
	
	if station_id != "" {
		station := m.Station{}
		handler.db.Table("station").Where("station_id = ?",station_id).First(&station)

		if station.StationId != "" {
			c.JSON(http.StatusOK, station)
		} else {
			respond(http.StatusBadRequest,"Station not found!",c,true)	
		}
	} else {
		respond(http.StatusBadRequest,"Invalid station id",c,true)
	}
}



