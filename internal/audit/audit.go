package audit

import (
	"example.com/inspection14/internal/domain"
	"example.com/inspection14/internal/store"
	"fmt"
)

type Service struct{ Store *store.Store }

func New(s *store.Store) *Service                     { return &Service{Store: s} }
func (a *Service) Explain(i domain.Inspection) string { return fmt.Sprintf("%s is %s", i.ID, i.Status) }
func (a *Service) CheckCompleteness(id string) (bool, error) {
	ps, e := a.Store.Photos()
	if e != nil {
		return false, e
	}
	cats := map[string]bool{}
	for _, p := range ps {
		if p.InspectionID == id {
			cats[p.Category] = true
		}
	}
	return cats["exterior"] && cats["interior"] && cats["engine"], nil
}
func (a *Service) Risk(i domain.Inspection) (string, error) {
	ok, e := a.CheckCompleteness(i.ID)
	if e != nil {
		return "", e
	}
	if !ok {
		return "high", nil
	}
	if i.Status == domain.StatusRejected {
		return "blocked", nil
	}
	return "normal", nil
}
