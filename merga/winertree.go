package merga

type FuncComp func(interface{}, interface{}) int8

type DataQueue interface {
	Pop() (interface{}, bool)
}

type mergaTreeNode struct {
	which int
	data  interface{}
}

func (n *mergaTreeNode) isNone() bool {
	return n.which == -1
}

func (n *mergaTreeNode) setNone() {
	n.which = -1
	n.data = nil
}

func nodeWin(win FuncComp, a *mergaTreeNode, b *mergaTreeNode) bool {
	return b.isNone() ||
		(!a.isNone() && win(a.data, b.data) == -1)
}

type WinerTree struct {
	funcWin FuncComp
	nodes   []*mergaTreeNode
	queues  []DataQueue
}

func onLeft(i int) bool {
	return i%2 == 1
}

func nLeft(i int) int {
	return i*2 + 1
}

func nRight(i int) int {
	return i*2 + 2
}

func nParent(i int) int {
	return (i - 1) / 2
}

func nSibling(i int) int {
	parent := nParent(i)
	if onLeft(i) {
		return nRight(parent)
	} else {
		return nLeft(parent)
	}
}

func NewWinerTree(queues []DataQueue, f FuncComp) *WinerTree {
	n := len(queues)*2 - 1

	t := &WinerTree{
		funcWin: f,
		nodes:   make([]*mergaTreeNode, n, n),
		queues:  queues,
	}

	for i := n - 1; i >= n/2; i-- {
		which := i + 1 - len(queues)
		data, ok := queues[which].Pop()

		t.nodes[i] = &mergaTreeNode{}
		if ok {
			t.nodes[i].which = which
			t.nodes[i].data = data
		} else {
			t.nodes[i].setNone()
		}
	}

	for i := n/2 - 1; i >= 0; i-- {
		left := nLeft(i)
		right := nRight(i)

		if nodeWin(t.funcWin, t.nodes[left], t.nodes[right]) {
			t.nodes[i] = t.nodes[left]
		} else {
			t.nodes[i] = t.nodes[right]
		}
	}

	return t
}

func (t *WinerTree) Extract() (ret interface{}, ok bool) {
	if t.nodes[0].isNone() {
		return nil, false
	}
	ret, ok = t.nodes[0].data, true

	which := t.nodes[0].which
	ni := which + len(t.queues) - 1
	data, ok0 := t.queues[which].Pop()

	if ok0 {
		t.nodes[ni].data = data
	} else {
		t.nodes[ni].setNone()
		if ni != 0 {
			ni = nSibling(ni)
		}
	}

	for ni > 0 {
		sibling := nSibling(ni)
		parent := nParent(ni)
		if nodeWin(t.funcWin, t.nodes[ni], t.nodes[sibling]) {
			t.nodes[parent] = t.nodes[ni]
		} else {
			t.nodes[parent] = t.nodes[sibling]
		}
		ni = parent
	}

	return
}
