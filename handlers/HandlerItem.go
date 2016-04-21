package handlers

import (
	"net/http"
	"time"
	
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	m "colgate/dgsi/api/models"
)

type ItemHandler struct {
	db *gorm.DB
}

func NewItemHandler(db *gorm.DB) *ItemHandler {
	return &ItemHandler{db}
}

// get all items
func (handler ItemHandler) Index(c *gin.Context) {
	items := []m.Item{}	
	handler.db.Table("item").Find(&items)
	c.JSON(http.StatusOK, &items)
}

// create new item
func (handler ItemHandler) Create(c *gin.Context) {
	now := time.Now().UTC()
	item_name := c.PostForm("item_name")
	
	if item_name == "" {
		respond(http.StatusBadRequest,"Please supply the item name!",c,true)
	} else {
		item := m.Item{}
		handler.db.Table("item").Where("item_name = ?",item_name).First(&item)
		if (item.ItemName == "") {
			handler.db.Exec("INSERT INTO item VALUES (null,?,?,?)",item_name,now,now)
			respond(http.StatusCreated,"New item successfully created!",c,false)		
		} else {
			respond(http.StatusBadRequest,"Item name alreay taken!",c,true)
		}
	}
}

// show an item by id
func(handler ItemHandler) Show(c *gin.Context) {
	id := c.Param("item_id")
	item := m.Item{}
	handler.db.Table("item").Where("id = ?",id).First(&item)

	if item.ItemName != "" {
		c.JSON(http.StatusOK, item)
	} else {
		respond(http.StatusBadRequest,"Item not found!",c,true)	
	}
}

//update item
func (handler ItemHandler) Update(c *gin.Context) {
	id := c.Param("item_id")
	new_item_name := c.PostForm("new_item_name")

	if id == "" {
		respond(http.StatusBadRequest,"Please supply the item id!",c,true)
	} else if new_item_name == "" {
		respond(http.StatusBadRequest,"Please supply the new item name!",c,true)
	} else {
		item := m.Item{}
		handler.db.Table("item").Where("id = ?",id).First(&item)

		if item.ItemName != "" {
			existingItem := m.Item{}
			handler.db.Table("item").Where("item_name = ?",new_item_name).First(&existingItem)
			if existingItem.ItemName != "" {
				respond(http.StatusBadRequest,"Your desired item name was already used!",c,true)
			} else {
				now := time.Now().UTC()
				handler.db.Exec("UPDATE item SET item_name = ?, date_updated = ? WHERE id = ?",new_item_name,now,id)
				respond(http.StatusOK,"Item successfully updated",c,false)		
			}
		} else {
			respond(http.StatusBadRequest,"Item not found!",c,true)	
		}
	}
}

