package query

import (
	"example.com/inspection14/internal/domain"
	"example.com/inspection14/internal/store"
	"path/filepath"
	"testing"
)

func TestWorkflowQueryExport(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	s.SaveVehicle(domain.Vehicle{ID: "v", VIN: "VIN12345678", Make: "Toyota", Model: "A", Year: 2020})
	q := New(s)
	vs, e := q.FindVehicle("toy")
	if e != nil || len(vs) != 1 {
		t.Fatal("query failed")
	}
}
