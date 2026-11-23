package entity

import (
	"time"

	"github.com/google/uuid"
)

// Department merepresentasikan jurusan/fakultas di ITS.
type Department struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Code      string    `json:"code" gorm:"type:varchar(10);uniqueIndex;not null"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Faculty   string    `json:"faculty" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:now()"`
}

// TableName mengembalikan nama tabel di PostgreSQL.
func (Department) TableName() string {
	return "departments"
}
