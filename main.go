package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"shopify-intern-demo/controller"
	"shopify-intern-demo/models"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env files")
	}

	r := gin.Default()
	err = r.SetTrustedProxies([]string{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//Database initialization
	db := initializeDb()
	//Migrations
	migrateDb(db)

	store := cookie.NewStore([]byte(os.Getenv("SESSION_COOKIE_KEY")))
	r.Use(sessions.Sessions("mysession", store))
	//loading html files
	r.LoadHTMLGlob("templates/*")
	//Controller setup
	ic := new(controller.InventoryController)
	ic.DB = db
	wc := new(controller.WarehouseLocationController)
	wc.DB = db

	//Routes setup
	r.GET("/", ic.Index)
	r.POST("/product", ic.Store)
	r.POST("/product/delete/:id", ic.Delete)
	r.POST("/product/update/:id", ic.Update)
	r.GET("/warehouse", wc.Index)
	r.POST("/warehouse", wc.Store)

	//Run app
	err = r.Run(":8080")
	if err != nil {
		panic("Error running server" + err.Error())
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func initializeDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("warehouse.sqlite"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func migrateDb(db *gorm.DB) {
	_ = db.AutoMigrate(models.Warehouse{})
	_ = db.AutoMigrate(models.Product{})
}
