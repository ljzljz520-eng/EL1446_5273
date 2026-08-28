package workflow14

import (
	"example.com/inspection14/internal/domain"
	"fmt"
)

func (s *Service) Review(id, reviewer, decision, comment string) error {
	if decision != "approve" && decision != "reject" {
		return fmt.Errorf("unknown decision")
	}
	r := domain.Review{ID: fmt.Sprintf("review-%s-%s", id, reviewer), InspectionID: id, Reviewer: reviewer, Decision: decision, Comment: comment, At: s.Clock()}
	if e := s.Store.SaveReview(r); e != nil {
		return e
	}
	if decision == "approve" {
		return s.Approve(id, reviewer)
	}
	return s.Reject(id, reviewer, comment)
}
func (s *Service) PhotoCount(id string) (int, error) {
	ps, e := s.Store.Photos()
	if e != nil {
		return 0, e
	}
	n := 0
	for _, p := range ps {
		if p.InspectionID == id {
			n++
		}
	}
	return n, nil
}
func (s *Service) HasCategory(id, cat string) (bool, error) {
	ps, e := s.Store.Photos()
	if e != nil {
		return false, e
	}
	for _, p := range ps {
		if p.InspectionID == id && p.Category == cat {
			return true, nil
		}
	}
	return false, nil
}
