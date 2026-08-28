package catalog

import (
	"example.com/inspection14/internal/domain"
	"example.com/inspection14/internal/store"
	"fmt"
	"strings"
)

type Service struct{ Store *store.Store }

func New(s *store.Store) *Service { return &Service{Store: s} }
func (c *Service) RegisterVehicle(v domain.Vehicle) error {
	if err := domain.ValidateVehicle(v); err != nil {
		return err
	}
	return c.Store.SaveVehicle(v)
}
func (c *Service) StartInspection(i domain.Inspection) error {
	if err := domain.ValidateInspection(i); err != nil {
		return err
	}
	if _, e := c.Store.Vehicle(i.VehicleID); e != nil {
		return fmt.Errorf("vehicle unavailable: %w", e)
	}
	return c.Store.SaveInspection(i)
}
func (c *Service) NormalizeVIN(v string) string { return strings.ToUpper(strings.TrimSpace(v)) }
func (c *Service) VehicleSummary(v domain.Vehicle) string {
	return fmt.Sprintf("%s %d (%d km)", v.Identity(), v.Year, v.Mileage)
}
