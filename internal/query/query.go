package query

import (
	"encoding/json"
	"example.com/inspection14/internal/domain"
	"example.com/inspection14/internal/store"
	"sort"
	"strings"
)

type Service struct{ Store *store.Store }

func New(s *store.Store) *Service { return &Service{Store: s} }
func (q *Service) FindVehicle(term string) ([]domain.Vehicle, error) {
	raw, e := q.Store.List("vehicles")
	if e != nil {
		return nil, e
	}
	term = strings.ToLower(strings.TrimSpace(term))
	out := []domain.Vehicle{}
	for _, b := range raw {
		var v domain.Vehicle
		if e := unmarshal(b, &v); e != nil {
			return nil, e
		}
		if term == "" || strings.Contains(strings.ToLower(v.VIN), term) || strings.Contains(strings.ToLower(v.Make), term) || strings.Contains(strings.ToLower(v.Model), term) {
			out = append(out, v)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].CreatedAt.Before(out[j].CreatedAt) })
	return out, nil
}
func (q *Service) InspectionPhotos(id string) ([]domain.Photo, error) {
	ps, e := q.Store.Photos()
	if e != nil {
		return nil, e
	}
	out := []domain.Photo{}
	for _, p := range ps {
		if p.InspectionID == id {
			out = append(out, p)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Sequence < out[j].Sequence })
	return out, nil
}
func (q *Service) Status(id string) (domain.InspectionStatus, error) {
	i, e := q.Store.Inspection(id)
	return i.Status, e
}
func unmarshal(b []byte, v any) error { return json.Unmarshal(b, v) }
