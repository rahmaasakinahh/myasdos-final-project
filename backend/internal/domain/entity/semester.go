package entity

import (
	"github.com/google/uuid"
)

// Semester merepresentasikan semester akademik.
type Semester struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	AcademicYear string    `json:"academic_year" gorm:"type:varchar(10);not null"`
	Term         string    `json:"term" gorm:"type:varchar(10);not null"` // GANJIL atau GENAP
	IsActive     bool      `json:"is_active" gorm:"type:boolean;default:false;index:idx_semesters_active"`
}

// TableName mengembalikan nama tabel di PostgreSQL.
func (Semester) TableName() string {
	return "semesters"
}
