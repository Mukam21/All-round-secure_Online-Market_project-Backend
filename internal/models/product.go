package models

type Product struct {
	ID            uint    `json:"id" gorm:"primaryKey"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	Description   string  `json:"description"`
	AgeRestricted bool    `json:"age_restricted" gorm:"default:false"`
}
