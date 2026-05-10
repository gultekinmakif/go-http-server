package models

import "gorm.io/datatypes"

// Databank backs the /values/{key} endpoint with a flexible JSONB blob per key.
//
// Schema notes:
//   - Key is the natural primary key (TEXT).
//   - Data is Postgres JSONB — supports arbitrary JSON, partial updates, and
//     can be indexed with GIN if you ever need to query inside it.
type Databank struct {
	Key  string         `gorm:"primaryKey;type:text" json:"key"`
	Data datatypes.JSON `gorm:"type:jsonb"           json:"data"`
}

// TableName pins the table name. GORM's default pluralizer would turn
// "Databank" into "databanks" which is fine, but pin it explicitly so
// migrations don't break if you rename the type.
func (Databank) TableName() string {
	return "databank"
}
