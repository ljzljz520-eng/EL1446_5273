package archive

import (
	"example.com/inspection14/internal/domain"
	"example.com/inspection14/internal/store"
	"path/filepath"
	"testing"
)

func TestWorkflowArchive(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	s.SaveInspection(domain.Inspection{ID: "i", Status: domain.StatusArchived})
	a := New(s)
	ok, e := a.Verify("i")
	if e != nil || !ok {
		t.Fatal("archive missing")
	}
}
