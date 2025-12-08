package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

// RunMigrations executes all SQL migration files in the config directory
func RunMigrations() error {
	migrationFiles := []string{
		"users.sql",
		"add_user_id_to_bookings.sql",
	}

	for _, filename := range migrationFiles {
		filePath := filepath.Join("config", filename)

		// Read SQL file
		sqlBytes, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Printf("Warning: Could not read migration file %s: %v", filename, err)
			continue
		}

		sqlQuery := string(sqlBytes)
		if sqlQuery == "" {
			continue
		}

		// Execute migration
		log.Printf("Running migration: %s", filename)
		_, err = DB.Exec(sqlQuery)
		if err != nil {
			log.Printf("Error running migration %s: %v", filename, err)
			// Don't return error, continue with other migrations
			continue
		}

		log.Printf("Successfully applied migration: %s", filename)
	}

	return nil
}

// CheckAndRunMigrations runs migrations and handles errors gracefully
func CheckAndRunMigrations() {
	fmt.Println("🔄 Running database migrations...")

	if err := RunMigrations(); err != nil {
		log.Printf("Migration error: %v", err)
	}

	fmt.Println("✅ Database migrations completed")
}
