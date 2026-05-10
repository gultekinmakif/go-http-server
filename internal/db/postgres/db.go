package postgres

import "gorm.io/gorm"

// db is the package-level GORM handle, initialized once via New() and accessed
// via Get(). Safe for concurrent use after init.
var db *gorm.DB

// New opens a Postgres connection via GORM, configures the pool, pings, and
// stores the handle in the package-level `db` variable.
func New(dsn string) error {

}

// Get returns the package-level handle. Panics if New() hasn't run yet —
// that's intentional: misuse is a programming error, not a runtime error.
func Get() *gorm.DB {

}

// Close shuts the underlying *sql.DB. Call on graceful shutdown.
func Close() error {

}

// Migrate runs schema migrations using GORM's AutoMigrate. Fine for templates
// and small projects; for production, swap to golang-migrate with versioned
// SQL files in /migrations.
func Migrate() error {
}
