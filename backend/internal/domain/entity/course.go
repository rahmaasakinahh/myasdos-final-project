package entity

import (
	"github.com/google/uuid"
)

// Course merepresentasikan mata kuliah di ITS.
type Course struct {
	ID           uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	DepartmentID uuid.UUID  `json:"department_id" gorm:"type:uuid;not null;index:idx_courses_department_id"`
	Department   Department `json:"department" gorm:"foreignKey:DepartmentID"`
	Code         string     `json:"code" gorm:"type:varchar(20);not null"`
	Name         string     `json:"name" gorm:"type:varchar(255);not null"`
	Credits      int        `json:"credits" gorm:"type:int;default:0"`
}

// TableName mengembalikan nama tabel di PostgreSQL.
func (Course) TableName() string {
	return "courses"
}
