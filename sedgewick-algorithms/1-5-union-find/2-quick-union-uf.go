package unionfind

type QuickUnionUF struct {
	root []int
}

func NewQuickUnionUF(n int) UnionFind {
	root := make([]int, n)
	for i := range root {
		root[i] = i
	}
	return &QuickUnionUF{root}
}

func (u *QuickUnionUF) Connected(i, j int) bool {
	return u.find(i) == u.find(j)
}

func (u *QuickUnionUF) Union(i, j int) {
	ri, rj := u.find(i), u.find(j)
	u.root[ri] = rj
}

func (u *QuickUnionUF) find(i int) int {
	for i != u.root[i] {
		i = u.root[i]
	}
	return i
}
