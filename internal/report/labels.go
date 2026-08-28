package report

var ExteriorLabels = []string{
	"front-bumper", "rear-bumper", "hood", "roof", "left-front-door", "right-front-door", "left-rear-door", "right-rear-door", "left-fender", "right-fender", "left-quarter", "right-quarter", "windshield", "rear-window", "left-mirror", "right-mirror", "front-grille", "headlamp-left", "headlamp-right", "taillamp-left", "taillamp-right", "license-front", "license-rear", "wheel-left-front", "wheel-right-front", "wheel-left-rear", "wheel-right-rear", "spare-wheel", "roof-rack", "sunroof", "trunk", "tailgate", "door-sill", "paint-depth", "panel-gap", "underbody", "frame", "exhaust", "catalyst", "fuel-door", "charging-port", "tow-hook", "antenna", "badge", "emblem", "handle-left-front", "handle-right-front", "handle-left-rear", "handle-right-rear", "glass-left-front", "glass-right-front", "glass-left-rear", "glass-right-rear", "wiper-front", "wiper-rear", "washer", "bumper-sensor", "parking-camera", "radar", "foglamp-left", "foglamp-right", "daylight-left", "daylight-right", "indicator-left", "indicator-right", "mirror-fold", "mirror-heater", "tire-tread-front-left", "tire-tread-front-right", "tire-tread-rear-left", "tire-tread-rear-right", "rim-front-left", "rim-front-right", "rim-rear-left", "rim-rear-right", "jack", "toolkit", "floor-mat", "cargo-cover", "trunk-light", "door-lock", "key-fob", "remote-start", "charging-cable", "inspection-plate", "odometer-photo", "vin-plate", "engine-bay", "coolant", "brake-fluid", "washer-fluid", "battery", "fuse-box", "air-filter", "oil-cap", "belts", "hoses", "radiator", "fan", "mounts", "shock-left", "shock-right", "brake-disc-left", "brake-disc-right", "brake-pad-left", "brake-pad-right", "suspension-left", "suspension-right", "steering-rack", "drive-shaft", "differential", "transmission", "leak-check", "rust-check", "corrosion-check", "seal-check", "weld-check", "alignment", "road-surface", "parking-context", "seller-context", "lighting-context", "distance-shot", "closeup-shot", "wide-shot", "detail-shot", "damage-marker", "repair-marker", "scrape-marker", "dent-marker", "crack-marker", "chip-marker", "stain-marker", "mismatch-marker", "color-marker", "reflection-marker", "measurement-marker", "scale-marker", "timestamp-marker", "gps-marker", "inspector-badge", "vehicle-side", "vehicle-front", "vehicle-rear", "vehicle-top", "vehicle-bottom", "vehicle-left", "vehicle-right", "vehicle-corner", "vehicle-cabin", "vehicle-cargo", "vehicle-dashboard", "vehicle-console", "vehicle-seat", "vehicle-pedal", "vehicle-airbag", "vehicle-belt", "vehicle-window", "vehicle-mirror", "vehicle-light", "vehicle-switch", "vehicle-screen", "vehicle-speaker", "vehicle-camera", "vehicle-sensor", "vehicle-connector", "vehicle-port", "vehicle-label", "vehicle-manual", "vehicle-record", "vehicle-service", "vehicle-tire", "vehicle-wheel", "vehicle-brake", "vehicle-fluid", "vehicle-engine", "vehicle-transmission", "vehicle-frame", "vehicle-structure", "vehicle-safety", "vehicle-accident", "vehicle-modification", "vehicle-accessory", "vehicle-document", "vehicle-owner", "vehicle-matching", "vehicle-confirmation", "vehicle-signoff", "vehicle-final", "vehicle-summary", "vehicle-overview", "vehicle-evidence", "vehicle-quality", "vehicle-review", "vehicle-approval", "vehicle-archive",
}

func LabelCount() int { return len(ExteriorLabels) }
func HasLabel(label string) bool {
	for _, v := range ExteriorLabels {
		if v == label {
			return true
		}
	}
	return false
}
func LabelAt(i int) string {
	if i < 0 || i >= len(ExteriorLabels) {
		return ""
	}
	return ExteriorLabels[i]
}
