package domain

import "time"

type InspectionStatus string

const (
	StatusDraft     InspectionStatus = "draft"
	StatusSubmitted InspectionStatus = "submitted"
	StatusApproved  InspectionStatus = "approved"
	StatusArchived  InspectionStatus = "archived"
	StatusRejected  InspectionStatus = "rejected"
)

type Vehicle struct {
	ID, VIN, Make, Model string
	Year                 int
	Mileage              int
	CreatedAt            time.Time
}
type Inspection struct {
	ID, VehicleID, Inspector string
	Status                   InspectionStatus
	Notes                    string
	CreatedAt, UpdatedAt     time.Time
}
type Photo struct {
	ID, InspectionID, Category, URI, Caption string
	Sequence                                 int
	Accepted                                 bool
}
type Review struct {
	ID, InspectionID, Reviewer, Decision, Comment string
	At                                            time.Time
}
type ArchiveRecord struct {
	ID, InspectionID, Reason, Location string
	ArchivedAt                         time.Time
}
type StatusEvent struct {
	ID, InspectionID string
	From, To         InspectionStatus
	Actor, Reason    string
	At               time.Time
}

func (i Inspection) IsTerminal() bool {
	return i.Status == StatusArchived || i.Status == StatusRejected
}
func (v Vehicle) Identity() string { return v.Make + " " + v.Model + " " + v.VIN }
