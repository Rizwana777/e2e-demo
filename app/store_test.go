package app

import (
	"testing"
)

func TestStore_AddAndList(t *testing.T) {
	store := &Store{}

	store.Add("Learn E2E")

	items := store.List()

	if len(items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(items))
	}

	if items[0] != "Learn E2E" {
		t.Fatalf("expected 'Learn E2E', got %s", items[0])
	}
}
