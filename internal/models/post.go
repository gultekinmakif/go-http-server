package models

import (
	"time"

	"gorm.io/gorm"
)

// Post is the canonical content entity behind /posts/{slug}.
//
// Slug is left empty by Create handlers for now — slug derivation (from
// Title) is intentionally a TODO. When implemented, add a `uniqueIndex` on
// Slug and enforce non-empty in handlers.
type Post struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Slug  string `gorm:"type:text"  json:"slug"` // TODO uniqueIndex once slug derivation lands
	Title string `gorm:"not null"   json:"title"`
	Body  string `gorm:"type:text"  json:"body"`

	// Auto-managed by GORM.
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // soft delete — auto-filtered on queries
}
