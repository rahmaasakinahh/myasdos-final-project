package entity

import (
	"time"

	"github.com/google/uuid"
)

// Application merepresentasikan lamaran asdos (PRD M3).
type Application struct {
	ID              uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	JobPostingID    uuid.UUID  `json:"job_posting_id" gorm:"type:uuid;not null;index:idx_applications_job_posting_id"`
	JobPosting      JobPosting `json:"job_posting" gorm:"foreignKey:JobPostingID"`
	StudentID       uuid.UUID  `json:"student_id" gorm:"type:uuid;not null;index:idx_applications_student_id"`
	Student         Student    `json:"student" gorm:"foreignKey:StudentID;references:UserID"`
	Reason          string     `json:"reason" gorm:"type:text;not null"`
	CVURL           string     `json:"cv_url" gorm:"type:varchar(500)"`
	AdditionalData  string     `json:"additional_data" gorm:"type:jsonb"` // JSONB untuk data tambahan
	Status          string     `json:"status" gorm:"type:varchar(50);default:'PENDING';index:idx_applications_status"` // PENDING, REVIEWED, ACCEPTED, REJECTED, WITHDRAWN
	LecturerNotes   string     `json:"lecturer_notes" gorm:"type:text"`
	AppliedAt       time.Time  `json:"applied_at" gorm:"type:timestamptz;default:now()"`
	DecidedAt       *time.Time `json:"decided_at,omitempty" gorm:"type:timestamptz"`
}

// TableName mengembalikan nama tabel di PostgreSQL.
func (Application) TableName() string {
	return "applications"
}
