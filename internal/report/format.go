package report

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Row struct{ Label, Value string }

func Rows(s Summary) []Row {
	return []Row{{"VIN", s.Vehicle.VIN}, {"车型", s.Heading()}, {"状态", string(s.Inspection.Status)}, {"照片", fmt.Sprint(s.PhotoTotal())}, {"审核", s.ReviewOutcome()}}
}
func Plain(s Summary) string {
	parts := []string{}
	for _, r := range Rows(s) {
		parts = append(parts, r.Label+":"+r.Value)
	}
	return strings.Join(parts, " | ")
}
func JSON(s Summary) ([]byte, error) { return json.Marshal(s) }
func CSV(s Summary) string {
	lines := []string{"label,value"}
	for _, r := range Rows(s) {
		lines = append(lines, r.Label+","+r.Value)
	}
	return strings.Join(lines, "\n")
}
func StatusColor(status string) string {
	switch status {
	case "approved":
		return "green"
	case "rejected":
		return "red"
	case "archived":
		return "blue"
	default:
		return "yellow"
	}
}
func Badge(s Summary) string { return "[" + StatusColor(string(s.Inspection.Status)) + "]" }
func Escape(v string) string {
	return strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;", "\"", "&quot;").Replace(v)
}
func Table(s Summary) string {
	rows := Rows(s)
	out := []string{}
	for _, r := range rows {
		out = append(out, "<tr><th>"+Escape(r.Label)+"</th><td>"+Escape(r.Value)+"</td></tr>")
	}
	return strings.Join(out, "")
}
