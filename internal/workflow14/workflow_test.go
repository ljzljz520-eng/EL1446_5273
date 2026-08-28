package workflow14

import (
	"example.com/inspection14/internal/domain"
	"example.com/inspection14/internal/store"
	"path/filepath"
	"testing"
	"time"
)

func setup(t *testing.T) (*Service, *store.Store) {
	s, e := store.Open(filepath.Join(t.TempDir(), "x.db"))
	if e != nil {
		t.Fatal(e)
	}
	t.Cleanup(func() { s.Close() })
	w := New(s)
	w.Clock = func() time.Time { return time.Unix(1700000000, 0) }
	if e = s.SaveVehicle(domain.Vehicle{ID: "v1", VIN: "VIN12345678", Make: "A", Model: "B", Year: 2020}); e != nil {
		t.Fatal(e)
	}
	if e = w.StartInspection(domain.Inspection{ID: "i1", VehicleID: "v1", Inspector: "ins", Status: domain.StatusDraft}); e != nil {
		t.Fatal(e)
	}
	return w, s
}
func TestWorkflowRegisterReviewArchive(t *testing.T) {
	w, _ := setup(t)
	if e := w.AddPhoto(domain.Photo{ID: "p1", InspectionID: "i1", Category: "exterior", URI: "a"}); e != nil {
		t.Fatal(e)
	}
	if e := w.Submit("i1", "ins"); e != nil {
		t.Fatal(e)
	}
	if e := w.Review("i1", "rev", "approve", "ok"); e != nil {
		t.Fatal(e)
	}
	if e := w.Archive("i1", "arch"); e != nil {
		t.Fatal(e)
	}
}
