package models

import (
	"time"
)

type User struct {
	ID            uint   `gorm:"primaryKey;autoIncrement"`
	UserName      string `gorm:"uniqueIndex"`
	Password      string `gorm:"not null"`
	Balance       float64
	Strikes       int
	LockedBalance float64
	Pools         []Pool  `gorm:"foreignKey:creator_id"`
	Wagers        []Wager `gorm:"foreignKey:user_id"`
}

type Pool struct {
	ID          uint    `gorm:"primaryKey;autoIncrement"`
	Title       string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	CreatorID   uint    `gorm:"not null"`
	Creator     User    `gorm:"foreignKey:creator_id"`
	Status      string  `gorm:"not null"`   // e.g., "open", "closed", "resolved"
	Outcomes    string  `gorm:"type:jsonb"` // JSON array of possible outcomes
	SLTotal     float64 `gorm:"not null;default:0"`
	SWTotal     float64 `gorm:"not null;default:0"`
	Outcome     string  `gorm:"not null"` // e.g., "outcome1", "outcome2"
}

type Wager struct {
	ID       uint      `gorm:"primaryKey;autoIncrement"`
	UserID   uint      `gorm:"not null"`
	PoolID   uint      `gorm:"not null"`
	Amount   float64   `gorm:"not null"`
	Vote     string    `gorm:"not null"` // e.g., "outcome1", "outcome2"
	Status   string    `gorm:"not null"` // e.g., "active", "canceled", "settled"
	PlacedAt time.Time `gorm:"autoCreateTime"`
	User     User      `gorm:"foreignKey:user_id"`
}
