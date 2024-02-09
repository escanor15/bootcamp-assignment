package controllers

import (
	"fmt"
	"go_bootcamp/H8-Assign/database"
	"go_bootcamp/H8-Assign/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var ItemDatas = []models.Item{}
var OrderDatas = []models.Order{}

func CreateItems(c *gin.Context) {
	var currentOrder models.Order

	c.ShouldBindJSON(&currentOrder)

	db := database.GetDB()
	err := db.Create(&currentOrder).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Can't create order",
			"error_detail":  err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"created_order": currentOrder,
	})
}

func GetAllItems(ctx *gin.Context) {
	var db = database.GetDB()
	if db == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
		return
	}

	var orders []models.Order

	// Fetch orders along with their associated items
	err := db.Preload("Items").Find(&orders).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Data": orders})
}

func UpdateItems(c *gin.Context) {
	orderID := c.Param("id")
	var db = database.GetDB()

	// Parse request body
	var updateData struct {
		CustomerName string `json:"customerName"`
		Items        []struct {
			ItemCode    string `json:"itemCode"`
			Description string `json:"description"`
			Quantity    int    `json:"quantity"`
		} `json:"items"`
	}

	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	// Fetch the order
	var order models.Order
	if err := db.Preload("Items").First(&order, "order_id = ?", orderID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Update order details
	order.CustomerName = updateData.CustomerName

	// Update items
	for _, updatedItem := range updateData.Items {
		// Check if item exists in order
		var itemExists bool
		var existingItem models.Item
		for _, item := range order.Items {
			if item.ItemCode == updatedItem.ItemCode {
				itemExists = true
				existingItem = item
				break
			}
		}

		if itemExists {
			// Update existing item
			existingItem.Description = updatedItem.Description
			existingItem.Quantity = uint(updatedItem.Quantity)
			if err := db.Save(&existingItem).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item"})
				return
			}
		} else {
			// Create new item
			newItem := models.Item{
				ItemCode:    updatedItem.ItemCode,
				Description: updatedItem.Description,
				Quantity:    uint(updatedItem.Quantity),
				OrderID:     order.OrderID,
			}
			if err := db.Create(&newItem).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create item"})
				return
			}
		}
	}

	// Save the updated order
	if err := db.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Order with id %v has been successfully updated", orderID)})
}

func DeleteOrder(ctx *gin.Context) {
	var db = database.GetDB()

	var order models.Order
	if err := db.First(&order, "order_id = ?", ctx.Param("id")).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Order not found!"})
		return
	}

	// Delete associated items first
	if err := db.Where("order_id = ?", order.OrderID).Delete(&models.Item{}).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete associated items"})
		return
	}

	// Then delete the order
	if err := db.Delete(&order).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
