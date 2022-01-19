package controller

import (
	"errors"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"shopify-intern-demo/models"
	"shopify-intern-demo/utils"
	"strconv"
)

type InventoryController struct {
	DB *gorm.DB
}

func (ic *InventoryController) Index(c *gin.Context) {
	sessionMessage := utils.GetSuccessErrorSession(c)
	var products []models.Product
	ic.DB.Preload("Warehouse").Find(&products)
	var warehouses []models.Warehouse
	ic.DB.Select("id,name").Find(&warehouses)
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"warehouses": warehouses,
		"error":      sessionMessage.Error,
		"success":    sessionMessage.Success,
		"products":   products,
	})
}

func (ic *InventoryController) Store(c *gin.Context) {
	session := sessions.Default(c)

	//Form validation
	type FormRequest struct {
		Warehouse uint    `form:"warehouse" binding:"required"`
		Name      string  `form:"name" binding:"required"`
		Unit      string  `form:"unit" binding:"required"`
		Price     float32 `form:"price"  binding:"required"`
		Quantity  float32 `form:"quantity" binding:"required"`
	}
	var formRequest FormRequest
	err := c.ShouldBind(&formRequest)
	if err != nil {
		session.AddFlash(err.Error(), "error")
		_ = session.Save()
		c.Redirect(http.StatusMovedPermanently, "/")
		return
	}

	//Check if warehouse is valid or not
	var warehouse models.Warehouse
	err = ic.DB.Find(&warehouse, formRequest.Warehouse).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		session.AddFlash("Invalid warehouse location", "error")
		_ = session.Save()
		return
	}

	//Store product to db
	product := models.Product{
		Name:        formRequest.Name,
		Price:       formRequest.Price,
		Quantity:    formRequest.Quantity,
		Unit:        formRequest.Unit,
		WarehouseId: formRequest.Warehouse,
		Warehouse:   warehouse,
	}
	err = ic.DB.Save(&product).Error
	if err != nil {
		fmt.Println(err.Error())
		session.AddFlash("Something went wrong", "error")
	} else {
		session.AddFlash("Product created successfully", "success")
	}

	_ = session.Save()
	c.Redirect(http.StatusMovedPermanently, "/")
	return
}

func (ic *InventoryController) Update(c *gin.Context) {
	type FormRequest struct {
		Warehouse uint    `form:"warehouse" binding:"required"`
		Name      string  `form:"name" binding:"required"`
		Unit      string  `form:"unit" binding:"required"`
		Price     float32 `form:"price"  binding:"required"`
		Quantity  float32 `form:"quantity" binding:"required"`
	}
	//Session initialization
	session := sessions.Default(c)

	//Form validation
	var formRequest FormRequest
	err := c.ShouldBind(&formRequest)
	if err != nil {
		session.AddFlash(err.Error(), "error")
		_ = session.Save()
		c.Redirect(http.StatusMovedPermanently, "/")
		return
	}

	//Check if the product is valid or not
	id, _ := strconv.Atoi(c.Param("id"))
	var product models.Product
	err = ic.DB.Find(&product, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		session.AddFlash("Product not found", "error")
		_ = session.Save()
		return
	}

	//Check if new warehouse is valid or not
	var warehouse models.Warehouse
	err = ic.DB.Find(&warehouse, formRequest.Warehouse).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		session.AddFlash("Invalid warehouse location", "error")
		_ = session.Save()
		return
	}

	//Saving on db
	product.Name = formRequest.Name
	product.Unit = formRequest.Unit
	product.Quantity = formRequest.Quantity
	product.Warehouse = warehouse
	product.Price = formRequest.Price
	err = ic.DB.Save(&product).Error

	if err != nil {
		fmt.Println(err.Error())
		session.AddFlash("Something went wrong", "error")
	} else {
		session.AddFlash("Product updated successfully", "success")
	}
	session.Save()
	c.Redirect(http.StatusMovedPermanently, "/")
	return
}

func (ic *InventoryController) Delete(c *gin.Context) {
	//check if product is valid or not
	id, _ := strconv.Atoi(c.Param("id"))
	session := sessions.Default(c)
	var product models.Product
	err := ic.DB.Find(&product, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		session.AddFlash("Product not found", "error")
		_ = session.Save()
		return
	}

	//Delete the product
	err = ic.DB.Delete(&product).Error

	if err != nil {
		fmt.Println(err.Error())
		session.AddFlash("Something went wrong", "error")
	} else {
		session.AddFlash("Product deleted successfully", "success")
	}
	_ = session.Save()
	c.Redirect(http.StatusMovedPermanently, "/")
	return
}
