package report

import "example.com/inspection14/internal/domain"

type Policy struct {
	RequiredCategories   []string
	MinimumPhotos        int
	RequireAccepted      bool
	AllowRejectedArchive bool
}

func DefaultPolicy() Policy {
	return Policy{[]string{"exterior", "interior", "engine"}, 3, true, false}
}
func (p Policy) Validate(s Summary) []string {
	issues := []string{}
	for _, c := range p.RequiredCategories {
		if s.CategoryTotal(c) == 0 {
			issues = append(issues, c)
		}
	}
	if s.PhotoTotal() < p.MinimumPhotos {
		issues = append(issues, "photo-count")
	}
	if p.RequireAccepted && s.AcceptedTotal() < p.MinimumPhotos {
		issues = append(issues, "accepted-count")
	}
	if !p.AllowRejectedArchive && s.Inspection.Status == domain.StatusRejected {
		issues = append(issues, "rejected")
	}
	return issues
}
func (p Policy) CanSubmit(s Summary) bool { return len(p.Validate(s)) == 0 }
func (p Policy) WithCategory(cat string) Policy {
	n := p
	n.RequiredCategories = append([]string{}, p.RequiredCategories...)
	n.RequiredCategories = append(n.RequiredCategories, cat)
	return n
}
func (p Policy) WithMinimum(n int) Policy { p.MinimumPhotos = n; return p }
func (p Policy) OptionalAccepted() Policy { p.RequireAccepted = false; return p }
