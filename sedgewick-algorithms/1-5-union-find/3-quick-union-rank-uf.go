package unionfind

type QuickUnionRankUF struct {
	root []int
	rank []int // subtree height
}

func NewQuickUnionRankUF(n int) UnionFind {
	root := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		root[i] = i
		rank[i] = 1
	}
	return &QuickUnionRankUF{root, rank}
}

func (u *QuickUnionRankUF) Connected(i, j int) bool {
	return u.find(i) == u.find(j)
}

func (u *QuickUnionRankUF) Union(i, j int) {
	ri, rj := u.find(i), u.find(j)
	if ri == rj {
		return
	}
	if u.root[ri] < u.root[rj] {
		u.root[ri] = rj
	} else if u.root[rj] < u.root[ri] {
		u.root[rj] = ri
	} else {
		u.root[ri] = rj
		u.rank[rj]++
	}
}

func (u *QuickUnionRankUF) find(i int) int {
	for i != u.root[i] {
		i = u.root[i]
	}
	return i
}
