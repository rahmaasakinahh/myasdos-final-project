package state

import (
	"fmt"
)

// ============================================================
// STATE PATTERN: Interface & Concrete States
// ============================================================

// ApplicationState mendefinisikan perilaku lamaran
// berdasarkan statusnya saat ini.
//
// Setiap concrete state akan implement method ini dengan
// return value yang berbeda-beda.
type ApplicationState interface {
	CanWithdraw() bool
	CanDecide() bool
	CanReview() bool
	GetStatus() string
	GetNextStates() []string
}

// -----------------------------------------------------------
// Concrete State 1: PENDING
// Mahasiswa baru submit. Dosen belum lihat.
// Boleh: tarik, review, decide
// -----------------------------------------------------------
type PendingState struct{}

func (s *PendingState) CanWithdraw() bool   { return true }
func (s *PendingState) CanDecide() bool      { return true }
func (s *PendingState) CanReview() bool      { return true }
func (s *PendingState) GetStatus() string    { return "PENDING" }
func (s *PendingState) GetNextStates() []string {
	return []string{"REVIEWED", "ACCEPTED", "REJECTED", "WITHDRAWN"}
}

// -----------------------------------------------------------
// Concrete State 2: REVIEWED
// Dosen sudah lihat, belum putuskan.
// Boleh: tarik, decide
// Tidak boleh: review lagi (sudah di-review)
// -----------------------------------------------------------
type ReviewedState struct{}

func (s *ReviewedState) CanWithdraw() bool   { return true }
func (s *ReviewedState) CanDecide() bool     { return true }
func (s *ReviewedState) CanReview() bool     { return false }
func (s *ReviewedState) GetStatus() string   { return "REVIEWED" }
func (s *ReviewedState) GetNextStates() []string {
	return []string{"ACCEPTED", "REJECTED", "WITHDRAWN"}
}

// -----------------------------------------------------------
// Concrete State 3: ACCEPTED
// Lamaran diterima. Status final.
// Tidak boleh: tarik, decide, review
// -----------------------------------------------------------
type AcceptedState struct{}

func (s *AcceptedState) CanWithdraw() bool  { return false }
func (s *AcceptedState) CanDecide() bool     { return false }
func (s *AcceptedState) CanReview() bool    { return false }
func (s *AcceptedState) GetStatus() string   { return "ACCEPTED" }
func (s *AcceptedState) GetNextStates() []string {
	return []string{} // terminal state
}

// -----------------------------------------------------------
// Concrete State 4: REJECTED
// Lamaran ditolak. Status final.
// Tidak boleh: tarik, decide, review
// -----------------------------------------------------------
type RejectedState struct{}

func (s *RejectedState) CanWithdraw() bool  { return false }
func (s *RejectedState) CanDecide() bool     { return false }
func (s *RejectedState) CanReview() bool    { return false }
func (s *RejectedState) GetStatus() string   { return "REJECTED" }
func (s *RejectedState) GetNextStates() []string {
	return []string{} // terminal state
}

// -----------------------------------------------------------
// Concrete State 5: WITHDRAWN
// Mahasiswa menarik lamaran. Status final.
// Tidak boleh: tarik, decide, review
// -----------------------------------------------------------
type WithdrawnState struct{}

func (s *WithdrawnState) CanWithdraw() bool  { return false }
func (s *WithdrawnState) CanDecide() bool     { return false }
func (s *WithdrawnState) CanReview() bool    { return false }
func (s *WithdrawnState) GetStatus() string   { return "WITHDRAWN" }
func (s *WithdrawnState) GetNextStates() []string {
	return []string{} // terminal state
}

// ============================================================
// STATE FACTORY: Mengkonversi string status → objek State
// ============================================================

// StateFactory membuat objek concrete state dari string status.
// Ini adalah helper function, bukan design pattern tersendiri.
func StateFactory(status string) (ApplicationState, error) {
	switch status {
	case "PENDING":
		return &PendingState{}, nil
	case "REVIEWED":
		return &ReviewedState{}, nil
	case "ACCEPTED":
		return &AcceptedState{}, nil
	case "REJECTED":
		return &RejectedState{}, nil
	case "WITHDRAWN":
		return &WithdrawnState{}, nil
	default:
		return nil, fmt.Errorf("unknown application status: %s", status)
	}
}

// ============================================================
// APPLICATION (Context dalam State Pattern)
// ============================================================

// ApplicationStateContext adalah struct yang "membungkus"
// entity Application dan menambahkan behavior state machine.
//
// Dalam terminologi GoF: ini adalah "Context".
type ApplicationStateContext struct {
	ID        string
	StudentID string
	JobID     string
	state     ApplicationState
}

// NewApplicationStateContext membuat context dengan state awal PENDING.
func NewApplicationStateContext(id, studentID, jobID string) *ApplicationStateContext {
	return &ApplicationStateContext{
		ID:        id,
		StudentID: studentID,
		JobID:     jobID,
		state:     &PendingState{},
	}
}

// SetState mengubah state internal (dipanggil setelah validasi).
func (ctx *ApplicationStateContext) SetState(state ApplicationState) {
	ctx.state = state
}

// GetStatus mengembalikan status string saat ini.
func (ctx *ApplicationStateContext) GetStatus() string {
	return ctx.state.GetStatus()
}

// Withdraw memvalidasi apakah lamaran boleh ditarik.
// Jika boleh, state berubah menjadi WithdrawnState.
func (ctx *ApplicationStateContext) Withdraw() error {
	if !ctx.state.CanWithdraw() {
		return fmt.Errorf("cannot withdraw application in %s state", ctx.GetStatus())
	}
	ctx.SetState(&WithdrawnState{})
	return nil
}

// Decide memvalidasi apakah dosen boleh memutuskan lamaran.
// Jika boleh, state berubah ke AcceptedState atau RejectedState.
func (ctx *ApplicationStateContext) Decide(decision string) error {
	if !ctx.state.CanDecide() {
		return fmt.Errorf("cannot decide application in %s state", ctx.GetStatus())
	}

	switch decision {
	case "ACCEPTED":
		ctx.SetState(&AcceptedState{})
	case "REJECTED":
		ctx.SetState(&RejectedState{})
	default:
		return fmt.Errorf("invalid decision: %s", decision)
	}
	return nil
}

// Review memvalidasi apakah lamaran boleh di-review.
// Jika boleh, state berubah menjadi ReviewedState.
func (ctx *ApplicationStateContext) Review() error {
	if !ctx.state.CanReview() {
		return fmt.Errorf("cannot review application in %s state", ctx.GetStatus())
	}
	ctx.SetState(&ReviewedState{})
	return nil
}
