package store

import (
	"encoding/json"
	"example.com/inspection14/internal/domain"
)

func (s *Store) Photos() ([]domain.Photo, error) {
	raw, e := s.List("photos")
	if e != nil {
		return nil, e
	}
	out := make([]domain.Photo, 0, len(raw))
	for _, b := range raw {
		var p domain.Photo
		if jsonErr := unmarshal(b, &p); jsonErr != nil {
			return nil, jsonErr
		}
		out = append(out, p)
	}
	return out, nil
}
func (s *Store) Reviews() ([]domain.Review, error) {
	raw, e := s.List("reviews")
	if e != nil {
		return nil, e
	}
	out := make([]domain.Review, 0, len(raw))
	for _, b := range raw {
		var p domain.Review
		if e := unmarshal(b, &p); e != nil {
			return nil, e
		}
		out = append(out, p)
	}
	return out, nil
}
func unmarshal(b []byte, v any) error { return json.Unmarshal(b, v) }
