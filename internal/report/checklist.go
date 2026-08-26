package report

import "example.com/inspection14/internal/domain"

type ChecklistItem struct {
	Key, Label         string
	Required, Complete bool
}

var DefaultChecklist = []ChecklistItem{{"exterior", "外观", true, false}, {"interior", "内饰", true, false}, {"engine", "发动机", true, false}, {"odometer", "里程", true, false}, {"documents", "证件", false, false}, {"tyres", "轮胎", false, false}}

func BuildChecklist(ps []domain.Photo) []ChecklistItem {
	out := make([]ChecklistItem, len(DefaultChecklist))
	copy(out, DefaultChecklist)
	for i := range out {
		for _, p := range ps {
			if p.Category == out[i].Key && p.URI != "" {
				out[i].Complete = true
			}
		}
	}
	return out
}
func Missing(items []ChecklistItem) []ChecklistItem {
	out := []ChecklistItem{}
	for _, i := range items {
		if i.Required && !i.Complete {
			out = append(out, i)
		}
	}
	return out
}
func Complete(items []ChecklistItem) bool {
	for _, i := range items {
		if i.Required && !i.Complete {
			return false
		}
	}
	return true
}
func Progress(items []ChecklistItem) float64 {
	if len(items) == 0 {
		return 1
	}
	n := 0
	for _, i := range items {
		if i.Complete {
			n++
		}
	}
	return float64(n) / float64(len(items))
}
func Keys(items []ChecklistItem) []string {
	out := []string{}
	for _, i := range items {
		out = append(out, i.Key)
	}
	return out
}
func Mark(items []ChecklistItem, key string) []ChecklistItem {
	out := make([]ChecklistItem, len(items))
	copy(out, items)
	for i := range out {
		if out[i].Key == key {
			out[i].Complete = true
		}
	}
	return out
}
