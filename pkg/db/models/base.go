package models

import "gorm.io/gorm"

type BaseModel struct {
	gorm.Model // ID, CreatedAt, UpdatedAt, DeletedAt
}

type CreatedUpdatedByModel struct {
	CreatedBy *uint `json:"Created_By"`
	UpdatedBy *uint `json:"Updated_By"`
}

type EnabledModel struct {
	Enabled bool `json:"Enabled"`
}
