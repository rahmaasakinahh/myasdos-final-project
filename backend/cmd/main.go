package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"

	"myasdos/internal/domain/entity"
	"myasdos/internal/domain/factory"
	"myasdos/internal/domain/state"
	"myasdos/internal/infrastructure/database"
	"myasdos/internal/infrastructure/repository"
)

func main() {
	log.Println("=== MyAsdos Backend - Design Pattern Demonstration ===")
	fmt.Println()

	// ============================================================
	// PATTERN 1: SINGLETON — Inisialisasi Database Connection
	// ============================================================
	log.Println("[Singleton] Initializing database connection...")
	db := database.GetDB()
	if db == nil {
		log.Println("Note: Database not configured. Using mock repositories for demo.")
	} else {
		log.Println("Database connection pool ready (Singleton initialized)")
	}
	fmt.Println()

	// ============================================================
	// PATTERN 3: REPOSITORY — Membuat lamaran baru
	// ============================================================
	log.Println("[Repository] Creating ApplicationRepository...")

	// Untuk demo tanpa DB, gunakan Mock Repository
	appRepo := repository.NewMockApplicationRepository()

	studentID := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	jobID := uuid.MustParse("22222222-2222-2222-2222-222222222222")
	applicationID := uuid.MustParse("33333333-3333-3333-3333-333333333333")

	app := &entity.Application{
		ID:           applicationID,
		JobPostingID: jobID,
		StudentID:    studentID,
		Reason:       "Saya ingin menjadi asdos karena ingin memperdalam pemahaman tentang materi kuliah.",
		Status:       "PENDING",
	}

	if err := appRepo.Create(app); err != nil {
		log.Fatalf("Failed to create application: %v", err)
	}
	log.Printf("Application created: ID=%s, Status=%s", app.ID, app.Status)
	fmt.Println()

	// ============================================================
	// PATTERN 4: STATE — State Machine untuk Lamaran
	// ============================================================
	log.Println("[State] Initializing Application State Machine...")

	// Buat context dengan state awal PENDING
	appCtx := state.NewApplicationStateContext(
		applicationID.String(),
		studentID.String(),
		jobID.String(),
	)
	log.Printf("Initial status: %s", appCtx.GetStatus())

	// Simulasi: Dosen melakukan REVIEW
	log.Println("Action: Lecturer reviews the application...")
	if err := appCtx.Review(); err != nil {
		log.Fatalf("Review failed: %v", err)
	}
	log.Printf("Status after review: %s", appCtx.GetStatus())

	// Simulasi: Dosen memutuskan ACCEPTED
	log.Println("Action: Lecturer decides ACCEPTED...")
	if err := appCtx.Decide("ACCEPTED"); err != nil {
		log.Fatalf("Decide failed: %v", err)
	}
	log.Printf("Status after decision: %s", appCtx.GetStatus())

	// Simulasi: Coba tarik lamaran (harus gagal karena sudah ACCEPTED)
	log.Println("Action: Student tries to withdraw (should fail)...")
	if err := appCtx.Withdraw(); err != nil {
		log.Printf("Expected error: %v", err)
	} else {
		log.Fatal("Withdraw should have failed!")
	}
	fmt.Println()

	// Update repository dengan status final
	appRepo.UpdateStatus(applicationID, appCtx.GetStatus(), "Selamat diterima!")
	log.Println("Application status persisted to repository")
	fmt.Println()

	// ============================================================
	// PATTERN 2: FACTORY METHOD — Membuat Notifikasi
	// ============================================================
	log.Println("[Factory Method] Creating notifications...")

	// Inisialisasi factory registry
	factoryRegistry := factory.NewNotificationFactoryRegistry()

	// Notifikasi 1: Ke mahasiswa — lamaran diterima
	notifData1 := factory.NotificationData{
		RecipientID:       studentID,
		RelatedEntityID:   &applicationID,
		RelatedEntityType: "APPLICATION",
		ActorName:         "Dr. Budi Santoso",
		CourseName:        "Pemrograman Web",
	}
	notif1, err := factoryRegistry.Create("APPLICATION_ACCEPTED", notifData1)
	if err != nil {
		log.Fatalf("Failed to create notification: %v", err)
	}
	log.Printf("Notification 1 created:")
	log.Printf("  Title: %s", notif1.Title)
	log.Printf("  Message: %s", notif1.Message)
	log.Printf("  Type: %s", notif1.Type)
	fmt.Println()

	// Notifikasi 2: Ke dosen — pelamar baru (simulasi scenario lain)
	lecturerID := uuid.MustParse("44444444-4444-4444-4444-444444444444")
	notifData2 := factory.NotificationData{
		RecipientID:       lecturerID,
		RelatedEntityID:   &applicationID,
		RelatedEntityType: "APPLICATION",
		ActorName:         "Ahmad Rizki",
		CourseName:        "Pemrograman Web",
	}
	notif2, err := factoryRegistry.Create("NEW_APPLICANT", notifData2)
	if err != nil {
		log.Fatalf("Failed to create notification: %v", err)
	}
	log.Printf("Notification 2 created:")
	log.Printf("  Title: %s", notif2.Title)
	log.Printf("  Message: %s", notif2.Message)
	log.Printf("  Type: %s", notif2.Type)
	fmt.Println()

	// Notifikasi 3: Laporan disetujui (simulasi M5)
	notifData3 := factory.NotificationData{
		RecipientID:       studentID,
		RelatedEntityID:   &applicationID,
		RelatedEntityType: "REPORT",
		ActorName:         "Dr. Budi Santoso",
		CourseName:        "Pemrograman Web",
		Period:            "Minggu 3",
	}
	notif3, err := factoryRegistry.Create("REPORT_APPROVED", notifData3)
	if err != nil {
		log.Fatalf("Failed to create notification: %v", err)
	}
	log.Printf("Notification 3 created:")
	log.Printf("  Title: %s", notif3.Title)
	log.Printf("  Message: %s", notif3.Message)
	log.Printf("  Type: %s", notif3.Type)
	fmt.Println()

	// ============================================================
	// RINGKASAN
	// ============================================================
	log.Println("=== Summary ===")
	log.Println("1. [Singleton] Database connection pool initialized once")
	log.Println("2. [Repository] Application persisted via repository abstraction")
	log.Println("3. [State] Application transitioned: PENDING → REVIEWED → ACCEPTED")
	log.Println("4. [Factory Method] 3 notifications created dynamically based on event type")
	log.Println("=== All 4 Design Patterns demonstrated successfully ===")
}
