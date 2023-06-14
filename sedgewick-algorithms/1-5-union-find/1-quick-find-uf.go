package unionfind

type QuickFindUF struct {
	root []int
}

func NewQuickFindUF(n int) UnionFind {
	root := make([]int, n)
	for i := range root {
		root[i] = i
	}
	return &QuickFindUF{root}
}

func (u *QuickFindUF) Connected(i, j int) bool {
	return u.root[i] == u.root[j]
}

func (u *QuickFindUF) Union(i, j int) {
	ri, rj := u.root[i], u.root[j]
	for k := 0; k < len(u.root); k++ {
		if u.root[k] == ri {
			u.root[k] = rj
		}
	}
}
