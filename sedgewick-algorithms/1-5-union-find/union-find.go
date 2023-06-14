package unionfind

type UnionFind interface {
	Connected(p, q int) bool
	Union(p, q int)
}
