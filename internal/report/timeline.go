package report

import (
	"example.com/inspection14/internal/domain"
	"sort"
	"time"
)

type TimelineEntry struct {
	At           time.Time
	Label, Actor string
}

func Timeline(events []domain.StatusEvent) []TimelineEntry {
	out := make([]TimelineEntry, 0, len(events))
	for _, e := range events {
		out = append(out, TimelineEntry{e.At, string(e.To), e.Actor})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].At.Before(out[j].At) })
	return out
}
func Latest(events []domain.StatusEvent) domain.StatusEvent {
	if len(events) == 0 {
		return domain.StatusEvent{}
	}
	out := events[0]
	for _, e := range events[1:] {
		if e.At.After(out.At) {
			out = e
		}
	}
	return out
}
func Duration(events []domain.StatusEvent) time.Duration {
	if len(events) < 2 {
		return 0
	}
	first, last := events[0], events[0]
	for _, e := range events {
		if e.At.Before(first.At) {
			first = e
		}
		if e.At.After(last.At) {
			last = e
		}
	}
	return last.At.Sub(first.At)
}
func Actors(events []domain.StatusEvent) []string {
	m := map[string]bool{}
	out := []string{}
	for _, e := range events {
		if !m[e.Actor] {
			m[e.Actor] = true
			out = append(out, e.Actor)
		}
	}
	return out
}
func HasStatus(events []domain.StatusEvent, status domain.InspectionStatus) bool {
	for _, e := range events {
		if e.To == status {
			return true
		}
	}
	return false
}
func CountStatus(events []domain.StatusEvent, status domain.InspectionStatus) int {
	n := 0
	for _, e := range events {
		if e.To == status {
			n++
		}
	}
	return n
}
