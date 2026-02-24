package models

import (
	"time"
)

type User struct {
	ID             uint    `gorm:"primaryKey;autoIncrement"`
	UserName       string  `gorm:"uniqueIndex"`
	Role           string  `gorm:"not null;default:'sheep'"`
	HashedPassword string  `gorm:"column:password;not null" json:"-"`
	SessionToken   string  `gorm:"not null;default:''" json:"-"`
	CSRFToken      string  `gorm:"not null;default:''" json:"-"`
	Balance        float64 `gorm:"not null;default:1000"`
	Strikes        int
	LockedBalance  float64 `gorm:"not null;default:0"`
	Pools          []Pool  `gorm:"foreignKey:creator_id"`
	Wagers         []Wager `gorm:"foreignKey:user_id"`
}

type Pool struct {
	ID          uint    `gorm:"primaryKey;autoIncrement"`
	Title       string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	CreatorID   uint    `gorm:"not null"`
	Creator     User    `gorm:"foreignKey:creator_id"`
	Status      string  `gorm:"not null"` // e.g., "open", "closed", "resolved"
	SLTotal     float64 `gorm:"not null;default:0"`
	SWTotal     float64 `gorm:"not null;default:0"`
	MaxPoolSize float64 `gorm:"not null;default:100"`
	MaxWager    float64 `gorm:"not null;default:10"`
	Approved    string  `gorm:"not null;default:'pending'"` //e.g "approved", "pending", "rejected"
	ApprovedAt  time.Time
	Outcome     bool `gorm:"not null"` // e.g., "outcome1", "outcome2"
	CreatedAt   time.Time
	UpdatedAt   time.Time
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
