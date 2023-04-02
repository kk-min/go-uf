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

func TestUnion2(t *testing.T) {
	sets := CreateSets(5)
	edges := [][]int{{0, 1}, {0, 2}, {3, 4}, {2, 3}}

	for _, edge := range edges {
		sets.Union(edge[0], edge[1])
	}

	for i := 0; i < 5; i++ {
		if sets.FindSet(i) != 0 {
			t.Error("Expected 0, got ", sets.FindSet(i))
		}
	}
}

func TestUnionLarge(t *testing.T) {
	sets := CreateSets(10000)

	for i := 0; i < 9999; i++ {
		sets.Union(i, i+1)
	}

	for i := 0; i < 10000; i++ {
		if sets.FindSet(i) != 0 {
			t.Error("Expected 0, got ", sets.FindSet(i))
		}
	}
}

func TestUnionSizeRanking(t *testing.T) {
	sets := CreateSets(6)
	edges := [][]int{{0, 1}, {2, 3}, {2, 4}, {2, 5}, {0, 4}}

	for _, edge := range edges {
		sets.Union(edge[0], edge[1])
	}

	if sets.FindSet(0) != 2 {
		t.Error("Expected 2, got ", sets.FindSet(0))
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
