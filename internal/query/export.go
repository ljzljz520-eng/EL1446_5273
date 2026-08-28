package query

import (
	"encoding/json"
	"example.com/inspection14/internal/domain"
	"fmt"
)

func (q *Service) ExportManifest(id string) ([]byte, error) {
	i, e := q.Store.Inspection(id)
	if e != nil {
		return nil, e
	}
	p, e := q.InspectionPhotos(id)
	if e != nil {
		return nil, e
	}
	m := struct {
		Inspection domain.Inspection `json:"inspection"`
		Photos     []domain.Photo    `json:"photos"`
	}{i, p}
	return json.MarshalIndent(m, "", "  ")
}
func (q *Service) Label(i domain.Inspection) string {
	if i.Status == domain.StatusArchived {
		return "archived"
	}
	if i.Status == domain.StatusApproved {
		return "approved"
	}
	return fmt.Sprintf("%s:%s", i.ID, i.Status)
}
