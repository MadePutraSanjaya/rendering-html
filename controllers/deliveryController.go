package controllers

import (
	"net/http"
	"renderin-html/initializers"
	"renderin-html/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var body struct {
	RobotId          string
	VendorId         string
	TaskId           string
	TaskType         string
	DestinationPoint string
}

func DeliverysCreate(c *gin.Context) {
	// Generate new UUID
	id := uuid.New()

	// create delivery
	c.Bind(&body)
	delivery := models.Delivery{ID: id, RobotId: body.RobotId, VendorId: body.VendorId, TaskId: body.TaskId, TaskType: body.TaskType, DestinationPoint: body.DestinationPoint}
	result := initializers.DB.Create(&delivery)

	if result.Error != nil {
		c.Status(400)
		return

	}

	// Response
	c.JSON(200, gin.H{
		"delivery": delivery,
	})
}

func DeliverysIndex(ctx *gin.Context) {
	// get data
	var deliverys []models.Delivery
	result := initializers.DB.Find(&deliverys)

	// send to html
	ctx.HTML(http.StatusOK, "index.delivery.html", gin.H{
		"deliverys": deliverys,
	})

	// Response
	if result.Error != nil {
		ctx.Status(400)
		return
	}

}

func DeliverysShow(ctx *gin.Context) {
	var delivery models.Delivery
	// get id
	if err := initializers.DB.Where("id = ?", ctx.Param("id")).First(&delivery).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Data not found!"})
		return
	}
	// get data
	ctx.HTML(http.StatusOK, "detail.delivery.html", gin.H{
		"id":                delivery.ID,
		"robot_id":          delivery.RobotId,
		"vendor_id":         delivery.VendorId,
		"task_id":           delivery.TaskId,
		"task_type":         delivery.TaskType,
		"destination_point": delivery.DestinationPoint,
	})

}

func DeliverysUpdate(c *gin.Context) {
	// get id
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid ID format"})
		return
	}

	// get data
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// find the delivery
	var delivery models.Delivery
	if err := initializers.DB.First(&delivery, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "delivery not found"})
		return
	}

	// update data
	delivery.RobotId = body.RobotId
	delivery.VendorId = body.VendorId
	delivery.TaskId = body.TaskId
	delivery.TaskType = body.TaskType
	delivery.DestinationPoint = body.DestinationPoint

	if err := initializers.DB.Save(&delivery).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update delivery"})
		return
	}

	// response
	c.JSON(http.StatusOK, gin.H{"delivery": delivery})
}

func DeliverysDelete(c *gin.Context) {
	// get id
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}

	// find the delivery
	var delivery models.Delivery
	if err := initializers.DB.First(&delivery, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "delivery not found"})
		return
	}

	// delete the delivery
	if err := initializers.DB.Delete(&delivery).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete delivery"})
		return
	}

	// response
	c.JSON(http.StatusOK, gin.H{"message": "Delivery successfully deleted"})
}
