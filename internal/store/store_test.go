package store

import (
	"example.com/inspection14/internal/domain"
	"path/filepath"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "persist.db")
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	v := domain.Vehicle{ID: "v", VIN: "VIN12345678", Make: "M", Model: "X", Year: 2021}
	if e = s.SaveVehicle(v); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	got, e := s.Vehicle("v")
	if e != nil || got.VIN != v.VIN {
		t.Fatal("vehicle not restored")
	}
}
