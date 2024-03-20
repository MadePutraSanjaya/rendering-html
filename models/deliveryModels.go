package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Delivery struct {
	ID               uuid.UUID `gorm:"type:uuid;primaryKey"`
	RobotId          string
	VendorId         string
	TaskId           string
	TaskType         string
	DestinationPoint string
	gorm.Model
}
