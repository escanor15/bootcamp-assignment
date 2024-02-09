package models

import "time"

type Item struct {
	ItemID      uint      `gorm:"primaryKey" json:"id"`
	ItemCode    string    `gorm:"not null;type:varchar(255)" json:"itemcode"`
	Description string    `gorm:"not null;type:varchar(255)" json:"Items"`
	Quantity    uint      `gorm:"not null;" json:"quantity"`
	OrderID     uint      `gorm:"not null;" json:"orderid"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
