package trade

type Trade struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	UserID    int       `gorm:"not null;index"`
	OrderID   int       `gorm:"not null;index"`
	Amount    float64   `gorm:"type:decimal(10,2);not null"`
	Direction Direction `gorm:"type:varchar(10);not null"` // Example values: "buy", "sell"
}

type Direction string

const (
	Buy  Direction = "buy"
	Sell Direction = "sell"
)
