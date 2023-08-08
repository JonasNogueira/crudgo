package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name   string
	Period string
	Lab    string
}

type StudentResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deteledAt,omitempty"`
	Name      string    `json:"name"`
	Period    string    `json:"period"`
	Lab       string    `json:"lab"`
}
