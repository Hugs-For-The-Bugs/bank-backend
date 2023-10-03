package model

type Transaction struct {
	ID            uint64 `gorm:"primaryKey"`
	FromAccountID string
	ToAccountID   string

	Amount string
}
