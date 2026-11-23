package entity

import (
	"time"

	"github.com/google/uuid"
)

// JobPosting merepresentasikan lowongan asdos (PRD M2).
type JobPosting struct {
	ID           uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	LecturerID   uuid.UUID  `json:"lecturer_id" gorm:"type:uuid;not null;index:idx_job_postings_lecturer_id"`
	Lecturer     Lecturer   `json:"lecturer" gorm:"foreignKey:LecturerID;references:UserID"`
	CourseID     uuid.UUID  `json:"course_id" gorm:"type:uuid;not null;index:idx_job_postings_course_id"`
	Course       Course     `json:"course" gorm:"foreignKey:CourseID"`
	SemesterID   uuid.UUID  `json:"semester_id" gorm:"type:uuid;not null;index:idx_job_postings_semester_id"`
	Semester     Semester   `json:"semester" gorm:"foreignKey:SemesterID"`
	Description  string     `json:"description" gorm:"type:text"`
	Requirements string     `json:"requirements" gorm:"type:text"`
	Quota        int        `json:"quota" gorm:"type:int;not null"`
	Status       string     `json:"status" gorm:"type:varchar(20);default:'OPEN';index:idx_job_postings_status"` // OPEN, FULL, CLOSED
	OpenDate     time.Time  `json:"open_date" gorm:"type:date;not null"`
	CloseDate    time.Time  `json:"close_date" gorm:"type:date;not null"`
	Notes        string     `json:"notes" gorm:"type:text"`
	CreatedAt    time.Time  `json:"created_at" gorm:"type:timestamptz;default:now()"`
	UpdatedAt    time.Time  `json:"updated_at" gorm:"type:timestamptz;default:now()"`
}

// TableName mengembalikan nama tabel di PostgreSQL.
func (JobPosting) TableName() string {
	return "job_postings"
}
