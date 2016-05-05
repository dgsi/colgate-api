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
	//"github.com/gin-gonic/contrib/jwt"
)

func main() {
	db := *InitDB()
	router := gin.Default()
	LoadAPIRoutes(router, &db)
}

var (
	mysupersecretpassword = "unicornsAreAwesome"
)

func LoadAPIRoutes(r *gin.Engine, db *gorm.DB) {
	private := r.Group("/api/v1")
	public := r.Group("/api/v1")
	//private.Use(jwt.Auth(mysupersecretpassword))

	//manage members
	memberHandler := h.NewMemberHandler(db)
	private.GET("/members", memberHandler.Index)
	private.POST("/register", memberHandler.Create)
	private.GET("/members/search", memberHandler.Search)
	private.GET("/members/search/:member_id", memberHandler.SearchById)

	//manage stations
	stationHandler := h.NewStationHandler(db)
	private.GET("/stations", stationHandler.Index)
	private.POST("/stations", stationHandler.Create)
	private.PUT("/stations/:station_id", stationHandler.Update)
	public.POST("/stations/auth", stationHandler.Login)

	//manage transactions
	transactionsHandler := h.NewTransactionHandler(db)
	private.GET("/transactions", transactionsHandler.Index)
	private.POST("/transactions", transactionsHandler.Create)
	private.GET("/transactions/member/:member_id", transactionsHandler.ShowMemberTransactions)
	private.GET("/transactions/stations/:station_id/:tx_type", transactionsHandler.ShowStationTransactions)

	//manage items
	itemHandler := h.NewItemHandler(db)
	private.GET("/items", itemHandler.Index)
	private.GET("/items/:item_id", itemHandler.Show)
	private.POST("/items", itemHandler.Create)
	private.PUT("/items/:item_id", itemHandler.Update)

	//manage rewards
	rewardsHanlder := h.NewRewardHandler(db)
	private.GET("/rewards", rewardsHanlder.Index)
	private.POST("/rewards", rewardsHanlder.Create)
	private.GET("/rewards/:member_id", rewardsHanlder.GetRewardsByUser)

	var port = os.Getenv("PORT")
	if port == "" {
		port = "9000"
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