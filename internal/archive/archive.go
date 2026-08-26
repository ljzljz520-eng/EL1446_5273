package archive

import (
	"example.com/inspection14/internal/domain"
	"example.com/inspection14/internal/store"
	"fmt"
)

type Service struct{ Store *store.Store }

func New(s *store.Store) *Service { return &Service{Store: s} }
func (a *Service) Locate(id string) (string, error) {
	v, e := a.Store.Inspection(id)
	if e != nil {
		return "", e
	}
	if v.Status != domain.StatusArchived {
		return "", fmt.Errorf("not archived")
	}
	return "vault/" + id, nil
}
func (a *Service) Verify(id string) (bool, error) {
	loc, e := a.Locate(id)
	return e == nil && loc != "", e
}
func (a *Service) Receipt(id string) string { return fmt.Sprintf("ARCHIVE-%s", id) }
