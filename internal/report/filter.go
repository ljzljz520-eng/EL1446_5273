package report

import (
	"example.com/inspection14/internal/domain"
	"strings"
)

func FilterPhotos(ps []domain.Photo, cat string) []domain.Photo {
	out := []domain.Photo{}
	for _, p := range ps {
		if cat == "" || p.Category == cat {
			out = append(out, p)
		}
	}
	return out
}
func SearchCaptions(ps []domain.Photo, term string) []domain.Photo {
	term = strings.ToLower(term)
	out := []domain.Photo{}
	for _, p := range ps {
		if strings.Contains(strings.ToLower(p.Caption), term) {
			out = append(out, p)
		}
	}
	return out
}
func SortBySequence(ps []domain.Photo) []domain.Photo {
	out := append([]domain.Photo{}, ps...)
	for i := 1; i < len(out); i++ {
		for j := i; j > 0 && out[j].Sequence < out[j-1].Sequence; j-- {
			out[j], out[j-1] = out[j-1], out[j]
		}
	}
	return out
}
func UniqueCategories(ps []domain.Photo) []string {
	m := map[string]bool{}
	out := []string{}
	for _, p := range ps {
		if !m[p.Category] {
			m[p.Category] = true
			out = append(out, p.Category)
		}
	}
	return out
}
func AcceptedOnly(ps []domain.Photo) []domain.Photo {
	out := []domain.Photo{}
	for _, p := range ps {
		if p.Accepted {
			out = append(out, p)
		}
	}
	return out
}
func CaptionRequired(p domain.Photo) bool { return p.Caption != "" || p.Category == "engine" }
