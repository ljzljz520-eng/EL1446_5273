package httpapi

import (
	"example.com/inspection14/internal/catalog"
	"example.com/inspection14/internal/query"
	"example.com/inspection14/internal/store"
	"example.com/inspection14/internal/workflow14"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	h := New(catalog.New(s), workflow14.New(s), query.New(s)).Routes()
	r := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if w.Code != 200 {
		t.Fatal(w.Code)
	}
}
