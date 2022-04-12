package models

type Category struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required,min=8"`
	AdParms     []AdParms
}

type AdParms struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	Name       string `json:"name" binding:"required"`
	Type       string `json:"type" binding:"required"`
	CategoryID int    `json:"category_id" gorm:"foreignkey:CategoryID"`
}
