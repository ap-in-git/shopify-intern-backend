package controller

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"shopify-intern-demo/models"
)

type WarehouseLocationController struct {
	DB *gorm.DB
}

func (wc WarehouseLocationController) Index(c *gin.Context) {
	var errorItem string
	var success string
	session := sessions.Default(c)
	val := session.Flashes("error")
	if val != nil {
		errorItem = val[0].(string)
	}
	val = session.Flashes("success")
	if val != nil {
		success = val[0].(string)
	}

	_ = session.Save()

	var warehouses []models.Warehouse
	wc.DB.Find(&warehouses)
	c.HTML(http.StatusOK, "warehouse.tmpl", gin.H{
		"error":      errorItem,
		"success":    success,
		"warehouses": warehouses,
	})
}

func (wc WarehouseLocationController) Store(c *gin.Context) {
	type FormRequest struct {
		Name     string `form:"name" binding:"required"`
		Location string `form:"location" binding:"required"`
	}
	session := sessions.Default(c)

	var formRequest FormRequest
	err := c.ShouldBind(&formRequest)
	if err != nil {
		session.AddFlash(err.Error(), "error")
		c.Redirect(http.StatusMovedPermanently, "/warehouse")
		return
	}
	var warehouse models.Warehouse
	warehouse.Name = formRequest.Name
	warehouse.Location = formRequest.Location
	err = wc.DB.Save(&warehouse).Error
	if err != nil {
		fmt.Println(err.Error())
		session.AddFlash("Something went wrong", "error")
	} else {
		session.AddFlash("Warehouse created successfully", "success")
	}
	session.Save()
	c.Redirect(http.StatusMovedPermanently, "/warehouse")
	return
}
