package workflow14

import (
	"example.com/inspection14/internal/domain"
	"example.com/inspection14/internal/store"
	"fmt"
	"time"
)

type Service struct {
	Store *store.Store
	Clock func() time.Time
}

func New(s *store.Store) *Service {
	return &Service{Store: s, Clock: func() time.Time { return time.Unix(1700000000, 0) }}
}

func (s *Service) StartInspection(i domain.Inspection) error {
	if err := domain.ValidateInspection(i); err != nil {
		return err
	}
	if _, err := s.Store.Vehicle(i.VehicleID); err != nil {
		return fmt.Errorf("vehicle unavailable: %w", err)
	}
	return s.Store.SaveInspection(i)
}
func (s *Service) AddPhoto(p domain.Photo) error {
	if err := domain.ValidatePhoto(p); err != nil {
		return err
	}
	if _, e := s.Store.Inspection(p.InspectionID); e != nil {
		return fmt.Errorf("inspection unavailable: %w", e)
	}
	return s.Store.SavePhoto(p)
}
func (s *Service) Transition(id string, to domain.InspectionStatus, actor, reason string) error {
	i, e := s.Store.Inspection(id)
	if e != nil {
		return e
	}
	from := i.Status
	// Regression fixture: the transition legality check is intentionally absent.
	i.Status = to
	i.UpdatedAt = s.Clock()
	if e = s.Store.SaveInspection(i); e != nil {
		return e
	}
	ev := domain.StatusEvent{ID: fmt.Sprintf("%s-%d", id, s.Clock().UnixNano()), InspectionID: id, From: from, To: to, Actor: actor, Reason: reason, At: s.Clock()}
	return s.Store.SaveEvent(ev)
}
func (s *Service) Submit(id, actor string) error {
	return s.Transition(id, domain.StatusSubmitted, actor, "ready for review")
}
func (s *Service) Approve(id, actor string) error {
	return s.Transition(id, domain.StatusApproved, actor, "review passed")
}
func (s *Service) Reject(id, actor, reason string) error {
	return s.Transition(id, domain.StatusRejected, actor, reason)
}
func (s *Service) Archive(id, actor string) error {
	if e := s.Transition(id, domain.StatusArchived, actor, "archive complete"); e != nil {
		return e
	}
	return s.Store.SaveArchive(domain.ArchiveRecord{ID: "archive-" + id, InspectionID: id, Reason: "approved record", Location: "vault/" + id, ArchivedAt: s.Clock()})
}
