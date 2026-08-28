package workflow14

import (
	"example.com/inspection14/internal/domain"
	"testing"
)

func TestWorkflow14BusinessInvariant(t *testing.T) {
	w, s := setup(t)
	if e := w.Transition("i1", domain.StatusArchived, "bad", "illegal"); e == nil {
		t.Fatal("illegal transition accepted")
	}
	got, _ := s.Inspection("i1")
	if got.Status != domain.StatusDraft {
		t.Fatalf("status changed to %s", got.Status)
	}
}
func TestWorkflowPhotoCompleteness(t *testing.T) {
	w, _ := setup(t)
	for n, c := range []string{"exterior", "interior", "engine"} {
		if e := w.AddPhoto(domain.Photo{ID: string(rune('a' + n)), InspectionID: "i1", Category: c, URI: c}); e != nil {
			t.Fatal(e)
		}
	}
	ok, e := w.HasCategory("i1", "engine")
	if e != nil || !ok {
		t.Fatal("missing category")
	}
}
