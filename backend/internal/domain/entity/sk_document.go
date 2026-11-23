package entity

import (
	"time"

	"github.com/google/uuid"
)

// SKDocument merepresentasikan Surat Keterangan (SK) digital untuk asdos (M6).
type SKDocument struct {
	ID                uuid.UUID       `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	ActiveAssistantID uuid.UUID       `json:"active_assistant_id" gorm:"type:uuid;not null;index:idx_sk_documents_active_assistant_id"`
	ActiveAssistant   ActiveAssistant `json:"active_assistant" gorm:"foreignKey:ActiveAssistantID"`
	FileURL           string          `json:"file_url" gorm:"type:varchar(500);not null"`
	IssuedBy          uuid.UUID       `json:"issued_by" gorm:"type:uuid;not null"`
	Admin             Admin           `json:"admin" gorm:"foreignKey:IssuedBy;references:UserID"`
	IssuedAt          time.Time       `json:"issued_at" gorm:"type:timestamptz;default:now()"`
}

// TableName mengembalikan nama tabel di PostgreSQL.
func (SKDocument) TableName() string {
	return "sk_documents"
}

// HonorariumReport merepresentasikan laporan pencairan honorarium (M7).
type HonorariumReport struct {
	ID                uuid.UUID       `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	ActiveAssistantID uuid.UUID       `json:"active_assistant_id" gorm:"type:uuid;not null;index:idx_honorarium_reports_active_assistant_id"`
	ActiveAssistant   ActiveAssistant `json:"active_assistant" gorm:"foreignKey:ActiveAssistantID"`
	Period            string          `json:"period" gorm:"type:varchar(50);not null"`
	FileURL           string          `json:"file_url" gorm:"type:varchar(500);not null"`
	UploadedBy        uuid.UUID       `json:"uploaded_by" gorm:"type:uuid;not null"`
	Admin             Admin           `json:"admin" gorm:"foreignKey:UploadedBy;references:UserID"`
	UploadedAt        time.Time       `json:"uploaded_at" gorm:"type:timestamptz;default:now()"`
}

// TableName mengembalikan nama tabel di PostgreSQL.
func (HonorariumReport) TableName() string {
	return "honorarium_reports"
}
