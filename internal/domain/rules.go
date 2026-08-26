package domain

import "fmt"

func ValidTransition(from, to InspectionStatus) bool {
	switch from {
	case StatusDraft:
		return to == StatusSubmitted || to == StatusRejected
	case StatusSubmitted:
		return to == StatusApproved || to == StatusRejected
	case StatusApproved:
		return to == StatusArchived
	case StatusRejected, StatusArchived:
		return false
	}
	return false
}
func ValidateVehicle(v Vehicle) error {
	if v.ID == "" || v.VIN == "" {
		return fmt.Errorf("vehicle identity required")
	}
	if len(v.VIN) < 8 {
		return fmt.Errorf("vin too short")
	}
	if v.Year < 1950 || v.Year > 2100 {
		return fmt.Errorf("year out of range")
	}
	if v.Mileage < 0 {
		return fmt.Errorf("mileage invalid")
	}
	return nil
}
func ValidateInspection(i Inspection) error {
	if i.ID == "" || i.VehicleID == "" || i.Inspector == "" {
		return fmt.Errorf("inspection fields required")
	}
	return nil
}
func ValidatePhoto(p Photo) error {
	if p.ID == "" || p.InspectionID == "" || p.URI == "" {
		return fmt.Errorf("photo fields required")
	}
	if p.Sequence < 0 {
		return fmt.Errorf("sequence invalid")
	}
	return nil
}
