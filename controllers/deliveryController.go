package controllers

import (
	"renderin-html/initializers"
	"renderin-html/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func DeliverysCreate(c *gin.Context) {
	// Generate new UUID
	id := uuid.New()

	// get data
	var body struct {
		RobotId          string
		VendorId         string
		TaskId           string
		TaskType         string
		DestinationPoint string
	}

	c.Bind(&body)

	// create delivery
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

func DeliverysIndex(c *gin.Context) {
	// get data
	var deliverys []models.Delivery
	result := initializers.DB.Find(&deliverys)

	// Response
	c.JSON(200, gin.H{
		"deliverys": deliverys,
	})

	if result.Error != nil {
		c.Status(400)
		return
	}

}

func DeliverysShow(c *gin.Context) {
	// get id
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid ID format"})
		return
	}

	// get data
	var deliverys models.Delivery
	result := initializers.DB.Find(&deliverys, id)

	// Response
	c.JSON(200, gin.H{
		"deliverys": deliverys,
	})

	if result.Error != nil {
		c.Status(400)
		return
	}
}

func DeliverysUpdate(c *gin.Context) {
	// Generate new UUID
	uuid := uuid.New()

	// get id
	id := c.Param("id")

	// get data
	var body struct {
		RobotId          string
		VendorId         string
		TaskId           string
		TaskType         string
		DestinationPoint string
	}

	c.Bind(&body)

	// find the post where will updating
	var deliverys models.Delivery
	result := initializers.DB.Find(&deliverys, id)

	// update data
	initializers.DB.Model(&deliverys).Updates(models.Delivery{
		ID:    uuid,
		RobotId: body.RobotId,
		VendorId:  body.VendorId,
		TaskId:  body.TaskId,
		TaskType:  body.TaskType,
		DestinationPoint:  body.DestinationPoint,
	})

	// response
	c.JSON(200, gin.H{
		"deliverys": deliverys,
	})

	if result.Error != nil {
		c.Status(400)
		return
	}

}

func DeliverysDelete(c *gin.Context) {
	// get id
	id := c.Param("id")

	// find the delivery where will deleting
	var deliverys models.Delivery
	result := initializers.DB.Delete(&deliverys, id)

	// response
	c.JSON(200, gin.H{
		"message": "Delete Delivery Successfully",
	})

	if result.Error != nil {
		c.Status(400)
		return
	}
}
