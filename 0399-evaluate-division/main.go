package _399_evaluate_division

type Node struct {
	Value  float64
	Parent *Node
}

func NewNode(value float64) *Node {
	node := &Node{
		Value: value,
	}
	node.Parent = node
	return node
}

func findParent(node *Node) *Node {
	if node == node.Parent {
		return node
	}
	node.Parent = findParent(node.Parent)
	return node.Parent
}

func union(node1, node2 *Node, num float64, maps map[string]*Node) {
	p1, p2 := findParent(node1), findParent(node2)
	if p1 != p2 {
		ratio := node2.Value * num / node1.Value
		for k, v := range maps {
			if findParent(v) == p1 {
				maps[k].Value *= ratio
			}
		}
		p1.Parent = p2
	}
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	res := make([]float64, 0)
	maps := make(map[string]*Node)

	for k, v := range equations {
		v1, ok1 := maps[v[0]]
		v2, ok2 := maps[v[1]]
		if !ok1 && !ok2 {
			p1, p2 := NewNode(values[k]), NewNode(1)
			maps[v[0]], maps[v[1]] = p1, p2
			p1.Parent = p2
		} else if !ok1 {
			p2 := findParent(v2)
			p1 := NewNode(v2.Value * values[k])
			maps[v[0]] = p1
			p1.Parent = p2
		} else if !ok2 {
			p1 := findParent(v1)
			p2 := NewNode(v1.Value / values[k])
			maps[v[1]] = p2
			p2.Parent = p1
		} else {
			union(v1, v2, values[k], maps)
		}
	}

	for _, v := range queries {
		v1, ok1 := maps[v[0]]
		v2, ok2 := maps[v[1]]
		if ok1 && ok2 && findParent(v1) == findParent(v2) {
			res = append(res, v1.Value/v2.Value)
		} else {
			res = append(res, -1.0)
		}
	}

	return res
}

func calcEquation1(equations [][]string, values []float64, queries [][]string) []float64 {
	uf := newUnionFind()
	for i, e := range equations {
		uf.Add(e[0])
		uf.Add(e[1])
		uf.Union(e[0], e[1], values[i])
	}
	ans := make([]float64, len(queries))
	for i, query := range queries {
		ans[i] = uf.Ratio(query[0], query[1])
	}
	return ans
}

type UnionFind struct {
	root  map[string]string
	ratio map[string]float64
}

func newUnionFind() *UnionFind {
	return &UnionFind{
		root:  map[string]string{},
		ratio: map[string]float64{},
	}
}

func (uf *UnionFind) Add(to string) {
	if uf.root[to] != "" {
		return
	}
	uf.ratio[to] = 1.0
	uf.root[to] = to
}

func (uf *UnionFind) Union(p, q string, ratio float64) {
	pr, qr := uf.findRoot(p), uf.findRoot(q)
	if pr == qr {
		return
	}
	uf.ratio[pr] = uf.ratio[q] / uf.ratio[p] * ratio
	uf.root[pr] = qr
}

func (uf *UnionFind) Ratio(from, to string) float64 {
	pr, qr := uf.findRoot(from), uf.findRoot(to)
	if pr != qr || pr == "" || qr == "" {
		return -1.0
	}
	return uf.ratio[from] / uf.ratio[to]
}

func (uf *UnionFind) findRoot(p string) string {
	if pr, ok := uf.root[p]; ok {
		if pr == p {
			return pr
		}
		prr := uf.findRoot(pr)
		uf.ratio[p] *= uf.ratio[pr]
		uf.root[p] = prr
		return prr
	}
	return ""
}
