package report

import (
	"example.com/inspection14/internal/domain"
	"fmt"
	"strings"
)

type Summary struct {
	Vehicle    domain.Vehicle
	Inspection domain.Inspection
	Photos     []domain.Photo
	Reviews    []domain.Review
	Warnings   []string
}

func NewSummary(v domain.Vehicle, i domain.Inspection) Summary {
	return Summary{Vehicle: v, Inspection: i, Photos: []domain.Photo{}, Reviews: []domain.Review{}, Warnings: []string{}}
}
func (s *Summary) AddPhoto(p domain.Photo)   { s.Photos = append(s.Photos, p) }
func (s *Summary) AddReview(r domain.Review) { s.Reviews = append(s.Reviews, r) }
func (s *Summary) AddWarning(w string) {
	if strings.TrimSpace(w) != "" {
		s.Warnings = append(s.Warnings, w)
	}
}
func (s Summary) PhotoTotal() int { return len(s.Photos) }
func (s Summary) AcceptedTotal() int {
	n := 0
	for _, p := range s.Photos {
		if p.Accepted {
			n++
		}
	}
	return n
}
func (s Summary) CategoryTotal(cat string) int {
	n := 0
	for _, p := range s.Photos {
		if p.Category == cat {
			n++
		}
	}
	return n
}
func (s Summary) ReviewOutcome() string {
	if len(s.Reviews) == 0 {
		return "pending"
	}
	return s.Reviews[len(s.Reviews)-1].Decision
}
func (s Summary) Ready() bool {
	return s.CategoryTotal("exterior") > 0 && s.CategoryTotal("interior") > 0 && s.CategoryTotal("engine") > 0
}
func (s Summary) Heading() string { return fmt.Sprintf("%s %s", s.Vehicle.Make, s.Vehicle.Model) }
func (s Summary) Lines() []string {
	return []string{s.Heading(), string(s.Inspection.Status), fmt.Sprintf("photos:%d", s.PhotoTotal()), fmt.Sprintf("reviews:%d", len(s.Reviews))}
}
func (s Summary) Copy() Summary {
	n := NewSummary(s.Vehicle, s.Inspection)
	n.Photos = append(n.Photos, s.Photos...)
	n.Reviews = append(n.Reviews, s.Reviews...)
	n.Warnings = append(n.Warnings, s.Warnings...)
	return n
}
