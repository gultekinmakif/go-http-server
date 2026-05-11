package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostContent struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Slug  string `json:"slug"`
}

// Post is the canonical content entity behind /posts/{slug}.
type Post struct {
	ID uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	PostContent

	// Auto-managed by GORM.
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // soft delete - auto-filtered on queries
}

func (P *Post) Sanitize() PostContent {
	return P.PostContent
}
