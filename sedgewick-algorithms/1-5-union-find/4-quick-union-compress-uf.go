package unionfind

type QuickUnionCompressUF struct {
	root []int
}

func NewQuickUnionCompressUF(n int) UnionFind {
	root := make([]int, n)
	for i := range root {
		root[i] = i
	}
	return &QuickUnionCompressUF{root}
}

func (u *QuickUnionCompressUF) Connected(i, j int) bool {
	return u.find(i) == u.find(j)
}

func (u *QuickUnionCompressUF) Union(i, j int) {
	ri, rj := u.find(i), u.find(j)
	u.root[ri] = rj
}

func (u *QuickUnionCompressUF) find(i int) int {
	for i != u.root[i] {
		u.root[i] = u.root[u.root[i]]
		i = u.root[i]
	}
	return i
}
