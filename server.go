package main

import (
	"os"
	"fmt"
	"log"

	 "github.com/gin-gonic/gin"
	 _ "github.com/go-sql-driver/mysql"
	 h "colgate/dgsi/api/handlers"
	 "colgate/dgsi/api/config"
	"github.com/jinzhu/gorm"
)

func main() {
	db := *InitDB()
	router := gin.Default()
	LoadAPIRoutes(router, &db)
}

func LoadAPIRoutes(r *gin.Engine, db *gorm.DB) {
	public := r.Group("/api/v1")

	//manage stations
	stationHandler := h.NewStationHandler(db)
	public.GET("/stations", stationHandler.Index)
	public.POST("/station", stationHandler.Create)
	public.PUT("/station/:station_id", stationHandler.Update)
	public.POST("/station/auth", stationHandler.Login)

	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("PORT ---> ",port)
	r.Run(fmt.Sprintf(":%s", port))
}

func InitDB() *gorm.DB {
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.GetString("DB_USER"), config.GetString("DB_PASS"),
		config.GetString("DB_HOST"), config.GetString("DB_PORT"),
		config.GetString("DB_NAME"))
	log.Printf("\nDatabase URL: %s\n", dbURL)

	_db, err := gorm.Open("mysql", dbURL)
	if err != nil {
		panic(fmt.Sprintf("Error connecting to the database:  %s", err))
	}
	_db.DB()
	_db.LogMode(true)
	//_db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&m.Category{})
	_db.Set("gorm:table_options", "ENGINE=InnoDB")
	return _db
}

func GetPort() string {
    var port = os.Getenv("PORT")
    // Set a default port if there is nothing in the environment
    if port == "" {
        port = "8000"
        fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
    }
    fmt.Println("port -----> ", port)
    return ":" + port
}