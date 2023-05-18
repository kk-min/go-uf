package uf

// Union-Find data structure for disjoint sets. Internally holds an array of ancestors for member i at parent[i], and the size of set i at size[i].
type UF struct {
	parent []int
	size   []int
}

// Creates n disjoint sets.
func CreateSets(n int) UF {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return UF{parent: parent, size: size}
}

// Finds the set that x belongs to.
func (sets *UF) FindSet(x int) int {
	if sets.parent[x] == x {
		return x
	}
	sets.parent[x] = sets.FindSet(sets.parent[x])
	return sets.parent[x]
}

// Unions the sets that x and y belong to. Representative set after union is the set with the larger size.
func (sets *UF) Union(x, y int) {
	parentx := sets.FindSet(x)
	parenty := sets.FindSet(y)

	if parentx != parenty {
		if sets.size[parentx] < sets.size[parenty] {
			sets.parent[parentx] = parenty
			sets.size[parenty] += sets.size[parentx]
		} else {
			sets.parent[parenty] = parentx
			sets.size[parentx] += sets.size[parenty]
		}
	}
}

// Unions the sets that x and y, with x being the resulting representative set.
func (sets *UF) UnionOrdered(x, y int) {
	parentx := sets.FindSet(x)
	parenty := sets.FindSet(y)

	if parentx != parenty {
		sets.parent[parenty] = parentx
		sets.size[parentx] += sets.size[parenty]
	}
}
