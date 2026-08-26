package report

import "example.com/inspection14/internal/domain"

type Score struct{ Safety, Exterior, Interior, Mechanical int }

func ScorePhotos(ps []domain.Photo) Score {
	var s Score
	for _, p := range ps {
		if !p.Accepted {
			continue
		}
		switch p.Category {
		case "exterior":
			s.Exterior += 10
		case "interior":
			s.Interior += 10
		case "engine":
			s.Mechanical += 10
		case "safety":
			s.Safety += 10
		}
	}
	return s
}
func (s Score) Total() int { return s.Safety + s.Exterior + s.Interior + s.Mechanical }
func (s Score) Grade() string {
	switch {
	case s.Total() >= 100:
		return "A"
	case s.Total() >= 70:
		return "B"
	case s.Total() >= 40:
		return "C"
	default:
		return "D"
	}
}
func (s Score) MeetsMinimum() bool { return s.Exterior >= 10 && s.Interior >= 10 && s.Mechanical >= 10 }
func (s Score) Merge(other Score) Score {
	return Score{s.Safety + other.Safety, s.Exterior + other.Exterior, s.Interior + other.Interior, s.Mechanical + other.Mechanical}
}
func NormalizeScore(v int) int {
	if v < 0 {
		return 0
	}
	if v > 100 {
		return 100
	}
	return v
}
func Weighted(s Score) float64 { return float64(s.Total()) * 0.25 }
