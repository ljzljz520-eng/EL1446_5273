package audit

import (
	"example.com/inspection14/internal/domain"
	"example.com/inspection14/internal/store"
	"path/filepath"
	"testing"
)

func TestAuditReportsRisk(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	s.SaveInspection(domain.Inspection{ID: "i", VehicleID: "v", Status: domain.StatusDraft})
	a := New(s)
	r, e := a.Risk(domain.Inspection{ID: "i", Status: domain.StatusDraft})
	if e != nil || r != "high" {
		t.Fatal(r, e)
	}
}
