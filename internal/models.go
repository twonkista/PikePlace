package main

import (
	"time"
)

type User struct {
	id             uint   `gorm:"primaryKey;autoIncrement"`
	user_name      string `gorm:"uniqueIndex"`
	user_password  string `gorm:"not null"`
	balance        float64
	strikes        int
	locked_balance float64
	Pools          []Pool  `gorm:"foreignKey:creator_id"`
	Wagers         []Wager `gorm:"foreignKey:user_id"`
}

type Pool struct {
	id          uint    `gorm:"primaryKey;autoIncrement"`
	title       string  `gorm:"not null"`
	description string  `gorm:"not null"`
	creator_id  uint    `gorm:"not null"`
	creator     User    `gorm:"foreignKey:creator_id"`
	status      string  `gorm:"not null"`   // e.g., "open", "closed", "resolved"
	outcomes    string  `gorm:"type:jsonb"` // JSON array of possible outcomes
	sl_total    float64 `gorm:"not null;default:0"`
	sw_total    float64 `gorm:"not null;default:0"`
	outcome     string  `gorm:"not null"` // e.g., "outcome1", "outcome2"
}

type Wager struct {
	id        uint      `gorm:"primaryKey;autoIncrement"`
	user_id   uint      `gorm:"not null"`
	pools_id  uint      `gorm:"not null"`
	amount    float64   `gorm:"not null"`
	vote      string    `gorm:"not null"` // e.g., "outcome1", "outcome2"
	status    string    `gorm:"not null"` // e.g., "active", "canceled", "settled"
	placed_at time.Time `gorm:"autoCreateTime"`
	User      User      `gorm:"foreignKey:user_id"`
}
