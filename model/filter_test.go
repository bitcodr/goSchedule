package model

import "testing"

func TestFilterModel(t *testing.T) {
	filter := Filter{
		Page:  1,
		Limit: 10,
	}

	if filter.Page != 1 {
		t.Errorf("Expected page to be 1, got %d", filter.Page)
	}

	if filter.Limit != 10 {
		t.Errorf("Expected limit to be 10, got %d", filter.Limit)
	}
}

func TestEmptyFilterModel(t *testing.T) {
	filter := Filter{}

	if filter.Page != 0 {
		t.Errorf("Expected page to be 0, got %d", filter.Page)
	}

	if filter.Limit != 0 {
		t.Errorf("Expected limit to be 0, got %d", filter.Limit)
	}
}

func TestFilterModelWithLargeValues(t *testing.T) {
	filter := Filter{
		Page:  1000,
		Limit: 100,
	}

	if filter.Page != 1000 {
		t.Errorf("Expected page to be 1000, got %d", filter.Page)
	}

	if filter.Limit != 100 {
		t.Errorf("Expected limit to be 100, got %d", filter.Limit)
	}
}
