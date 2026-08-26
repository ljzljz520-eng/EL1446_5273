package catalog

import (
	"example.com/inspection14/internal/domain"
	"example.com/inspection14/internal/store"
	"path/filepath"
	"testing"
)

func TestCatalogRegistersVehicle(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	c := New(s)
	if e := c.RegisterVehicle(domain.Vehicle{ID: "v", VIN: "vin12345678", Make: "m", Model: "x", Year: 2022}); e != nil {
		t.Fatal(e)
	}
	v, e := s.Vehicle("v")
	if e != nil || v.VIN != "vin12345678" {
		t.Fatal("not saved")
	}
}
