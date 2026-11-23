package factory

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"myasdos/internal/domain/entity"
)

// ============================================================
// FACTORY METHOD PATTERN
// ============================================================

// NotificationCreator adalah interface Factory Method.
// Setiap tipe notifikasi akan punya implementasi konkretnya sendiri.
type NotificationCreator interface {
	CreateNotification(data NotificationData) entity.Notification
}

// NotificationData berisi data mentah yang dibutuhkan
// oleh factory untuk membuat notifikasi.
type NotificationData struct {
	RecipientID       uuid.UUID
	RelatedEntityID   *uuid.UUID
	RelatedEntityType string
	ActorName         string   // Nama dosen/mahasiswa yang memicu event
	CourseName        string   // Nama mata kuliah terkait
	Period            string   // Periode laporan (opsional)
}

// -----------------------------------------------------------
// Concrete Factory 1: Notifikasi Lowongan Baru
// Trigger: Dosen membuat lowongan asdos
// Penerima: Mahasiswa di jurusan terkait (STUDENT)
// -----------------------------------------------------------
type NewJobPostingFactory struct{}

func (f *NewJobPostingFactory) CreateNotification(data NotificationData) entity.Notification {
	return entity.Notification{
		ID:                uuid.New(),
		UserID:            data.RecipientID,
		Type:              "NEW_JOB_POSTING",
		Title:             "Lowongan Asdos Baru",
		Message:           fmt.Sprintf("Lowongan baru: Asdos %s oleh %s. Segera daftar!", data.CourseName, data.ActorName),
		RelatedEntityID:   data.RelatedEntityID,
		RelatedEntityType: data.RelatedEntityType,
		IsRead:            false,
		CreatedAt:         time.Now(),
	}
}

// -----------------------------------------------------------
// Concrete Factory 2: Notifikasi Lamaran Diterima
// Trigger: Dosen menyetujui lamaran
// Penerima: Mahasiswa pelamar (STUDENT)
// -----------------------------------------------------------
type ApplicationAcceptedFactory struct{}

func (f *ApplicationAcceptedFactory) CreateNotification(data NotificationData) entity.Notification {
	return entity.Notification{
		ID:                uuid.New(),
		UserID:            data.RecipientID,
		Type:              "APPLICATION_ACCEPTED",
		Title:             "Lamaran Diterima",
		Message:           fmt.Sprintf("Selamat! Lamaran kamu untuk asdos %s telah diterima oleh %s.", data.CourseName, data.ActorName),
		RelatedEntityID:   data.RelatedEntityID,
		RelatedEntityType: data.RelatedEntityType,
		IsRead:            false,
		CreatedAt:         time.Now(),
	}
}

// -----------------------------------------------------------
// Concrete Factory 3: Notifikasi Laporan Disetujui
// Trigger: Dosen memverifikasi laporan dengan status APPROVED
// Penerima: Mahasiswa asdos (STUDENT)
// -----------------------------------------------------------
type ReportApprovedFactory struct{}

func (f *ReportApprovedFactory) CreateNotification(data NotificationData) entity.Notification {
	return entity.Notification{
		ID:                uuid.New(),
		UserID:            data.RecipientID,
		Type:              "REPORT_APPROVED",
		Title:             "Laporan Disetujui",
		Message:           fmt.Sprintf("Laporan periode %s untuk %s telah disetujui oleh %s.", data.Period, data.CourseName, data.ActorName),
		RelatedEntityID:   data.RelatedEntityID,
		RelatedEntityType: data.RelatedEntityType,
		IsRead:            false,
		CreatedAt:         time.Now(),
	}
}

// -----------------------------------------------------------
// Concrete Factory 4: Notifikasi Pelamar Baru
// Trigger: Mahasiswa submit lamaran ke lowongan
// Penerima: Dosen pemilik lowongan (LECTURER)
// -----------------------------------------------------------
type NewApplicantFactory struct{}

func (f *NewApplicantFactory) CreateNotification(data NotificationData) entity.Notification {
	return entity.Notification{
		ID:                uuid.New(),
		UserID:            data.RecipientID,
		Type:              "NEW_APPLICANT",
		Title:             "Pelamar Baru Masuk",
		Message:           fmt.Sprintf("%s mendaftar ke lowongan asdos %s.", data.ActorName, data.CourseName),
		RelatedEntityID:   data.RelatedEntityID,
		RelatedEntityType: data.RelatedEntityType,
		IsRead:            false,
		CreatedAt:         time.Now(),
	}
}

// ============================================================
// FACTORY REGISTRY (Client yang menggunakan Factory Method)
// ============================================================

// NotificationFactoryRegistry menyimpan mapping antara
// event type (string) dengan concrete factory-nya.
// Ini adalah "Factory of Factories" — memudahkan lookup runtime.
type NotificationFactoryRegistry struct {
	factories map[string]NotificationCreator
}

// NewNotificationFactoryRegistry membuat registry baru
// dan mendaftarkan semua concrete factory yang tersedia.
func NewNotificationFactoryRegistry() *NotificationFactoryRegistry {
	registry := &NotificationFactoryRegistry{
		factories: make(map[string]NotificationCreator),
	}

	// Daftarkan semua concrete factory
	registry.Register("NEW_JOB_POSTING", &NewJobPostingFactory{})
	registry.Register("APPLICATION_ACCEPTED", &ApplicationAcceptedFactory{})
	registry.Register("REPORT_APPROVED", &ReportApprovedFactory{})
	registry.Register("NEW_APPLICANT", &NewApplicantFactory{})

	return registry
}

// Register menambahkan concrete factory baru ke registry.
func (r *NotificationFactoryRegistry) Register(eventType string, factory NotificationCreator) {
	r.factories[eventType] = factory
}

// Create adalah entry point utama. Client cukup panggil ini
// dengan event type dan data, tanpa perlu tahu factory konkret mana yang dipakai.
func (r *NotificationFactoryRegistry) Create(eventType string, data NotificationData) (*entity.Notification, error) {
	factory, exists := r.factories[eventType]
	if !exists {
		return nil, fmt.Errorf("unknown notification event type: %s", eventType)
	}

	notification := factory.CreateNotification(data)
	return &notification, nil
}
