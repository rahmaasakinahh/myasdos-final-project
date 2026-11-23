package repository

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"myasdos/internal/domain/entity"
	"myasdos/internal/infrastructure/database"
)

// ============================================================
// REPOSITORY PATTERN
// ============================================================

// ApplicationRepository adalah interface (kontrak) yang didefinisikan
// oleh Domain Layer. Infrastructure Layer menyediakan implementasi-nya.
type ApplicationRepository interface {
	Create(app *entity.Application) error
	GetByID(id uuid.UUID) (*entity.Application, error)
	GetByStudentID(studentID uuid.UUID) ([]entity.Application, error)
	UpdateStatus(id uuid.UUID, status string, notes string) error
}

// -----------------------------------------------------------
// Concrete Implementation: ApplicationRepositoryGORM
// Menggunakan GORM untuk query ke PostgreSQL.
// -----------------------------------------------------------
type ApplicationRepositoryGORM struct {
	db *gorm.DB
}

// NewApplicationRepository membuat instance repository GORM.
// Mengambil DB instance dari Singleton DatabaseConnectionManager.
func NewApplicationRepository() ApplicationRepository {
	return &ApplicationRepositoryGORM{
		db: database.GetDB(),
	}
}

func (r *ApplicationRepositoryGORM) Create(app *entity.Application) error {
	return r.db.Create(app).Error
}

func (r *ApplicationRepositoryGORM) GetByID(id uuid.UUID) (*entity.Application, error) {
	var app entity.Application
	if err := r.db.First(&app, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &app, nil
}

func (r *ApplicationRepositoryGORM) GetByStudentID(studentID uuid.UUID) ([]entity.Application, error) {
	var apps []entity.Application
	if err := r.db.Where("student_id = ?", studentID).Find(&apps).Error; err != nil {
		return nil, err
	}
	return apps, nil
}

func (r *ApplicationRepositoryGORM) UpdateStatus(id uuid.UUID, status string, notes string) error {
	updates := map[string]interface{}{
		"status": status,
	}
	if notes != "" {
		updates["lecturer_notes"] = notes
	}
	if status == "ACCEPTED" || status == "REJECTED" {
		now := time.Now()
		updates["decided_at"] = &now
	}

	result := r.db.Model(&entity.Application{}).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("application not found: %s", id)
	}
	return nil
}

// -----------------------------------------------------------
// Mock Implementation (untuk testing / development tanpa DB)
// -----------------------------------------------------------

type MockApplicationRepository struct {
	data map[uuid.UUID]*entity.Application
}

// NewMockApplicationRepository membuat mock repository in-memory.
func NewMockApplicationRepository() ApplicationRepository {
	return &MockApplicationRepository{
		data: make(map[uuid.UUID]*entity.Application),
	}
}

func (r *MockApplicationRepository) Create(app *entity.Application) error {
	r.data[app.ID] = app
	return nil
}

func (r *MockApplicationRepository) GetByID(id uuid.UUID) (*entity.Application, error) {
	if app, ok := r.data[id]; ok {
		return app, nil
	}
	return nil, fmt.Errorf("not found")
}

func (r *MockApplicationRepository) GetByStudentID(studentID uuid.UUID) ([]entity.Application, error) {
	var result []entity.Application
	for _, app := range r.data {
		if app.StudentID == studentID {
			result = append(result, *app)
		}
	}
	return result, nil
}

func (r *MockApplicationRepository) UpdateStatus(id uuid.UUID, status string, notes string) error {
	if app, ok := r.data[id]; ok {
		app.Status = status
		if notes != "" {
			app.LecturerNotes = notes
		}
		return nil
	}
	return fmt.Errorf("not found")
}
