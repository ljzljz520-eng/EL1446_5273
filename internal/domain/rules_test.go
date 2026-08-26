package domain

import "testing"

func TestRulesValidateTransitions(t *testing.T) {
	if !ValidTransition(StatusDraft, StatusSubmitted) {
		t.Fatal("expected valid")
	}
	if ValidTransition(StatusArchived, StatusDraft) {
		t.Fatal("expected invalid")
	}
}
