package database

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DatabaseConnectionManager adalah Singleton yang mengelola
// satu instance koneksi database PostgreSQL via GORM.
//
// Singleton memastikan bahwa seluruh aplikasi menggunakan
// connection pool yang sama, sehingga hemat resource.
type DatabaseConnectionManager struct {
	instance *gorm.DB
	once     sync.Once
}

// singletonInstance adalah instance global (private)
var singletonInstance *DatabaseConnectionManager

// getSingletonInstance mengembalikan pointer ke singleton manager.
// Menggunakan lazy initialization — baru dibuat saat pertama kali dipanggil.
func getSingletonInstance() *DatabaseConnectionManager {
	if singletonInstance == nil {
		singletonInstance = &DatabaseConnectionManager{}
	}
	return singletonInstance
}

// GetInstance mengembalikan *gorm.DB yang sudah terkoneksi ke PostgreSQL.
// sync.Once memastikan fungsi init hanya dijalankan satu kali meskipun
// dipanggil dari banyak goroutine secara concurrent.
func (d *DatabaseConnectionManager) GetInstance() *gorm.DB {
	d.once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			getEnv("DB_HOST", "localhost"),
			getEnv("DB_USER", "postgres"),
			getEnv("DB_PASSWORD", "postgres"),
			getEnv("DB_NAME", "myasdos"),
			getEnv("DB_PORT", "5432"),
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		// Connection pool configuration untuk performa
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Failed to get underlying sql.DB: %v", err)
		}
		sqlDB.SetMaxOpenConns(25)     // Maksimum 25 koneksi terbuka
		sqlDB.SetMaxIdleConns(10)     // Maksimum 10 koneksi idle
		sqlDB.SetConnMaxLifetime(300) // Koneksi di-recycle setiap 5 menit

		log.Println("Database connection established (Singleton initialized)")
		d.instance = db
	})

	return d.instance
}

// GetDB adalah helper function global untuk mendapatkan instance DB
// tanpa harus mengakses struct DatabaseConnectionManager langsung.
func GetDB() *gorm.DB {
	return getSingletonInstance().GetInstance()
}

// getEnv membaca environment variable dengan fallback value.
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
