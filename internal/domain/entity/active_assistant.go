package entity

import (
	"time"

	"github.com/google/uuid"
)

// ActiveAssistant merepresentasikan asdos yang sedang aktif mengajar (M4).
type ActiveAssistant struct {
	ID           uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	JobPostingID uuid.UUID  `json:"job_posting_id" gorm:"type:uuid;not null;index:idx_active_assistants_job_posting_id"`
	JobPosting   JobPosting `json:"job_posting" gorm:"foreignKey:JobPostingID"`
	StudentID    uuid.UUID  `json:"student_id" gorm:"type:uuid;not null;index:idx_active_assistants_student_id"`
	Student      Student    `json:"student" gorm:"foreignKey:StudentID;references:UserID"`
	SemesterID   uuid.UUID  `json:"semester_id" gorm:"type:uuid;not null;index:idx_active_assistants_semester_id"`
	Semester     Semester   `json:"semester" gorm:"foreignKey:SemesterID"`
	StartedAt    time.Time  `json:"started_at" gorm:"type:timestamptz;default:now()"`
	EndedAt      *time.Time `json:"ended_at,omitempty" gorm:"type:timestamptz"`
	IsActive     bool       `json:"is_active" gorm:"type:boolean;default:true"`
}

// TableName mengembalikan nama tabel di PostgreSQL.
func (ActiveAssistant) TableName() string {
	return "active_assistants"
}
