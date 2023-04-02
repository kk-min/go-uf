# Go Union Find / Disjoint Set

Union-Find / Disjoint Set data structure implementation for Go.

## Features

- **Path Compression:** FindSet(n) will point the parent value of n and it's intermediate parents to the root
- **Ranking via Size:** The set with a bigger size will act as the representative during a Union

## Usage

```go
// Create 5 initial disjoint sets
sets := uf.CreateSets(5)

// Union the sets based on edge (connection) information
edges := [][]int{{0,1}, {2,3}, {3,4}}
for _, edge := range edges {
	sets.Union(edge[0], edge[1])
}

// Get root (parent set) of a number
uf.FindSet(0) // 0
uf.FindSet(1) // 0
uf.FindSet(2) // 2
uf.FindSet(3) // 2
uf.FindSet(4) // 2
```
