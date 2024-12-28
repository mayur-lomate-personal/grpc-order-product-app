package v1

type Product struct {
	ID          int32   `gorm:"primaryKey"`
	Name        string  `gorm:"size:255;not null"`
	Description string  `gorm:"size:500"`
	Price       float64 `gorm:"not null"`
	Stock       int32   `gorm:"not null"`
}
