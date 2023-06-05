package unionfind

import "testing"

func TestQuickFindUF(t *testing.T) {
	testUnionFind(t, NewQuickUnionUF)
}

func TestQuickUnionUF(t *testing.T) {
	testUnionFind(t, NewQuickUnionUF)
}

func TestQuickUnionRankUF(t *testing.T) {
	testUnionFind(t, NewQuickUnionRankUF)
}

func TestQuickUnionCompresskUF(t *testing.T) {
	testUnionFind(t, NewQuickUnionCompressUF)
}

func testUnionFind(t *testing.T, NewUnionFind func(n int) UnionFind) {
	u := NewUnionFind(10)

	u.Union(1, 2)
	u.Union(2, 5)
	u.Union(5, 6)
	u.Union(6, 7)
	u.Union(3, 8)

	if !u.Connected(1, 5) {
		t.Error()
	}
	if !u.Connected(5, 7) {
		t.Error()
	}
	if u.Connected(4, 9) {
		t.Error()
	}

	u.Union(9, 4)

	if !u.Connected(4, 9) {
		t.Error()
	}
}
