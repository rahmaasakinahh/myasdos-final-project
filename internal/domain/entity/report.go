package entity

import (
	"time"

	"github.com/google/uuid"
)

// Report merepresentasikan laporan tanggung jawab atau rencana kerja asdos (M5).
type Report struct {
	ID                uuid.UUID       `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	ActiveAssistantID uuid.UUID       `json:"active_assistant_id" gorm:"type:uuid;not null;index:idx_reports_active_assistant_id"`
	ActiveAssistant   ActiveAssistant `json:"active_assistant" gorm:"foreignKey:ActiveAssistantID"`
	Type              string          `json:"type" gorm:"type:varchar(50);not null;index:idx_reports_type"` // TANGGUNG_JAWAB atau RENCANA_KERJA
	Period            string          `json:"period" gorm:"type:varchar(50);not null"`
	FileURL           string          `json:"file_url" gorm:"type:varchar(500);not null"`
	Status            string          `json:"status" gorm:"type:varchar(50);default:'SUBMITTED';index:idx_reports_status"` // SUBMITTED, APPROVED, REVISION, REJECTED
	Feedback          string          `json:"feedback" gorm:"type:text"`
	VerifiedFileURL   string          `json:"verified_file_url" gorm:"type:varchar(500)"`
	SubmittedAt       time.Time       `json:"submitted_at" gorm:"type:timestamptz;default:now()"`
	VerifiedAt        *time.Time      `json:"verified_at,omitempty" gorm:"type:timestamptz"`
}

// TableName mengembalikan nama tabel di PostgreSQL.
func (Report) TableName() string {
	return "reports"
}
