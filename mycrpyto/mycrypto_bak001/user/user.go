package user

type User struct {
	ID       int     `gorm:"primaryKey;autoIncrement"`
	Username string  `gorm:"type:varchar(100);not null;unique"`
	Email    string  `gorm:"type:varchar(100);not null;unique"`
	Password string  `gorm:"type:varchar(100);not null"`
	Balance  float64 `gorm:"type:decimal(10,2);not null;default:0.0"`
}
