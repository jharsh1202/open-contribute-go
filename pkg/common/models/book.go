package models

type Book struct {
	BaseModel
	CreatedUpdatedByModel
	EnabledModel
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}
