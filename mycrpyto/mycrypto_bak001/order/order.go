package order

type Order struct {
	ID           int     `gorm:"primaryKey;autoIncrement"`
	UserID       int     `gorm:"not null;index"`
	FromCurrency string  `gorm:"type:varchar(10);not null"`
	ToCurrency   string  `gorm:"type:varchar(10);not null"`
	Amount       float64 `gorm:"type:decimal(10,2);not null"`
	Status       string  `gorm:"type:varchar(20);not null"` // Example values: "pending", "completed", "canceled"
}
