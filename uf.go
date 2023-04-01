package uf

type UF struct {
	parent []int
	size   []int
}

func CreateSets(n int) UF {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return UF{parent: parent}
}

func (sets *UF) FindSet(x int) int {
	if sets.parent[x] == x {
		return x
	}
	sets.parent[x] = sets.FindSet(sets.parent[x])
	return sets.parent[x]
}

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
