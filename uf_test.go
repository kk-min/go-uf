package uf

import (
	"testing"
)

func TestCreateSets(t *testing.T) {
	sets := CreateSets(10)
	if len(sets.parent) != 10 || len(sets.size) != 10 {
		t.Error("Expected 10 sets, got ", len(sets.parent))
	}
}

func TestUnion(t *testing.T) {
	sets := CreateSets(2)
	sets.Union(0, 1)

	if sets.FindSet(0) != 0 {
		t.Error("Expected 0, got ", sets.FindSet(0))
	}

	if sets.FindSet(1) != 0 {
		t.Error("Expected 0, got ", sets.FindSet(1))
	}
}

func TestFindSet(t *testing.T) {
	sets := CreateSets(3)
	sets.Union(0, 1)
	sets.Union(1, 2)

	if sets.FindSet(1) != 0 {
		t.Error("Expected 0, got ", sets.FindSet(1))
	}

	if sets.FindSet(2) != 0 {
		t.Error("Expected 0, got ", sets.FindSet(2))
	}
}
