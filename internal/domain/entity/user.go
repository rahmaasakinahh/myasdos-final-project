package entity

import (
	"time"

	"github.com/google/uuid"
)

// User merepresentasikan data dasar semua pengguna (STUDENT, LECTURER, ADMIN).
type User struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	MyITSID     string    `json:"myits_id" gorm:"type:varchar(50);uniqueIndex;not null"`
	Email       string    `json:"email" gorm:"type:varchar(255);uniqueIndex;not null"`
	Name        string    `json:"name" gorm:"type:varchar(255);not null"`
	Role        string    `json:"role" gorm:"type:varchar(20);not null;index:idx_users_role"`       // STUDENT, LECTURER, ADMIN
	Status      string    `json:"status" gorm:"type:varchar(20);default:'ACTIVE';index:idx_users_status"` // ACTIVE, SUSPENDED
	AvatarURL   string    `json:"avatar_url" gorm:"type:varchar(500)"`
	PhoneNumber string    `json:"phone_number" gorm:"type:varchar(20)"`
	CreatedAt   time.Time `json:"created_at" gorm:"type:timestamptz;default:now()"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"type:timestamptz;default:now()"`
}

// TableName mengembalikan nama tabel di PostgreSQL.
func (User) TableName() string {
	return "users"
}

// Student merepresentasikan profil mahasiswa (extends User).
type Student struct {
	UserID       uuid.UUID  `json:"user_id" gorm:"type:uuid;primary_key"`
	User         User       `json:"user" gorm:"foreignKey:UserID"`
	NRP          string     `json:"nrp" gorm:"type:varchar(20);uniqueIndex;not null"`
	DepartmentID uuid.UUID  `json:"department_id" gorm:"type:uuid;not null;index:idx_students_department_id"`
	Department   Department `json:"department" gorm:"foreignKey:DepartmentID"`
	Bio          string     `json:"bio" gorm:"type:text"`
	PortfolioURL string     `json:"portfolio_url" gorm:"type:varchar(500)"`
	GPA          float64    `json:"gpa" gorm:"type:decimal(3,2)"`
	Semester     int        `json:"semester" gorm:"type:int"`
}

// TableName mengembalikan nama tabel di PostgreSQL.
func (Student) TableName() string {
	return "students"
}

// Lecturer merepresentasikan profil dosen (extends User).
type Lecturer struct {
	UserID       uuid.UUID  `json:"user_id" gorm:"type:uuid;primary_key"`
	User         User       `json:"user" gorm:"foreignKey:UserID"`
	NIP          string     `json:"nip" gorm:"type:varchar(20);uniqueIndex;not null"`
	DepartmentID uuid.UUID  `json:"department_id" gorm:"type:uuid;not null;index:idx_lecturers_department_id"`
	Department   Department `json:"department" gorm:"foreignKey:DepartmentID"`
}

// TableName mengembalikan nama tabel di PostgreSQL.
func (Lecturer) TableName() string {
	return "lecturers"
}

// Admin merepresentasikan profil admin (extends User).
type Admin struct {
	UserID     uuid.UUID `json:"user_id" gorm:"type:uuid;primary_key"`
	User       User      `json:"user" gorm:"foreignKey:UserID"`
	AdminLevel string    `json:"admin_level" gorm:"type:varchar(20);default:'REGULAR'"` // SUPER, REGULAR
}

// TableName mengembalikan nama tabel di PostgreSQL.
func (Admin) TableName() string {
	return "admins"
}
